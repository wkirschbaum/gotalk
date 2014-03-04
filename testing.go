import (
        "testing"
)

func TestMyFunction(t *testing.T) {
        c, err := aetest.NewContext(nil)
        if err != nil {
                t.Fatal(err)
        }
        defer c.Close()

        // Run code and tests requiring the appengine.Context using c.
}