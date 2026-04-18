package clients

import "fmt"

func (c *Clients) Close() error {
	var errs []error
	if err := c.authConn.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := c.hotelConn.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := c.bookingConn.Close(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("errors closing connections: %v", errs)
	}

	return nil
}
