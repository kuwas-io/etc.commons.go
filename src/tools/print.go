/** @imports */
package tools
import (
  "github.com/davecgh/go-spew/spew"
  "github.com/sirupsen/logrus"
)

/** @parameters */
/** @exports */
func Print( t string , s [] string , o interface {} , d bool ) {
  var key = append( s , "%v" )
  var template = Title( key , false )
  if d { o = spew.Sdump( o ) }
  switch t {
    case "p" :  logrus.Panicf( template , o )
    case "f" :  logrus.Fatalf( template , o )
    case "e" :  logrus.Errorf( template , o )
    case "w" :  logrus.Warnf(  template , o )
    case "i" :  logrus.Infof(  template , o )
    case "d" :  logrus.Debugf( template , o )
    case "t" :  logrus.Tracef( template , o )
  }
  return
}
