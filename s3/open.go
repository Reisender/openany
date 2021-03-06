package s3

import (
	"io"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	//"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Open(s3URL string) (io.ReadCloser, error) {
	var region string
	var ok bool
	if region, ok = os.LookupEnv("AWS_DEFAULT_REGION"); !ok {
		region = "us-east-1" // default region
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, // loades ~/.aws/config as well
	}))
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(region),
	})

	parsed, err := url.Parse(s3URL)
	if err != nil {
		return nil, err
	}

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(parsed.Host),
		Key:    aws.String(parsed.Path),
	})

	return obj.Body, err
}
