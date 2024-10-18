package provider

import (
	"context"
	"os"

	client2 "github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider = &ProxmoxSDNProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ProxmoxSDNProvider{
			version: version,
		}
	}
}

type ProxmoxSDNProvider struct {
	version string
}

func (p *ProxmoxSDNProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "proxmox-sdn"
	resp.Version = p.version
}

func (p *ProxmoxSDNProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "The Proxmox host",
				Required:    true,
			},
			"username": schema.StringAttribute{
				Description: "The Proxmox username",
				Required:    true,
			},
			"password": schema.StringAttribute{
				Description: "The Proxmox password",
				Required:    true,
				Sensitive:   true,
			},
		},
	}
}

type ProxmoxSDNProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *ProxmoxSDNProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring ProxmoxSDNProvider Client")
	var config ProxmoxSDNProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "proxmoSDNHost", config.Host)
	ctx = tflog.SetField(ctx, "proxmoSDNUsername", config.Username)
	ctx = tflog.SetField(ctx, "proxmoSDNPassword", config.Password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "proxmoSDNPassword")

	tflog.Debug(ctx, "Creating ProxmoxSDNProvider Client")

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Proxmox host",
			"The Provider cannot create the client without a host",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Proxmox username",
			"The Provider cannot create the client without a username",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Proxmox password",
			"The Provider cannot create the client without a password",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("PROXMOXSDN_HOST")
	username := os.Getenv("PROXMOXSDN_USERNAME")
	password := os.Getenv("PROXMOXSDN_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Proxmox host",
			"The Provider cannot create the client without a host"+
				"Set the PROXMOXSDN_HOST environment variable or provide the host attribute"+
				"if either is already set, ensure that the value is not empty",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Proxmox username",
			"The Provider cannot create the client without a username"+
				"Set the PROXMOXSDN_USERNAME environment variable or provide the username attribute"+
				"if either is already set, ensure that the value is not empty",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Proxmox password",
			"The Provider cannot create the client without a password"+
				"Set the PROXMOXSDN_PASSWORD environment variable or provide the password attribute"+
				"if either is already set, ensure that the value is not empty",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	client, err := client2.NewSSHProxmoxClient(username, password, host)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create client: %w", err.Error())
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured proxmox-sdn client", map[string]any{"success": true})
}

func (p *ProxmoxSDNProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProxmoxSDNZoneDataSource,
	}
}

func (p *ProxmoxSDNProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProxmoxSDNZoneResource,
	}
}
