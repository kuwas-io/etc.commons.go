/** @imports */
package tools
import ( "github.com/onsi/ginkgo" )

/** @parameters */
/** @exports */
func Should( o string , c func() ) ( bool ) {
  return ginkgo.It( "should return results " + o , c )
}
