package tests

import(
    "testing"
    "PortProgram/myapp/cmd/utils"
)

func testansiicolors(t *testing.T) {
    if utils.Colors["red"] != "\x1b[31m" {
        t.Errorf("Color was not red")
    }
    if utils.Colors["green"] != "\x1b[32m" {
        t.Errorf("Color was not green")
    }
    if utils.Colors["yellow"] != "\x1b[33m" {
        t.Errorf("Color was not yellow")
    }
    if utils.Colors["blue"] != "\x1b[34m" {
        t.Errorf("Color was not blue")
    }
    if utils.Colors["magenta"] != "\x1b[35m" {
        t.Errorf("Color was not magenta")
    }
    if utils.Colors["cyan"] != "\x1b[36m" {
        t.Errorf("Color was not cyan")
    }
    if utils.Colors["white"] != "\x1b[37m" {
        t.Errorf("Color was not white")
    }
    if utils.Colors["brightred"] != "\x1b[91m" {
        t.Errorf("Color was not brightred")
    }
    if utils.Colors["brightgreen"] != "\x1b[92m" {
        t.Errorf("Color was not brightgreen")
    }
    if utils.Colors["brightyellow"] != "\x1b[93m" {
        t.Errorf("Color was not brightyellow")
    }
    if utils.Colors["brightblue"] != "\x1b[94m" {
        t.Errorf("Color was not brightblue")
    }
    if utils.Colors["brightmagenta"] != "\x1b[95m" {
        t.Errorf("Color was not brightmagenta")
    }
    if utils.Colors["brightcyan"] != "\x1b[96m" {
        t.Errorf("Color was not brightcyan")
    }
    if utils.Colors["brightwhite"] != "\x1b[97m" {
        t.Errorf("Color was not brightwhite")
    }
    if utils.Colors["reset"] != "\x1b[0m" {
        t.Errorf("Color was not reset")
    }
}

func TestClearScreen(t *testing.T) {
    clearscreenstring := utils.ClearScreen()

    if clearscreenstring != "\033[H\033[2J" {
		t.Errorf("Screen was not cleared properly")
	}
}

func TestCheckForIllegalCharacters(t *testing.T) {
    illegalwords := []string{"Peop!e", "Peo@ple", "Peo$ple", "Peo%ple", "Peo^ple", "Peo&ple", "Peo*ple", "Peo(ple", "Peo)ple", "Peo+ple", "Peo=ple", "Peo[ple", "Peo]ple", "Peo{ple", "Peo}ple", "Peo|ple", "Peo;ple", "Peo:ple", "Peo'ple", "Peo,ple", "Peo<ple", "Peo>ple", "Peo/ple", "Peo?ple", "Peo`ple", "Peo~ple", "Peo\\ple"}

    for _, word := range illegalwords {
        if utils.CheckForIllegalCharacters(word) == false {
            t.Errorf("Illegal character should have been detected")
        }
    }


    legalwords := []string{"Unlike","most","MSDOS","programs","at","the","time","Microsoft","Word","was","designed","to","be","used","with","a","mouse","Advertisements","depicted","the","Microsoft","Mouse","and","described","Word","as","a","WYSIWYG","windowed","word","processor","with","the","ability","to","undo","and","display","bold","italic","and","underlined","text","although","it","could","not","render","fonts","It","was","not","initially","popular","since","its","user","interface","was","different","from","the","leading","word","processor","at","the","time","WordStar","However","Microsoft","steadily","improved","the","product","releasing","versions","through","","over","the","next","six","years","In","","Microsoft","ported","Word","to","the","classic","Mac","OS","known","as","Macintosh","System","Software","at","the","time","This","was","made","easier","by","Word","for","DOS","having","been","designed","for","use","with","highresolution","displays","and","laser","printers","even","though","none","were","yet","available","to","the","general","public","It","was","also","notable","for","its","very","fast","cutandpaste","function","and","unlimited","number","of","undo","operations","which","are","due","to","its","usage","of","the","piece","table","data","structure"}

    for _, current := range legalwords {
        if utils.CheckForIllegalCharacters(current) == true {
            t.Error("Only legal characters")
        }
    }
}
