resource "aws_s3_bucket" "repo" {
	bucket = "coffee-artifacts"
    acl    = "private"
}

# vim:ts=4:sw=4:sts=4:expandtab:syntax=terraform
