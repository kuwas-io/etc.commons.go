/** @imports */
package cli
import (
  "github.com/sirupsen/logrus"
  "github.com/urfave/cli/v2"
  "os"
)

/** @parameters */
/** @exports */
func App( v string , title string , description string , cmd cli.Commands ) ( e error ) {

  var o = & cli.App {

    /* Cores */
    Version :     v ,
    Name :        title ,
    Description : description ,
    Usage :       " " ,
    Authors :     [] * cli.Author {
      & cli.Author { "@kuwas" , "40496186+kuwas@users.noreply.github.com" } ,
    } ,

    /* Hooks */
    Before : func( c * cli.Context ) ( e error ) {
      logrus.SetOutput( os.Stdout )
      logrus.SetFormatter( & logrus.TextFormatter { FullTimestamp : true } )
      if c.Bool( "debug" ) { logrus.SetLevel( logrus.DebugLevel ) }
      return
    } ,

    /* Flags */
    Flags : [] cli.Flag {
      & cli.BoolFlag { Name : "debug" , Aliases : [] string { "d" } } ,
    } ,

    /* Commands */
    Commands : cmd ,
    /*****/

  }
  return o.Run( os.Args )

}
