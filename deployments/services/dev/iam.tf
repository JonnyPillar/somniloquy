# Defines a user that should be able to write to you test bucket
resource "aws_iam_user" "somniloquy_upload_user" {
    name = "dev_somniloquy_upload_user"
}

resource "aws_iam_access_key" "somniloquy_upload_user" {
    user = "${aws_iam_user.somniloquy_upload_user.name}"
}

resource "aws_iam_user_policy" "somniloquy_upload_user_ro" {
    name = "somniloquy_upload_user_policy"
    user = "${aws_iam_user.somniloquy_upload_user.name}"
    policy= <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "s3:*",
            "Resource": [
                "arn:aws:s3:::${var.audio_upload_bucket_name}",
                "arn:aws:s3:::${var.audio_upload_bucket_name}/*"
            ]
        }
    ]
}
EOF
}