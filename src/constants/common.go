package constants

const AppName = "WednesDaySolutions"
const Port = ":5176"

var BookingState = struct {
	Completed string
	Booked    string
	Failed    string
}{
	Completed: "COMPLETED",
	Booked:    "BOOKED",
	Failed:    "FAILED",
}
