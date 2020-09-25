// protocol buffers version 
syntax = "proto3"; //proto1 / proto2 / proto3

import "<folder_name>/<file_name>"
import "go-protocol-buffers/example.proto"

package "<name>"

message MessageName {
    int32 id = 1;   // [type] [name] = [order];
    string first_name = 2;
    bool is_valid = 3;

    // coments
    /*
        multiline
        comment
    */

    repeated string phone_numbers = 4; //array

    enum EyeColor {
        UNKNOWN_EYE_COLOR = 1;
        EYE_GREEN = 1;
        EYE_BROWN = 2;
        EYE_BLUE = 2;
    }

    EyeColor eye_color = 8;

    ... <package>.<type> <name> = <order> //eg.
    my.package.Date date = 9;
    
}