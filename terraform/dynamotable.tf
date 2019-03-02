resource "aws_dynamodb_table" "artifacts" {
	name 						= "Artifacts"
	billing_mode 		= "PROVISIONED"
	read_capacity 	= 2
	write_capacity 	= 2
	hash_key        = "artifact"

	attribute {
		name = "artifact"
		type = "S"
	}

}
