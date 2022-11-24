package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRandomPaymentName(t *testing.T, paymentName string) PaymentBy {
	paymentBy, err := testQueries.CreatePaymentName(context.Background(), paymentName)
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentBy)
	assert.NotEqual(t, "", paymentBy.PaymentName) // ngak boleh empty
	assert.Equal(t, paymentName, paymentBy.PaymentName)
	assert.NotZero(t, paymentBy.ID)

	return paymentBy
}

func TestCreatePaymentName(t *testing.T) {
	var paymentNames []string = []string{"GoPay", "ShopeePay", "OVO"}
	//var paymentNamesResults []PaymentBy

	t.Run("PaymentNames not empty", func(t *testing.T) {

		for _, paymentName := range paymentNames {
			//paymentNamesResults = append(paymentNamesResults, createRandomPaymentName(t, paymentName))
			createRandomPaymentName(t, paymentName)
		}

		// for _, result := range paymentNamesResults {
		// }
	})

	// t.Run("PaymentName is Empty", func(t *testing.T) {
	// 	createRandomPaymentName(t, "")
	// })

}

func TestDeletePaymentBy(t *testing.T) {
	t.Run("Delete PaymentName with Existing Id", func(t *testing.T) {
		pymtName := createRandomPaymentName(t, "Livin")
		assert.NotEmpty(t, pymtName)

		err := testQueries.DeletePaymenBy(context.Background(), pymtName.ID)
		assert.NoError(t, err)
	})

	// t.Run("Delete PaymentName with non-Existing Id", func(t *testing.T) {
	// 	err := testQueries.DeletePaymenBy(context.Background(), 999999999)
	// 	assert.Error(t, err)
	// })
}
