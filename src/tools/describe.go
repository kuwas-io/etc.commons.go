/** @imports */
package tools
import ( "github.com/onsi/ginkgo" )

/** @parameters */
/** @exports */
func Describe( s string , c func() , b func() , a func() ) bool {
  var t = Title( [] string { s } , true )
  return ginkgo.Describe( t , func() {
    ginkgo.BeforeEach( b ) ; c() ;
    ginkgo.AfterEach( a ) ;
  })
}
