
/*
  Chessgame module


  Sergio de Mingo

*/



pgnData = [
    '[Event "Euro Club Cup"]',
    '[Site "Kallithea GRE"]',
    '[Date "2008.10.18"]',
    '[EventDate "2008.10.17"]',
    '[Round "2"]',
    '[Result "1-0"]',
    '[White "Simon Ansell"]',
    '[Black "J Garcia-Ortega Mendez"]',
    '[ECO "B27"]',
    '[WhiteElo "2410"]',
    '[BlackElo "2223"]',
    '[PlyCount "29"]',
    '',
    /*'1. e4 c5 2. Nf3 g6 3. d4 cxd4 {bla lbo blo} 4. Qxd4 Nf6',
      '5. e5 Nc6 6. Qa4 Nd5 7. Qe4 Ndb4 8. Bb5 Qa5 9. Nc3 d5 10. exd6 Bf5 11. d7+ Kd8 12. Qc4 Nxc2+ 13. Ke2 Nxa1 14. Rd1 Be6 15. Qxe6 1-0'*/
    '1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}',
    '4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8  10. d4 Nbd7',
    '11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5',
    'Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6',
    '23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5',
    'hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5',
    '35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6',
    'Nf2 42. g4 Bd3 43. Re6 1/2-1/2'
]



var chessgame = (function(){

    var game
    var history
    var board
    var cur
    var headers

    var extractHeaders = function(pgn){
	var headersRe=/\[(.*) \"(.*)\"\]/g
	var h = pgn.replace(headersRe,"\"$1\":\"$2\"")
	var hd = h.match(/\".*\":\".*\"/g)
	if (hd){
	    headers=JSON.parse("{"+hd.join(",")+"}")
	}
    }

    var loadPGN = function(pgn){

	var headersRe=/\[.*\]/g
	pgn = pgn.replace(headersRe,"")

	//alert(pgn)

	var ok = game.load_pgn(pgn,{sloppy: true});
	if (!ok){
	    alert("error pgn")
	    return
	}
	headers=game.header()
	history=game.history({verbose:"true"})

	reset()
    }

    var reset = function(){
	game.reset()
	cur=0
	board.position(game.fen());
    }

    var prev = function(){
	if (cur > 0){
	    game.undo()
	    cur--
	    board.position(game.fen());
	}
    }

    var next = function(){
	if (cur < history.length){
	    game.move(history[cur].san);
	    board.position(game.fen());
	    cur++;
	}
    }

    var init = function(){
	board = ChessBoard('board1', 'start');
	game = new Chess();
    }

    return{
	init:init,
	load:loadPGN,
	next:next,
	prev:prev,
	reset:reset
    }
})()







