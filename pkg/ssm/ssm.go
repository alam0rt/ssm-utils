package ssm

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Parameter struct {
	*ssm.Parameter
}

var (
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc = ssm.New(sess, &aws.Config{
		MaxRetries: aws.Int(3),
	})
	// Parameters is a list of SSM parameters
	recursive      = true
	withDecryption = true
)

func (o Parameter) OutputToInput() ssm.PutParameterInput {
	return ssm.PutParameterInput{
		Name:     o.Name,
		Value:    o.Value,
		DataType: o.DataType,
		Type:     o.Type,
	}
}

// GetParameters takes a path and returns all parameters under it recursively
func GetParameters(path string) ([]*ssm.Parameter, error) {
	var parameters []*ssm.Parameter
	input := &ssm.GetParametersByPathInput{
		Path:           &path,
		Recursive:      &recursive,
		WithDecryption: &withDecryption,
	}
	for {
		resp, err := svc.GetParametersByPath(input)
		if err != nil {
			return nil, err
		}

		parameters = append(parameters, resp.Parameters...)
		if resp.NextToken == nil {
			break
		}
		input.SetNextToken(*resp.NextToken)
	}
	if len(parameters) == 0 {
		err := errors.New("there are no parameters at the given path")
		return nil, err
	}
	return parameters, nil
}
