package main

import (
	"main/ascii"
	"testing"
)

func TestWrongInput(t *testing.T) {
	examples := []string{"حسن", "😊", "doesn't لا يعمل work"}
	for _, example := range examples {
		_, err := ascii.GenerateAscii(example, "standard")
		if err == nil {
			t.Errorf(`%s: should have failed`, example)
		}
	}
}

func TestWrongBanner(t *testing.T) {
	_, err := ascii.GenerateAscii("test", "sagasdgasaj")
	if err == nil {
		t.Error(`should have failed due to wrong banner type`)
	}
}

func TestStandardGenerate(t *testing.T) {
	examples := []string{"test", "Reboot01", "123??", "{123}\n<Hello> (World)!"}
	results := []string{
		` _                  _    
| |                | |   
| |_    ___   ___  | |_  
| __|  / _ \ / __| | __| 
\ |_  |  __/ \__ \ \ |_  
 \__|  \___| |___/  \__| 

`,
		` _____           _                       _                
|  __ \         | |                     | |     ___    _  
| |__) |   ___  | |__     ___     ___   | |_   / _ \  / | 
|  _  /   / _ \ | '_ \   / _ \   / _ \  | __| | | | | | | 
| | \ \  |  __/ | |_) | | (_) | | (_) | \ |_  | |_| | | | 
|_|  \_\  \___| |_.__/   \___/   \___/   \__|  \___/  |_| 

`,
		`                     ___    ___   
 _   ____    _____  |__ \  |__ \  
/ | |___ \  |___ /     ) |    ) | 
| |   __) |   |_ \    / /    / /  
| |  / __/   ___) |  |_|    |_|   
|_| |_____| |____/   (_)    (_)   

`,
		`   __                     __    
  / /  _   ____    _____  \ \   
 | |  / | |___ \  |___ /   | |  
/ /   | |   __) |   |_ \    \ \ 
\ \   | |  / __/   ___) |   / / 
 | |  |_| |_____| |____/   | |  
  \_\                     /_/   

   __  _    _          _   _          __            __ __          __                 _       _  __    _  
  / / | |  | |        | | | |         \ \          / / \ \        / /                | |     | | \ \  | | 
 / /  | |__| |   ___  | | | |   ___    \ \        | |   \ \  /\  / /    ___    _ __  | |   __| |  | | | | 
< <   |  __  |  / _ \ | | | |  / _ \    > >       | |    \ \/  \/ /    / _ \  | '__| | |  / _' |  | | | | 
 \ \  | |  | | |  __/ | | | | | (_) |  / /        | |     \  /\  /    | (_) | | |    | | | (_| |  | | |_| 
  \_\ |_|  |_|  \___| |_| |_|  \___/  /_/         | |      \/  \/      \___/  |_|    |_|  \__,_|  | | (_) 
                                                   \_\                                           /_/      

`}
	for i, example := range examples {
		ascii, err := ascii.GenerateAscii(example, "standard")
		if err != nil {
			t.Errorf(`Error: %s`, err.Error())
		} else if ascii != results[i] {
			t.Errorf("ascii generated for %s doesn't match\nwant:\n%s\ngot:\n%s", example, results[i], ascii)
		}
	}

}

func TestShadowGenerate(t *testing.T) {
	examples := []string{"test", "Reboot01", `$% "=`}
	results := []string{
		`                                    
  _|                         _|     
_|_|_|_|   _|_|     _|_|_| _|_|_|_| 
  _|     _|_|_|_| _|_|       _|     
  _|     _|           _|_|   _|     
    _|_|   _|_|_| _|_|_|       _|_| 

`,
		`                                                                  
_|_|_|            _|                           _|       _|     _| 
_|    _|   _|_|   _|_|_|     _|_|     _|_|   _|_|_|_| _|  _| _|_| 
_|_|_|   _|_|_|_| _|    _| _|    _| _|    _|   _|     _|  _|   _| 
_|    _| _|       _|    _| _|    _| _|    _|   _|     _|  _|   _| 
_|    _|   _|_|_| _|_|_|     _|_|     _|_|       _|_|   _|     _| 

`,
		`                        _|  _|            
  _|   _|_|    _|       _|  _|            
_|_|_| _|_|  _|                _|_|_|_|_| 
_|_|       _|                             
  _|_|   _|  _|_|              _|_|_|_|_| 
_|_|_| _|    _|_|                         
  _|                                      

`,
	}
	for i, example := range examples {
		ascii, err := ascii.GenerateAscii(example, "shadow")
		if err != nil {
			t.Errorf(`Error: %s`, err.Error())
		} else if ascii != results[i] {
			t.Errorf("ascii generated for %s doesn't match\nwant:\n%s\ngot:\n%s", example, results[i], ascii)
		}
	}

}

func TestThinkertoyGenerate(t *testing.T) {
	examples := []string{"test", "Reboot01", "123 T/fs#R"}
	results := []string{
		`                 
 o           o  
 |           |  
-o- o-o o-o -o- 
 |  |-'  \   |  
 o  o-o o-o  o  

`,
		`                                       
o--o      o             o   o-o    0   
|   |     |             |  o  /o  /|   
O-Oo  o-o O-o  o-o o-o -o- | / | o |   
|  \  |-' |  | | | | |  |  o/  o   |   
o   o o-o o-o  o-o o-o  o   o-o  o-o-o 

`,
		`                                                       
  0    --  o-o        o-O-o     o  o-o      | |  o--o  
 /|   o  o    |         |      /   |       -O-O- |   | 
o |     /   oo          |     o   -O-  o-o  | |  O-Oo  
  |    /      |         |    /     |    \  -O-O- |  \  
o-o-o o--o o-o          o   o      o   o-o  | |  o   o 

`,
	}
	for i, example := range examples {
		ascii, err := ascii.GenerateAscii(example, "thinkertoy")
		if err != nil {
			t.Errorf(`Error: %s`, err.Error())
		} else if ascii != results[i] {
			t.Errorf("ascii generated for %s doesn't match\nwant:\n%s\ngot:\n%s", example, results[i], ascii)
		}
	}

}
