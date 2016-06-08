
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
    '1. e4 c5 2. Nf3 g6 3. d4 cxd4 4. Qxd4 Nf6 5. e5 Nc6 6. Qa4 Nd5 7. Qe4 Ndb4 8. Bb5 Qa5 9. Nc3 d5 10. exd6 Bf5 11. d7+ Kd8 12. Qc4 Nxc2+ 13. Ke2 Nxa1 14. Rd1 Be6 15. Qxe6 1-0'
]



var chessgame = (function(){

    var game
    var history
    var board
    var cur


    var loadPGN = function(pgn){
	game.load_pgn(pgnData.join('\n'));
	history=game.history({verbose:"true"})
    }


    var prev = function(){
	game.undo()
	cur--
	board.position(game.fen());
    }

    var next = function(){
	game.move(history[cur].san);
	board.position(game.fen());
	cur++;
    }

    var init = function(){
	board = ChessBoard('board1', 'start');
	game = new Chess();

	game.load_pgn(pgnData.join('\n'));
	history=game.history({verbose:"true"})

	game.reset()
	cur=0
    }

    return{
	init:init,
	loadPGN:loadPGN,
	next:next,
	prev:prev
    }
})()







