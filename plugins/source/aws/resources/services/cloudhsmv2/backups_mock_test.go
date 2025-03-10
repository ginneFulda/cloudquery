package cloudhsmv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildHSMBackups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudhsmv2Client(ctrl)

	var backups []types.Backup
	if err := faker.FakeObject(&backups); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeBackups(
		gomock.Any(),
		&cloudhsmv2.DescribeBackupsInput{},
		gomock.Any(),
	).Return(
		&cloudhsmv2.DescribeBackupsOutput{Backups: backups},
		nil,
	)

	return client.Services{Cloudhsmv2: mock}
}

func TestBackups(t *testing.T) {
	client.AwsMockTestHelper(t, Backups(), buildHSMBackups, client.TestOptions{})
}
