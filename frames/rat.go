package frames

// RatDance maps the command endpoint registration
var RatDance = DefaultFrameType(ratDanceFrames)

var ratDanceFrames = []string{
	// Frame 1: Left Step
	`
       _.._  ,-----..__
    .-' ._.'`  _ )_    `-.
   ((_.-'    .' /_\\      \
            /   \_/       |
        ()__\  \_|        /
      _ .--..`""  `---.._.'
     ( (     )         
      \ \   / /        
      /_/  /_/         
	`,

	// Frame 2: Neutral Lean
	`
       _.._       _..--''
    .-' ._.'`---"`_ )_   `-.
   ((_.-'       .' /_\\     \
               /   \_/      |
           ()__\  \_|       /
         _ .--..`""  `---..'
        ( (     )
         \ \   / \
          \_\  \_\
	`,

	// Frame 3: Right Step
	`
    __..-----,  _.._ 
 .-'    _(_ _  `'._. '-.
/      //_\ \       '-._))
|       \_/   \
\        |_/  /__()
 '..---`  ""`..--. _
             (     ) )
              \ \   / /
              /_/  /_/
	`,

	// Frame 4: Neutral Lean (Return transition)
	`
       _.._       _..--''
    .-' ._.'`---"`_ )_   `-.
   ((_.-'       .' /_\\     \
               /   \_/      |
           ()__\  \_|       /
         _ .--..`""  `---..'
        ( (     )
         \ \   / \
          \_\  \_\
	`,
}
