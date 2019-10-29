
data "coinbase_address" "api_address" {}


output "endpoint" {
  value = ["${data.coinbase_address.api_address.api_endpoint}"]
}
