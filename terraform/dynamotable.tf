resource "aws_dynamodb_table" "artifacts" {
	name 						= "Artifacts"
	billing_mode 		= "PROVISIONED"
	read_capacity 	= 2
	write_capacity 	= 2
	hash_key        = "Artifact"

	attribute {
		name = "Artifact"
		type = "S"
	}

	attribute {
		name = "SourceChecksum"
		type = "N"
	}

	attribute {
		name = "Last deployment"
		type = "S"
	}

}
