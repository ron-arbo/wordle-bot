package board

var bigLetters = map[rune][]string{
	'A': {
		"  AAA  ",
		" AA AA ",
		"AA   AA",
		"AAAAAAA",
		"AA   AA",
		"AA   AA",
		"AA   AA",
	},
	'B': {
		"BBBBBB ",
		"BB    B",
		"BB    B",
		"BBBBBB ",
		"BB    B",
		"BB    B",
		"BBBBBB ",
	},
	'C': {
		"  CCCCCC ",
		" CC      ",
		"CC       ",
		"CC       ",
		"CC       ",
		" CC      ",
		"  CCCCCC ",
	},
	'D': {
		"DDDDD  ",
		"DD   DD",
		"DD    D",
		"DD    D",
		"DD    D",
		"DD   DD",
		"DDDDD  ",
	},
	'E': {
		"EEEEEEE",
		"EE     ",
		"EE     ",
		"EEEEEE ",
		"EE     ",
		"EE     ",
		"EEEEEEE",
	},
	'F': {
		"FFFFFFF",
		"FF     ",
		"FF     ",
		"FFFFF  ",
		"FF     ",
		"FF     ",
		"FF     ",
	},
	'G': {
		"  GGGG ",
		" GG    ",
		"GG     ",
		"GG  GGG",
		"GG   GG",
		"GG   GG",
		" GGGGG ",
	},
	'H': {
		"HH   HH",
		"HH   HH",
		"HH   HH",
		"HHHHHHH",
		"HH   HH",
		"HH   HH",
		"HH   HH",
	},
	'I': {
		"IIIIIII",
		"  III  ",
		"  III  ",
		"  III  ",
		"  III  ",
		"  III  ",
		"IIIIIII",
	},
	'J': {
		"   JJJJ",
		"     JJ",
		"     JJ",
		"     JJ",
		"     JJ",
		"JJ   JJ",
		" JJJJJ ",
	},
	'K': {
		"KK   KK",
		"KK  KK ",
		"KK KK  ",
		"KKKK   ",
		"KK KK  ",
		"KK  KK ",
		"KK   KK",
	},
	'L': {
		"LL     ",
		"LL     ",
		"LL     ",
		"LL     ",
		"LL     ",
		"LL     ",
		"LLLLLLL",
	},
	'M': {
		"MM   MM",
		"MMM MMM",
		"MM M MM",
		"MM   MM",
		"MM   MM",
		"MM   MM",
		"MM   MM",
	},
	'N': {
		"NN   NN",
		"NNN  NN",
		"NNN  NN",
		"NN N NN",
		"NN  NNN",
		"NN  NNN",
		"NN   NN",
	},
	'O': {
		"  OOO  ",
		" OO OO ",
		"OO   OO",
		"OO   OO",
		"OO   OO",
		" OO OO ",
		"  OOO  ",
	},
	'P': {
		"PPPPP  ",
		"PP   PP",
		"PP   PP",
		"PPPPP  ",
		"PP     ",
		"PP     ",
		"PP     ",
	},
	'Q': {
		"  QQQ  ",
		" QQ QQ ",
		"QQ   QQ",
		"QQ   QQ",
		"QQ  QQQ",
		" QQ  QQ",
		"  QQQQQ",
	},
	'R': {
		"RRRRRR ",
		"RR   RR",
		"RR   RR",
		"RRRRRR ",
		"RR  RR ",
		"RR   RR",
		"RR   RR",
	},
	'S': {
		" SSSSS ",
		"SS    S",
		"SS     ",
		" SSSSS ",
		"     SS",
		"S    SS",
		" SSSSS ",
	},
	'T': {
		"TTTTTTT",
		"  TTT  ",
		"  TTT  ",
		"  TTT  ",
		"  TTT  ",
		"  TTT  ",
		"  TTT  ",
	},
	'U': {
		"UU   UU",
		"UU   UU",
		"UU   UU",
		"UU   UU",
		"UU   UU",
		"UU   UU",
		" UUUUU ",
	},
	'V': {
		"VV   VV",
		"VV   VV",
		"VV   VV",
		"VV   VV",
		" VV VV ",
		"  VVV  ",
		"   V   ",
	},
	'W': {
		"WW   WW",
		"WW   WW",
		"WW W WW",
		"WW W WW",
		"WW W WW",
		"WW W WW",
		" WW WW ",
	},
	'X': { 
		"XX   XX",
		" XX XX ",
		"  XXX  ",
		"   X   ",
		"  XXX  ",
		" XX XX ",
		"XX   XX",
	},
	'Y': {
		"YY   YY",
		" YY YY ",
		"  YYYY ",
		"   YY  ",
		"   YY  ",
		"   YY  ",
		"   YY  ",
	},
	'Z': {
		"ZZZZZZZ",
		"     ZZ",
		"    ZZ ",
		"   ZZ  ",
		"  ZZ   ",
		" ZZ    ",
		"ZZZZZZZ",
	},
}
