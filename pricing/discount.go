package pricing

// Factory Function
// Returns a closure that remembers the discount percentage
func CreateDiscounter(percent float64) func(float64) float64 {

	return func(price float64) float64 {

		discount := price * percent / 100

		return price - discount
	}

}