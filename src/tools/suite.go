/** @imports */
package tools
import (
  "github.com/onsi/ginkgo"
  "github.com/onsi/gomega"
  "testing"
)

/** @parameters */
/** @exports */
func Suite( o string , c * testing.T ) ( bool ) {
  var t = Title( [] string { o } , false )
  gomega.RegisterFailHandler( ginkgo.Fail )
  return ginkgo.RunSpecs( c , t )
}
