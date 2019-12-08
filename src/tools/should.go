/** @imports */
package tools
import ( "github.com/onsi/ginkgo" )

/** @parameters */
/** @exports */
func Should( s string , c func() ) bool {
  return ginkgo.It( "should return results " + s , c )
}
