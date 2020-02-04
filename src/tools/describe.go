/** @imports */
package tools
import ( "github.com/onsi/ginkgo" )

/** @parameters */
/** @exports */
func Describe( o string , c func() , b func() , a func() ) ( bool ) {
  var t = Title( [] string { o } , true )
  return ginkgo.Describe( t , func() {
    ginkgo.BeforeEach( b ) ; c() ;
    ginkgo.AfterEach( a ) ;
  })
}
