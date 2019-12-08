/** @imports */
package tools
import (
  "github.com/onsi/ginkgo"
  "github.com/onsi/gomega"
  "testing"
)

/** @parameters */
/** @exports */
func Suite( s string , c * testing.T ) bool {
  var t = Title( [] string { s } , false )
  gomega.RegisterFailHandler( ginkgo.Fail )
  return ginkgo.RunSpecs( c , t )
}
