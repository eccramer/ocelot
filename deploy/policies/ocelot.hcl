# configuration data
path "secret/data/config/ocelot/*"
{
  capabilities = ["create", "read", "update", "delete", "list"]
}

# user credential data
path "secret/data/creds/*"
{
  capabilities = ["create", "read", "update", "delete", "list"]
}