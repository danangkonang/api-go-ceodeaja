package kec

import "github.com/danangkonang/rest-api/config"

func Kec6() {
	db := config.Connect()
	db.Exec(`INSERT INTO kecamatan (id,kecamatan,kota_id,provinsi_id) VALUES
		(5001, 'DAMSOL (DAMPELAS SOJOL)', 359, 26),
		(5002, 'LABUAN', 359, 26),
		(5003, 'PINEMBANI', 359, 26),
		(5004, 'RIO PAKAVA (RIOPAKAWA)', 359, 26),
		(5005, 'SINDUE', 359, 26),
		(5006, 'SINDUE TOBATA', 359, 26),
		(5007, 'SINDUE TOMBUSABORA', 359, 26),
		(5008, 'SIRENJA', 359, 26),
		(5009, 'SOJOL', 359, 26),
		(5010, 'SOJOL UTARA', 359, 26),
		(5011, 'TANANTOVEA', 359, 26),
		(5012, 'BAHODOPI', 360, 26),
		(5013, 'BUMI RAYA', 360, 26),
		(5014, 'BUNGKU BARAT', 360, 26),
		(5015, 'BUNGKU PESISIR', 360, 26),
		(5016, 'BUNGKU SELATAN', 360, 26),
		(5017, 'BUNGKU TENGAH', 360, 26),
		(5018, 'BUNGKU TIMUR', 360, 26),
		(5019, 'BUNGKU UTARA', 360, 26),
		(5020, 'LEMBO', 360, 26),
		(5021, 'LEMBO RAYA', 360, 26),
		(5022, 'MAMOSALATO', 360, 26),
		(5023, 'MENUI KEPULAUAN', 360, 26),
		(5024, 'MORI ATAS', 360, 26),
		(5025, 'MORI UTARA', 360, 26),
		(5026, 'PETASIA', 360, 26),
		(5027, 'PETASIA BARAT', 360, 26),
		(5028, 'PETASIA TIMUR', 360, 26),
		(5029, 'SOYO JAYA', 360, 26),
		(5030, 'WITA PONDA', 360, 26),
		(5031, 'MANTIKULORE', 361, 26),
		(5032, 'PALU BARAT', 361, 26),
		(5033, 'PALU SELATAN', 361, 26),
		(5034, 'PALU TIMUR', 361, 26),
		(5035, 'PALU UTARA', 361, 26),
		(5036, 'TATANGA', 361, 26),
		(5037, 'TAWAELI', 361, 26),
		(5038, 'ULUJADI', 361, 26),
		(5039, 'AMPIBABO', 362, 26),
		(5040, 'BALINGGI', 362, 26),
		(5041, 'BOLANO', 362, 26),
		(5042, 'BOLANO LAMBUNU/LAMBULU', 362, 26),
		(5043, 'KASIMBAR', 362, 26),
		(5044, 'MEPANGA', 362, 26),
		(5045, 'MOUTONG', 362, 26),
		(5046, 'ONGKA MALINO', 362, 26),
		(5047, 'PALASA', 362, 26),
		(5048, 'PARIGI', 362, 26),
		(5049, 'PARIGI BARAT', 362, 26),
		(5050, 'PARIGI SELATAN', 362, 26),
		(5051, 'PARIGI TENGAH', 362, 26),
		(5052, 'PARIGI UTARA', 362, 26),
		(5053, 'SAUSU', 362, 26),
		(5054, 'SINIU', 362, 26),
		(5055, 'TAOPA', 362, 26),
		(5056, 'TINOMBO', 362, 26),
		(5057, 'TINOMBO SELATAN', 362, 26),
		(5058, 'TOMINI', 362, 26),
		(5059, 'TORIBULU', 362, 26),
		(5060, 'TORUE', 362, 26),
		(5061, 'LAGE', 363, 26),
		(5062, 'LORE BARAT', 363, 26),
		(5063, 'LORE PIORE', 363, 26),
		(5064, 'LORE SELATAN', 363, 26),
		(5065, 'LORE TENGAH', 363, 26),
		(5066, 'LORE TIMUR', 363, 26),
		(5067, 'LORE UTARA', 363, 26),
		(5068, 'PAMONA BARAT', 363, 26),
		(5069, 'PAMONA PUSELEMBA', 363, 26),
		(5070, 'PAMONA SELATAN', 363, 26),
		(5071, 'PAMONA TENGGARA', 363, 26),
		(5072, 'PAMONA TIMUR', 363, 26),
		(5073, 'PAMONA UTARA', 363, 26),
		(5074, 'POSO KOTA', 363, 26),
		(5075, 'POSO KOTA SELATAN', 363, 26),
		(5076, 'POSO KOTA UTARA', 363, 26),
		(5077, 'POSO PESISIR', 363, 26),
		(5078, 'POSO PESISIR SELATAN', 363, 26),
		(5079, 'POSO PESISIR UTARA', 363, 26),
		(5080, 'DOLO', 364, 26),
		(5081, 'DOLO BARAT', 364, 26),
		(5082, 'DOLO SELATAN', 364, 26),
		(5083, 'GUMBASA', 364, 26),
		(5084, 'KINOVARU', 364, 26),
		(5085, 'KULAWI', 364, 26),
		(5086, 'KULAWI SELATAN', 364, 26),
		(5087, 'LINDU', 364, 26),
		(5088, 'MARAWOLA', 364, 26),
		(5089, 'MARAWOLA BARAT', 364, 26),
		(5090, 'NOKILALAKI', 364, 26),
		(5091, 'PALOLO', 364, 26),
		(5092, 'PIPIKORO', 364, 26),
		(5093, 'SIGI BIROMARU', 364, 26),
		(5094, 'TANAMBULAVA', 364, 26),
		(5095, 'AMPANA KOTA', 365, 26),
		(5096, 'AMPANA TETE', 365, 26),
		(5097, 'TOGEAN', 365, 26),
		(5098, 'TOJO', 365, 26),
		(5099, 'TOJO BARAT', 365, 26),
		(5100, 'ULU BONGKA', 365, 26),
		(5101, 'UNA - UNA', 365, 26),
		(5102, 'WALEA BESAR', 365, 26),
		(5103, 'WALEA KEPULAUAN', 365, 26),
		(5104, 'BAOLAN', 366, 26),
		(5105, 'BASIDONDO', 366, 26),
		(5106, 'DAKO PAMEAN', 366, 26),
		(5107, 'DAMPAL SELATAN', 366, 26),
		(5108, 'DAMPAL UTARA', 366, 26),
		(5109, 'DONDO', 366, 26),
		(5110, 'GALANG', 366, 26),
		(5111, 'LAMPASIO', 366, 26),
		(5112, 'OGO DEIDE', 366, 26),
		(5113, 'TOLITOLI UTARA', 366, 26),
		(5114, 'BOTUMOITA (BOTUMOITO)', 367, 27),
		(5115, 'DULUPI', 367, 27),
		(5116, 'MANANGGU', 367, 27),
		(5117, 'PAGUYAMAN', 367, 27),
		(5118, 'PAGUYAMAN PANTAI', 367, 27),
		(5119, 'TILAMUTA', 367, 27),
		(5120, 'WONOSARI', 367, 27),
		(5121, 'BONE', 368, 27),
		(5122, 'BONE RAYA', 368, 27),
		(5123, 'BONEPANTAI', 368, 27),
		(5124, 'BOTU PINGGE', 368, 27),
		(5125, 'BULANGO SELATAN', 368, 27),
		(5126, 'BULANGO TIMUR', 368, 27),
		(5127, 'BULANGO ULU', 368, 27),
		(5128, 'BULANGO UTARA', 368, 27),
		(5129, 'BULAWA', 368, 27),
		(5130, 'KABILA', 368, 27),
		(5131, 'KABILA BONE', 368, 27),
		(5132, 'PINOGU', 368, 27),
		(5133, 'SUWAWA', 368, 27),
		(5134, 'SUWAWA SELATAN', 368, 27),
		(5135, 'SUWAWA TENGAH', 368, 27),
		(5136, 'SUWAWA TIMUR', 368, 27),
		(5137, 'TAPA', 368, 27),
		(5138, 'TILONGKABILA', 368, 27),
		(5139, 'ASPARAGA', 369, 27),
		(5140, 'BATUDAA', 369, 27),
		(5141, 'BATUDAA PANTAI', 369, 27),
		(5142, 'BILATO', 369, 27),
		(5143, 'BILUHU', 369, 27),
		(5144, 'BOLIOHUTO (BOLIYOHUTO)', 369, 27),
		(5145, 'BONGOMEME', 369, 27),
		(5146, 'DUMBO RAYA', 369, 27),
		(5147, 'DUNGALIYO', 369, 27),
		(5148, 'DUNGINGI', 369, 27),
		(5149, 'HULONTHALANGI', 369, 27),
		(5150, 'KOTA BARAT', 369, 27),
		(5151, 'KOTA SELATAN', 369, 27),
		(5152, 'KOTA TENGAH', 369, 27),
		(5153, 'KOTA TIMUR', 369, 27),
		(5154, 'KOTA UTARA', 369, 27),
		(5155, 'LIMBOTO', 369, 27),
		(5156, 'LIMBOTO BARAT', 369, 27),
		(5157, 'MOOTILANGO', 369, 27),
		(5158, 'PULUBALA', 369, 27),
		(5159, 'SIPATANA', 369, 27),
		(5160, 'TABONGO', 369, 27),
		(5161, 'TELAGA', 369, 27),
		(5162, 'TELAGA BIRU', 369, 27),
		(5163, 'TELAGA JAYA', 369, 27),
		(5164, 'TIBAWA', 369, 27),
		(5165, 'TILANGO', 369, 27),
		(5166, 'TOLANGOHULA', 369, 27),
		(5167, 'ANGGREK', 370, 27),
		(5168, 'ATINGGOLA', 370, 27),
		(5169, 'BIAU', 370, 27),
		(5170, 'GENTUMA RAYA', 370, 27),
		(5171, 'KWANDANG', 370, 27),
		(5172, 'MONANO', 370, 27),
		(5173, 'PONELO KEPULAUAN', 370, 27),
		(5174, 'SUMALATA', 370, 27),
		(5175, 'SUMALATA TIMUR', 370, 27),
		(5176, 'TOLINGGULA', 370, 27),
		(5177, 'TOMOLITO', 370, 27),
		(5178, 'BUNTULIA', 371, 27),
		(5179, 'DENGILO', 371, 27),
		(5180, 'DUHIADAA', 371, 27),
		(5181, 'LEMITO', 371, 27),
		(5182, 'MARISA', 371, 27),
		(5183, 'PAGUAT', 371, 27),
		(5184, 'PATILANGGIO', 371, 27),
		(5185, 'POPAYATO', 371, 27),
		(5186, 'POPAYATO BARAT', 371, 27),
		(5187, 'POPAYATO TIMUR', 371, 27),
		(5188, 'RANDANGAN', 371, 27),
		(5189, 'TALUDITI (TALUDUTI)', 371, 27),
		(5190, 'WANGGARASI', 371, 27),
		(5191, 'AERTEMBAGA (BITUNG TIMUR)', 372, 28),
		(5192, 'GIRIAN', 372, 28),
		(5193, 'LEMBEH SELATAN (BITUNG SELATAN)', 372, 28),
		(5194, 'LEMBEH UTARA', 372, 28),
		(5195, 'MADIDIR (BITUNG TENGAH)', 372, 28),
		(5196, 'MAESA', 372, 28),
		(5197, 'MATUARI (BITUNG BARAT)', 372, 28),
		(5198, 'RANOWULU (BITUNG UTARA)', 372, 28),
		(5199, 'BILALANG', 373, 28),
		(5200, 'BOLAANG', 373, 28),
		(5201, 'BOLAANG TIMUR', 373, 28),
		(5202, 'DUMOGA', 373, 28),
		(5203, 'DUMOGA BARAT', 373, 28),
		(5204, 'DUMOGA TENGAH', 373, 28),
		(5205, 'DUMOGA TENGGARA', 373, 28),
		(5206, 'DUMOGA TIMUR', 373, 28),
		(5207, 'DUMOGA UTARA', 373, 28),
		(5208, 'LOLAK', 373, 28),
		(5209, 'LOLAYAN', 373, 28),
		(5210, 'PASSI BARAT', 373, 28),
		(5211, 'PASSI TIMUR', 373, 28),
		(5212, 'POIGAR', 373, 28),
		(5213, 'SANGTOMBOLANG', 373, 28),
		(5214, 'BOLAANG UKI', 374, 28),
		(5215, 'PINOLOSIAN', 374, 28),
		(5216, 'PINOLOSIAN TENGAH', 374, 28),
		(5217, 'PINOLOSIAN TIMUR', 374, 28),
		(5218, 'POSIGADAN', 374, 28),
		(5219, 'KOTABUNAN', 375, 28),
		(5220, 'MODAYAG', 375, 28),
		(5221, 'MODAYAG BARAT', 375, 28),
		(5222, 'NUANGAN', 375, 28),
		(5223, 'TUTUYAN', 375, 28),
		(5224, 'BINTAUNA', 376, 28),
		(5225, 'BOLANG ITANG BARAT', 376, 28),
		(5226, 'BOLANG ITANG TIMUR', 376, 28),
		(5227, 'KAIDIPANG', 376, 28),
		(5228, 'PINOGALUMAN', 376, 28),
		(5229, 'SANGKUB', 376, 28),
		(5230, 'KENDAHE', 377, 28),
		(5231, 'KEPULAUAN MARORE', 377, 28),
		(5232, 'MANGANITU', 377, 28),
		(5233, 'MANGANITU SELATAN', 377, 28),
		(5234, 'NUSA TABUKAN', 377, 28),
		(5235, 'TABUKAN SELATAN', 377, 28),
		(5236, 'TABUKAN SELATAN TENGAH', 377, 28),
		(5237, 'TABUKAN SELATAN TENGGARA', 377, 28),
		(5238, 'TABUKAN TENGAH', 377, 28),
		(5239, 'TABUKAN UTARA', 377, 28),
		(5240, 'TAHUNA', 377, 28),
		(5241, 'TAHUNA BARAT', 377, 28),
		(5242, 'TAHUNA TIMUR', 377, 28),
		(5243, 'TAMAKO', 377, 28),
		(5244, 'TATOARENG', 377, 28),
		(5245, 'BIARO', 378, 28),
		(5246, 'SIAU BARAT', 378, 28),
		(5247, 'SIAU BARAT SELATAN', 378, 28),
		(5248, 'SIAU BARAT UTARA', 378, 28),
		(5249, 'SIAU TENGAH', 378, 28),
		(5250, 'SIAU TIMUR', 378, 28),
		(5251, 'SIAU TIMUR SELATAN', 378, 28),
		(5252, 'TAGULANDANG', 378, 28),
		(5253, 'TAGULANDANG SELATAN', 378, 28),
		(5254, 'TAGULANDANG UTARA', 378, 28),
		(5255, 'BEO', 379, 28),
		(5256, 'BEO SELATAN', 379, 28),
		(5257, 'BEO UTARA', 379, 28),
		(5258, 'DAMAO (DAMAU)', 379, 28),
		(5259, 'ESSANG', 379, 28),
		(5260, 'ESSANG SELATAN', 379, 28),
		(5261, 'GEMEH', 379, 28),
		(5262, 'KABARUAN', 379, 28),
		(5263, 'KALONGAN', 379, 28),
		(5264, 'LIRUNG', 379, 28),
		(5265, 'MELONGUANE', 379, 28),
		(5266, 'MELONGUANE TIMUR', 379, 28),
		(5267, 'MIANGAS', 379, 28),
		(5268, 'MORONGE', 379, 28),
		(5269, 'NANUSA', 379, 28),
		(5270, 'PULUTAN', 379, 28),
		(5271, 'RAINIS', 379, 28),
		(5272, 'SALIBABU', 379, 28),
		(5273, 'TAMPAN AMMA', 379, 28),
		(5274, 'KOTAMOBAGU BARAT', 380, 28),
		(5275, 'KOTAMOBAGU SELATAN', 380, 28),
		(5276, 'KOTAMOBAGU TIMUR', 380, 28),
		(5277, 'KOTAMOBAGU UTARA', 380, 28),
		(5278, 'BUNAKEN', 381, 28),
		(5279, 'BUNAKEN KEPULAUAN', 381, 28),
		(5280, 'MALALAYANG', 381, 28),
		(5281, 'MAPANGET', 381, 28),
		(5282, 'PAAL DUA', 381, 28),
		(5283, 'SARIO', 381, 28),
		(5284, 'SINGKIL', 381, 28),
		(5285, 'TIKALA', 381, 28),
		(5286, 'TUMINITING', 381, 28),
		(5287, 'WANEA', 381, 28),
		(5288, 'WENANG', 381, 28),
		(5289, 'ERIS', 382, 28),
		(5290, 'KAKAS', 382, 28),
		(5291, 'KAKAS BARAT', 382, 28),
		(5292, 'KAWANGKOAN', 382, 28),
		(5293, 'KAWANGKOAN BARAT', 382, 28),
		(5294, 'KAWANGKOAN UTARA', 382, 28),
		(5295, 'KOMBI', 382, 28),
		(5296, 'LANGOWAN BARAT', 382, 28),
		(5297, 'LANGOWAN SELATAN', 382, 28),
		(5298, 'LANGOWAN TIMUR', 382, 28),
		(5299, 'LANGOWAN UTARA', 382, 28),
		(5300, 'LEMBEAN TIMUR', 382, 28),
		(5301, 'MANDOLANG', 382, 28),
		(5302, 'PINELENG', 382, 28),
		(5303, 'REMBOKEN', 382, 28),
		(5304, 'SONDER', 382, 28),
		(5305, 'TOMBARIRI', 382, 28),
		(5306, 'TOMBARIRI TIMUR', 382, 28),
		(5307, 'TOMBULU', 382, 28),
		(5308, 'TOMPASO', 382, 28),
		(5309, 'TOMPASO BARAT', 382, 28),
		(5310, 'TONDANO BARAT', 382, 28),
		(5311, 'TONDANO SELATAN', 382, 28),
		(5312, 'TONDANO TIMUR', 382, 28),
		(5313, 'TONDANO UTARA', 382, 28),
		(5314, 'AMURANG', 383, 28),
		(5315, 'AMURANG BARAT', 383, 28),
		(5316, 'AMURANG TIMUR', 383, 28),
		(5317, 'KUMELEMBUAI', 383, 28),
		(5318, 'MAESAAN', 383, 28),
		(5319, 'MODOINDING', 383, 28),
		(5320, 'MOTOLING', 383, 28),
		(5321, 'MOTOLING BARAT', 383, 28),
		(5322, 'MOTOLING TIMUR', 383, 28),
		(5323, 'RANOYAPO', 383, 28),
		(5324, 'SINONSAYANG', 383, 28),
		(5325, 'SULUUN TARERAN', 383, 28),
		(5326, 'TARERAN', 383, 28),
		(5327, 'TATAPAAN', 383, 28),
		(5328, 'TENGA', 383, 28),
		(5329, 'TOMPASO BARU', 383, 28),
		(5330, 'TUMPAAN', 383, 28),
		(5331, 'BELANG', 384, 28),
		(5332, 'PASAN', 384, 28),
		(5333, 'PUSOMAEN', 384, 28),
		(5334, 'RATAHAN', 384, 28),
		(5335, 'RATAHAN TIMUR', 384, 28),
		(5336, 'RATATOTOK', 384, 28),
		(5337, 'SILIAN RAYA', 384, 28),
		(5338, 'TOMBATU', 384, 28),
		(5339, 'TOMBATU TIMUR', 384, 28),
		(5340, 'TOMBATU UTARA', 384, 28),
		(5341, 'TOULUAAN', 384, 28),
		(5342, 'TOULUAAN SELATAN', 384, 28),
		(5343, 'AIRMADIDI', 385, 28),
		(5344, 'DIMEMBE', 385, 28),
		(5345, 'KALAWAT', 385, 28),
		(5346, 'KAUDITAN', 385, 28),
		(5347, 'KEMA', 385, 28),
		(5348, 'LIKUPANG BARAT', 385, 28),
		(5349, 'LIKUPANG SELATAN', 385, 28),
		(5350, 'LIKUPANG TIMUR', 385, 28),
		(5351, 'TALAWAAN', 385, 28),
		(5352, 'WORI', 385, 28),
		(5353, 'TOMOHON BARAT', 386, 28),
		(5354, 'TOMOHON SELATAN', 386, 28),
		(5355, 'TOMOHON TENGAH', 386, 28),
		(5356, 'TOMOHON TIMUR', 386, 28),
		(5357, 'TOMOHON UTARA', 386, 28),
		(5358, 'BAGUALA', 387, 29),
		(5359, 'LEITIMUR SELATAN', 387, 29),
		(5360, 'NUSANIWE (NUSANIVE)', 387, 29),
		(5361, 'SIRIMAU', 387, 29),
		(5362, 'TELUK AMBON', 387, 29),
		(5363, 'AIRBUAYA', 388, 29),
		(5364, 'BATABUAL', 388, 29),
		(5365, 'FENA LEISELA', 388, 29),
		(5366, 'LILIALY', 388, 29),
		(5367, 'LOLONG GUBA', 388, 29),
		(5368, 'NAMLEA', 388, 29),
		(5369, 'TELUK KAIELY', 388, 29),
		(5370, 'WAEAPO', 388, 29),
		(5371, 'WAELATA', 388, 29),
		(5372, 'WAPLAU', 388, 29),
		(5373, 'AMBALAU', 389, 29),
		(5374, 'FENA FAFAN', 389, 29),
		(5375, 'KEPALA MADAN', 389, 29),
		(5376, 'LEKSULA', 389, 29),
		(5377, 'NAMROLE', 389, 29),
		(5378, 'WAESAMA', 389, 29),
		(5379, 'ARU SELATAN', 390, 29),
		(5380, 'ARU SELATAN TIMUR', 390, 29),
		(5381, 'ARU SELATAN UTARA', 390, 29),
		(5382, 'ARU TENGAH', 390, 29),
		(5383, 'ARU TENGAH SELATAN', 390, 29),
		(5384, 'ARU TENGAH TIMUR', 390, 29),
		(5385, 'ARU UTARA', 390, 29),
		(5386, 'ARU UTARA TIMUR BATULEY', 390, 29),
		(5387, 'PULAU-PULAU ARU', 390, 29),
		(5388, 'SIR-SIR', 390, 29),
		(5389, 'DAMER', 391, 29),
		(5390, 'DAWELOR DAWERA', 391, 29),
		(5391, 'KEPULAUAN ROMANG', 391, 29),
		(5392, 'KISAR UTARA', 391, 29),
		(5393, 'MDONA HYERA/HIERA', 391, 29),
		(5394, 'MOA LAKOR', 391, 29),
		(5395, 'PULAU LAKOR', 391, 29),
		(5396, 'PULAU LETTI (LETI MOA LAKOR)', 391, 29),
		(5397, 'PULAU MASELA', 391, 29),
		(5398, 'PULAU PULAU BABAR', 391, 29),
		(5399, 'PULAU PULAU TERSELATAN', 391, 29),
		(5400, 'PULAU WETANG', 391, 29),
		(5401, 'PULAU-PULAU BABAR TIMUR', 391, 29),
		(5402, 'WETAR', 391, 29),
		(5403, 'WETAR BARAT', 391, 29),
		(5404, 'WETAR TIMUR', 391, 29),
		(5405, 'WETAR UTARA', 391, 29),
		(5406, 'AMAHAI', 392, 29),
		(5407, 'BANDA', 392, 29),
		(5408, 'LEIHITU', 392, 29),
		(5409, 'LEIHITU BARAT', 392, 29),
		(5410, 'MASOHI KOTA', 392, 29),
		(5411, 'NUSALAUT', 392, 29),
		(5412, 'PULAU HARUKU', 392, 29),
		(5413, 'SALAHUTU', 392, 29),
		(5414, 'SAPARUA', 392, 29),
		(5415, 'SAPARUA TIMUR', 392, 29),
		(5416, 'SERAM UTARA', 392, 29),
		(5417, 'SERAM UTARA BARAT', 392, 29),
		(5418, 'SERAM UTARA TIMUR KOBI', 392, 29),
		(5419, 'SERAM UTARA TIMUR SETI', 392, 29),
		(5420, 'TEHORU', 392, 29),
		(5421, 'TELUK ELPAPUTIH', 392, 29),
		(5422, 'TELUTIH', 392, 29),
		(5423, 'TEON NILA SERUA', 392, 29),
		(5424, 'HOAT SORBAY', 393, 29),
		(5425, 'KEI BESAR', 393, 29),
		(5426, 'KEI BESAR SELATAN', 393, 29),
		(5427, 'KEI BESAR SELATAN BARAT', 393, 29),
		(5428, 'KEI BESAR UTARA BARAT', 393, 29),
		(5429, 'KEI BESAR UTARA TIMUR', 393, 29),
		(5430, 'KEI KECIL', 393, 29),
		(5431, 'KEI KECIL BARAT', 393, 29),
		(5432, 'KEI KECIL TIMUR', 393, 29),
		(5433, 'KEI KECIL TIMUR SELATAN', 393, 29),
		(5434, 'MANYEUW', 393, 29),
		(5435, 'KORMOMOLIN', 394, 29),
		(5436, 'MOLU MARU', 394, 29),
		(5437, 'NIRUNMAS', 394, 29),
		(5438, 'SELARU', 394, 29),
		(5439, 'TANIMBAR SELATAN', 394, 29),
		(5440, 'TANIMBAR UTARA', 394, 29),
		(5441, 'WERMAKATIAN (WER MAKTIAN)', 394, 29),
		(5442, 'WERTAMRIAN', 394, 29),
		(5443, 'WUARLABOBAR', 394, 29),
		(5444, 'YARU', 394, 29),
		(5445, 'AMALATU', 395, 29),
		(5446, 'ELPAPUTIH', 395, 29),
		(5447, 'HUAMUAL', 395, 29),
		(5448, 'HUAMUAL BELAKANG (WAISALA)', 395, 29),
		(5449, 'INAMOSOL', 395, 29),
		(5450, 'KAIRATU', 395, 29),
		(5451, 'KAIRATU BARAT', 395, 29),
		(5452, 'KEPULAUAN MANIPA', 395, 29),
		(5453, 'SERAM BARAT', 395, 29),
		(5454, 'TANIWEL', 395, 29),
		(5455, 'TANIWEL TIMUR', 395, 29),
		(5456, 'BULA', 396, 29),
		(5457, 'BULA BARAT', 396, 29),
		(5458, 'GOROM TIMUR', 396, 29),
		(5459, 'KIAN DARAT', 396, 29),
		(5460, 'KILMURY', 396, 29),
		(5461, 'PULAU GORONG (GOROM)', 396, 29),
		(5462, 'PULAU PANJANG', 396, 29),
		(5463, 'SERAM TIMUR', 396, 29),
		(5464, 'SIRITAUN WIDA TIMUR', 396, 29),
		(5465, 'SIWALALAT', 396, 29),
		(5466, 'TELUK WARU', 396, 29),
		(5467, 'TEOR', 396, 29),
		(5468, 'TUTUK TOLU', 396, 29),
		(5469, 'WAKATE', 396, 29),
		(5470, 'WERINAMA', 396, 29),
		(5471, 'KUR SELATAN', 397, 29),
		(5472, 'PULAU DULLAH SELATAN', 397, 29),
		(5473, 'PULAU DULLAH UTARA', 397, 29),
		(5474, 'PULAU TAYANDO TAM', 397, 29),
		(5475, 'PULAU-PULAU KUR', 397, 29),
		(5476, 'IBU', 398, 30),
		(5477, 'IBU SELATAN', 398, 30),
		(5478, 'IBU UTARA', 398, 30),
		(5479, 'JAILOLO', 398, 30),
		(5480, 'JAILOLO SELATAN', 398, 30),
		(5481, 'LOLODA', 398, 30),
		(5482, 'SAHU', 398, 30),
		(5483, 'SAHU TIMUR', 398, 30),
		(5484, 'BACAN', 399, 30),
		(5485, 'BACAN BARAT', 399, 30),
		(5486, 'BACAN BARAT UTARA', 399, 30),
		(5487, 'BACAN SELATAN', 399, 30),
		(5488, 'BACAN TIMUR', 399, 30),
		(5489, 'BACAN TIMUR SELATAN', 399, 30),
		(5490, 'BACAN TIMUR TENGAH', 399, 30),
		(5491, 'GANE BARAT', 399, 30),
		(5492, 'GANE BARAT SELATAN', 399, 30),
		(5493, 'GANE BARAT UTARA', 399, 30),
		(5494, 'GANE TIMUR', 399, 30),
		(5495, 'GANE TIMUR SELATAN', 399, 30),
		(5496, 'GANE TIMUR TENGAH', 399, 30),
		(5497, 'KASIRUTA BARAT', 399, 30),
		(5498, 'KASIRUTA TIMUR', 399, 30),
		(5499, 'KAYOA', 399, 30),
		(5500, 'KAYOA BARAT', 399, 30),
		(5501, 'KAYOA SELATAN', 399, 30),
		(5502, 'KAYOA UTARA', 399, 30),
		(5503, 'KEPULAUAN BOTANGLOMANG', 399, 30),
		(5504, 'KEPULAUAN JORONGA', 399, 30),
		(5505, 'MAKIAN (PULAU MAKIAN)', 399, 30),
		(5506, 'MAKIAN BARAT (PULAU MAKIAN)', 399, 30),
		(5507, 'MANDIOLI SELATAN', 399, 30),
		(5508, 'MANDIOLI UTARA', 399, 30),
		(5509, 'OBI', 399, 30),
		(5510, 'OBI BARAT', 399, 30),
		(5511, 'OBI SELATAN', 399, 30),
		(5512, 'OBI TIMUR', 399, 30),
		(5513, 'OBI UTARA', 399, 30),
		(5514, 'PATANI', 400, 30),
		(5515, 'PATANI BARAT', 400, 30),
		(5516, 'PATANI UTARA', 400, 30),
		(5517, 'PULAU GEBE', 400, 30),
		(5518, 'WEDA', 400, 30),
		(5519, 'WEDA SELATAN', 400, 30),
		(5520, 'WEDA TENGAH', 400, 30),
		(5521, 'WEDA UTARA', 400, 30),
		(5522, 'KOTA MABA', 401, 30),
		(5523, 'MABA', 401, 30),
		(5524, 'MABA SELATAN', 401, 30),
		(5525, 'MABA TENGAH', 401, 30),
		(5526, 'MABA UTARA', 401, 30),
		(5527, 'WASILE', 401, 30),
		(5528, 'WASILE SELATAN', 401, 30),
		(5529, 'WASILE TENGAH', 401, 30),
		(5530, 'WASILE TIMUR', 401, 30),
		(5531, 'WASILE UTARA', 401, 30),
		(5532, 'GALELA', 402, 30),
		(5533, 'GALELA BARAT', 402, 30),
		(5534, 'GALELA SELATAN', 402, 30),
		(5535, 'GALELA UTARA', 402, 30),
		(5536, 'KAO', 402, 30),
		(5537, 'KAO BARAT', 402, 30),
		(5538, 'KAO TELUK', 402, 30),
		(5539, 'KAO UTARA', 402, 30),
		(5540, 'LOLODA KEPULAUAN', 402, 30),
		(5541, 'LOLODA UTARA', 402, 30),
		(5542, 'MALIFUT', 402, 30),
		(5543, 'TOBELO', 402, 30),
		(5544, 'TOBELO BARAT', 402, 30),
		(5545, 'TOBELO SELATAN', 402, 30),
		(5546, 'TOBELO TENGAH', 402, 30),
		(5547, 'TOBELO TIMUR', 402, 30),
		(5548, 'TOBELO UTARA', 402, 30),
		(5549, 'LEDE', 403, 30),
		(5550, 'MANGOLI BARAT', 403, 30),
		(5551, 'MANGOLI SELATAN', 403, 30),
		(5552, 'MANGOLI TENGAH', 403, 30),
		(5553, 'MANGOLI TIMUR', 403, 30),
		(5554, 'MANGOLI UTARA', 403, 30),
		(5555, 'MANGOLI UTARA TIMUR', 403, 30),
		(5556, 'SANANA', 403, 30),
		(5557, 'SANANA UTARA', 403, 30),
		(5558, 'SULABESI BARAT', 403, 30),
		(5559, 'SULABESI SELATAN', 403, 30),
		(5560, 'SULABESI TENGAH', 403, 30),
		(5561, 'SULABESI TIMUR', 403, 30),
		(5562, 'TALIABU BARAT', 403, 30),
		(5563, 'TALIABU BARAT LAUT', 403, 30),
		(5564, 'TALIABU SELATAN', 403, 30),
		(5565, 'TALIABU TIMUR', 403, 30),
		(5566, 'TALIABU TIMUR SELATAN', 403, 30),
		(5567, 'TALIABU UTARA', 403, 30),
		(5568, 'MOROTAI JAYA', 404, 30),
		(5569, 'MOROTAI SELATAN', 404, 30),
		(5570, 'MOROTAI SELATAN BARAT', 404, 30),
		(5571, 'MOROTAI TIMUR', 404, 30),
		(5572, 'MOROTAI UTARA', 404, 30),
		(5573, 'MOTI (PULAU MOTI)', 405, 30),
		(5574, 'PULAU BATANG DUA', 405, 30),
		(5575, 'PULAU HIRI', 405, 30),
		(5576, 'PULAU TERNATE', 405, 30),
		(5577, 'TERNATE SELATAN (KOTA)', 405, 30),
		(5578, 'TERNATE TENGAH (KOTA)', 405, 30),
		(5579, 'TERNATE UTARA (KOTA)', 405, 30),
		(5580, 'OBA', 406, 30),
		(5581, 'OBA SELATAN', 406, 30),
		(5582, 'OBA TENGAH', 406, 30),
		(5583, 'OBA UTARA', 406, 30),
		(5584, 'TIDORE (PULAU TIDORE)', 406, 30),
		(5585, 'TIDORE SELATAN', 406, 30),
		(5586, 'TIDORE TIMUR (PULAU TIDORE)', 406, 30),
		(5587, 'TIDORE UTARA', 406, 30),
		(5588, 'AMBALAWI', 407, 31),
		(5589, 'ASAKOTA', 407, 31),
		(5590, 'BELO', 407, 31),
		(5591, 'BOLO', 407, 31),
		(5592, 'DONGGO', 407, 31),
		(5593, 'LAMBITU', 407, 31),
		(5594, 'LAMBU', 407, 31),
		(5595, 'LANGGUDU', 407, 31),
		(5596, 'MADAPANGGA', 407, 31),
		(5597, 'MONTA', 407, 31),
		(5598, 'MPUNDA', 407, 31),
		(5599, 'PALIBELO', 407, 31),
		(5600, 'PARADO', 407, 31),
		(5601, 'RABA', 407, 31),
		(5602, 'RASANAE BARAT', 407, 31),
		(5603, 'RASANAE TIMUR', 407, 31),
		(5604, 'SANGGAR', 407, 31),
		(5605, 'SAPE', 407, 31),
		(5606, 'SOROMANDI', 407, 31),
		(5607, 'TAMBORA', 407, 31),
		(5608, 'WERA', 407, 31),
		(5609, 'WOHA', 407, 31),
		(5610, 'DOMPU', 408, 31),
		(5611, 'HU\\U', 408, 31),
		(5612, 'KEMPO', 408, 31),
		(5613, 'KILO', 408, 31),
		(5614, 'MENGGELEWA (MANGGELEWA)', 408, 31),
		(5615, 'PAJO', 408, 31),
		(5616, 'PEKAT', 408, 31),
		(5617, 'WOJA', 408, 31),
		(5618, 'BATU LAYAR', 409, 31),
		(5619, 'GERUNG', 409, 31),
		(5620, 'GUNUNGSARI', 409, 31),
		(5621, 'KEDIRI', 409, 31),
		(5622, 'KURIPAN', 409, 31),
		(5623, 'LABUAPI', 409, 31),
		(5624, 'LEMBAR', 409, 31),
		(5625, 'LINGSAR', 409, 31),
		(5626, 'NARMADA', 409, 31),
		(5627, 'SEKOTONG', 409, 31),
		(5628, 'BATUKLIANG', 410, 31),
		(5629, 'BATUKLIANG UTARA', 410, 31),
		(5630, 'JANAPRIA', 410, 31),
		(5631, 'JONGGAT', 410, 31),
		(5632, 'KOPANG', 410, 31),
		(5633, 'PRAYA', 410, 31),
		(5634, 'PRAYA BARAT', 410, 31),
		(5635, 'PRAYA BARAT DAYA', 410, 31),
		(5636, 'PRAYA TENGAH', 410, 31),
		(5637, 'PRAYA TIMUR', 410, 31),
		(5638, 'PRINGGARATA', 410, 31),
		(5639, 'PUJUT', 410, 31),
		(5640, 'AIKMEL', 411, 31),
		(5641, 'JEROWARU', 411, 31),
		(5642, 'KERUAK', 411, 31),
		(5643, 'LABUHAN HAJI', 411, 31),
		(5644, 'MASBAGIK', 411, 31),
		(5645, 'MONTONG GADING', 411, 31),
		(5646, 'PRINGGABAYA', 411, 31),
		(5647, 'PRINGGASELA', 411, 31),
		(5648, 'SAKRA', 411, 31),
		(5649, 'SAKRA BARAT', 411, 31),
		(5650, 'SAKRA TIMUR', 411, 31),
		(5651, 'SAMBALIA (SAMBELIA)', 411, 31),
		(5652, 'SELONG', 411, 31),
		(5653, 'SEMBALUN', 411, 31),
		(5654, 'SIKUR', 411, 31),
		(5655, 'SUELA (SUWELA)', 411, 31),
		(5656, 'SUKAMULIA', 411, 31),
		(5657, 'SURALAGA', 411, 31),
		(5658, 'TERARA', 411, 31),
		(5659, 'WANASABA', 411, 31),
		(5660, 'BAYAN', 412, 31),
		(5661, 'GANGGA', 412, 31),
		(5662, 'KAYANGAN', 412, 31),
		(5663, 'PEMENANG', 412, 31),
		(5664, 'TANJUNG', 412, 31),
		(5665, 'AMPENAN', 413, 31),
		(5666, 'CAKRANEGARA', 413, 31),
		(5667, 'MATARAM', 413, 31),
		(5668, 'SANDUBAYA (SANDUJAYA)', 413, 31),
		(5669, 'SEKARBELA', 413, 31),
		(5670, 'SELAPARANG (SELAPRANG)', 413, 31),
		(5671, 'ALAS', 414, 31),
		(5672, 'ALAS BARAT', 414, 31),
		(5673, 'BATULANTEH', 414, 31),
		(5674, 'BUER', 414, 31),
		(5675, 'EMPANG', 414, 31),
		(5676, 'LABANGKA', 414, 31),
		(5677, 'LABUHAN BADAS', 414, 31),
		(5678, 'LANTUNG', 414, 31),
		(5679, 'LAPE (LAPE LOPOK)', 414, 31),
		(5680, 'LENANGGUAR', 414, 31),
		(5681, 'LOPOK', 414, 31),
		(5682, 'LUNYUK', 414, 31),
		(5683, 'MARONGE', 414, 31),
		(5684, 'MOYO HILIR', 414, 31),
		(5685, 'MOYO HULU', 414, 31),
		(5686, 'MOYO UTARA', 414, 31),
		(5687, 'ORONG TELU', 414, 31),
		(5688, 'PLAMPANG', 414, 31),
		(5689, 'RHEE', 414, 31),
		(5690, 'ROPANG', 414, 31),
		(5691, 'SUMBAWA', 414, 31),
		(5692, 'TARANO', 414, 31),
		(5693, 'UNTER IWES (UNTERWIRIS)', 414, 31),
		(5694, 'UTAN', 414, 31),
		(5695, 'BRANG ENE', 415, 31),
		(5696, 'BRANG REA', 415, 31),
		(5697, 'JEREWEH', 415, 31),
		(5698, 'MALUK', 415, 31),
		(5699, 'POTO TANO', 415, 31),
		(5700, 'SATELUK (SETELUK)', 415, 31),
		(5701, 'SEKONGKANG', 415, 31),
		(5702, 'TALIWANG', 415, 31),
		(5703, 'ALOR BARAT DAYA', 416, 32),
		(5704, 'ALOR BARAT LAUT', 416, 32),
		(5705, 'ALOR SELATAN', 416, 32),
		(5706, 'ALOR TENGAH UTARA', 416, 32),
		(5707, 'ALOR TIMUR', 416, 32),
		(5708, 'ALOR TIMUR LAUT', 416, 32),
		(5709, 'KABOLA', 416, 32),
		(5710, 'LEMBUR', 416, 32),
		(5711, 'MATARU', 416, 32),
		(5712, 'PANTAR', 416, 32),
		(5713, 'PANTAR BARAT', 416, 32),
		(5714, 'PANTAR BARAT LAUT', 416, 32),
		(5715, 'PANTAR TENGAH', 416, 32),
		(5716, 'PANTAR TIMUR', 416, 32),
		(5717, 'PULAU PURA', 416, 32),
		(5718, 'PUREMAN', 416, 32),
		(5719, 'TELUK MUTIARA', 416, 32),
		(5720, 'ATAMBUA BARAT', 417, 32),
		(5721, 'ATAMBUA KOTA', 417, 32),
		(5722, 'ATAMBUA SELATAN', 417, 32),
		(5723, 'BOTIN LEO BELE', 417, 32),
		(5724, 'IO KUFEU', 417, 32),
		(5725, 'KAKULUK MESAK', 417, 32),
		(5726, 'KOBALIMA', 417, 32),
		(5727, 'KOBALIMA TIMUR', 417, 32),
		(5728, 'LAEN MANEN', 417, 32),
		(5729, 'LAMAKNEN', 417, 32),
		(5730, 'LAMAKNEN SELATAN', 417, 32),
		(5731, 'LASIOLAT', 417, 32),
		(5732, 'MALAKA BARAT', 417, 32),
		(5733, 'MALAKA TENGAH', 417, 32),
		(5734, 'MALAKA TIMUR', 417, 32),
		(5735, 'NANAET DUABESI', 417, 32),
		(5736, 'RAIHAT', 417, 32),
		(5737, 'RAIMANUK', 417, 32),
		(5738, 'RINHAT', 417, 32),
		(5739, 'SASITAMEAN', 417, 32),
		(5740, 'TASIFETO BARAT', 417, 32),
		(5741, 'TASIFETO TIMUR', 417, 32),
		(5742, 'WELIMAN', 417, 32),
		(5743, 'WEWIKU', 417, 32),
		(5744, 'DETUKELI', 418, 32),
		(5745, 'DETUSOKO', 418, 32),
		(5746, 'ENDE', 418, 32),
		(5747, 'ENDE SELATAN', 418, 32),
		(5748, 'ENDE TENGAH', 418, 32),
		(5749, 'ENDE TIMUR', 418, 32),
		(5750, 'ENDE UTARA', 418, 32),
		(5751, 'KELIMUTU', 418, 32),
		(5752, 'KOTABARU', 418, 32),
		(5753, 'LEPEMBUSU KELISOKE', 418, 32),
		(5754, 'LIO TIMUR', 418, 32),
		(5755, 'MAUKARO', 418, 32),
		(5756, 'MAUROLE', 418, 32),
		(5757, 'NANGAPANDA', 418, 32),
		(5758, 'NDONA', 418, 32),
		(5759, 'NDONA TIMUR', 418, 32),
		(5760, 'NDORI', 418, 32),
		(5761, 'PULAU ENDE', 418, 32),
		(5762, 'WEWARIA', 418, 32),
		(5763, 'WOLOJITA', 418, 32),
		(5764, 'WOLOWARU', 418, 32),
		(5765, 'ADONARA', 419, 32),
		(5766, 'ADONARA BARAT', 419, 32),
		(5767, 'ADONARA TENGAH', 419, 32),
		(5768, 'ADONARA TIMUR', 419, 32),
		(5769, 'DEMON PAGONG', 419, 32),
		(5770, 'ILE BOLENG', 419, 32),
		(5771, 'ILE BURA', 419, 32),
		(5772, 'ILE MANDIRI', 419, 32),
		(5773, 'KELUBAGOLIT (KLUBAGOLIT)', 419, 32),
		(5774, 'LARANTUKA', 419, 32),
		(5775, 'LEWOLEMA', 419, 32),
		(5776, 'SOLOR BARAT', 419, 32),
		(5777, 'SOLOR SELATAN', 419, 32),
		(5778, 'SOLOR TIMUR', 419, 32),
		(5779, 'TANJUNG BUNGA', 419, 32),
		(5780, 'TITEHENA', 419, 32),
		(5781, 'WITIHAMA (WATIHAMA)', 419, 32),
		(5782, 'WOTAN ULUMADO', 419, 32),
		(5783, 'WULANGGITANG', 419, 32),
		(5784, 'ALAK', 420, 32),
		(5785, 'AMABI OEFETO', 420, 32),
		(5786, 'AMABI OEFETO TIMUR', 420, 32),
		(5787, 'AMARASI', 420, 32),
		(5788, 'AMARASI BARAT', 420, 32),
		(5789, 'AMARASI SELATAN', 420, 32),
		(5790, 'AMARASI TIMUR', 420, 32),
		(5791, 'AMFOANG BARAT DAYA', 420, 32),
		(5792, 'AMFOANG BARAT LAUT', 420, 32),
		(5793, 'AMFOANG SELATAN', 420, 32),
		(5794, 'AMFOANG TENGAH', 420, 32),
		(5795, 'AMFOANG TIMUR', 420, 32),
		(5796, 'AMFOANG UTARA', 420, 32),
		(5797, 'FATULEU', 420, 32),
		(5798, 'FATULEU BARAT', 420, 32),
		(5799, 'FATULEU TENGAH', 420, 32),
		(5800, 'KELAPA LIMA', 420, 32),
		(5801, 'KOTA LAMA', 420, 32),
		(5802, 'KOTA RAJA', 420, 32),
		(5803, 'KUPANG BARAT', 420, 32),
		(5804, 'KUPANG TENGAH', 420, 32),
		(5805, 'KUPANG TIMUR', 420, 32),
		(5806, 'MAULAFA', 420, 32),
		(5807, 'NEKAMESE', 420, 32),
		(5808, 'OEBOBO', 420, 32),
		(5809, 'SEMAU', 420, 32),
		(5810, 'SEMAU SELATAN', 420, 32),
		(5811, 'SULAMU', 420, 32),
		(5812, 'TAEBENU', 420, 32),
		(5813, 'TAKARI', 420, 32),
		(5814, 'ATADEI', 421, 32),
		(5815, 'BUYASURI (BUYASARI)', 421, 32),
		(5816, 'ILE APE', 421, 32),
		(5817, 'ILE APE TIMUR', 421, 32),
		(5818, 'LEBATUKAN', 421, 32),
		(5819, 'NAGAWUTUNG', 421, 32),
		(5820, 'NUBATUKAN', 421, 32),
		(5821, 'OMESURI', 421, 32),
		(5822, 'WULANDONI (WULANDIONI)', 421, 32),
		(5823, 'CIBAL', 422, 32),
		(5824, 'CIBAL BARAT', 422, 32),
		(5825, 'LANGKE REMBONG', 422, 32),
		(5826, 'LELAK', 422, 32),
		(5827, 'RAHONG UTARA', 422, 32),
		(5828, 'REOK', 422, 32),
		(5829, 'REOK BARAT', 422, 32),
		(5830, 'RUTENG', 422, 32),
		(5831, 'SATAR MESE', 422, 32),
		(5832, 'SATAR MESE BARAT', 422, 32),
		(5833, 'WAE RII', 422, 32),
		(5834, 'BOLENG', 423, 32),
		(5835, 'KOMODO', 423, 32),
		(5836, 'KUWUS', 423, 32),
		(5837, 'LEMBOR', 423, 32),
		(5838, 'LEMBOR SELATAN', 423, 32),
		(5839, 'MACANG PACAR', 423, 32),
		(5840, 'MBELILING', 423, 32),
		(5841, 'NDOSO', 423, 32),
		(5842, 'SANO NGGOANG', 423, 32),
		(5843, 'WELAK', 423, 32),
		(5844, 'BORONG', 424, 32),
		(5845, 'ELAR', 424, 32),
		(5846, 'ELAR SELATAN', 424, 32),
		(5847, 'KOTA KOMBA', 424, 32),
		(5848, 'LAMBA LEDA', 424, 32),
		(5849, 'POCO RANAKA', 424, 32),
		(5850, 'POCO RANAKA TIMUR', 424, 32),
		(5851, 'RANA MESE', 424, 32),
		(5852, 'SAMBI RAMPAS', 424, 32),
		(5853, 'AESESA', 425, 32),
		(5854, 'AESESA SELATAN', 425, 32),
		(5855, 'BOAWAE', 425, 32),
		(5856, 'KEO TENGAH', 425, 32),
		(5857, 'MAUPONGGO', 425, 32),
		(5858, 'NANGARORO', 425, 32),
		(5859, 'WOLOWAE', 425, 32),
		(5860, 'AIMERE', 426, 32),
		(5861, 'BAJAWA', 426, 32),
		(5862, 'BAJAWA UTARA', 426, 32),
		(5863, 'GOLEWA', 426, 32),
		(5864, 'GOLEWA BARAT', 426, 32),
		(5865, 'GOLEWA SELATAN', 426, 32),
		(5866, 'INERIE', 426, 32),
		(5867, 'JEREBUU', 426, 32),
		(5868, 'RIUNG', 426, 32),
		(5869, 'RIUNG BARAT', 426, 32),
		(5870, 'SOA', 426, 32),
		(5871, 'WOLOMEZE (RIUNG SELATAN)', 426, 32),
		(5872, 'LANDU LEKO', 427, 32),
		(5873, 'LOBALAIN', 427, 32),
		(5874, 'NDAO NUSE', 427, 32),
		(5875, 'PANTAI BARU', 427, 32),
		(5876, 'ROTE BARAT', 427, 32),
		(5877, 'ROTE BARAT DAYA', 427, 32),
		(5878, 'ROTE BARAT LAUT', 427, 32),
		(5879, 'ROTE SELATAN', 427, 32),
		(5880, 'ROTE TENGAH', 427, 32),
		(5881, 'ROTE TIMUR', 427, 32),
		(5882, 'HAWU MEHARA', 428, 32),
		(5883, 'RAIJUA', 428, 32),
		(5884, 'SABU BARAT', 428, 32),
		(5885, 'SABU LIAE', 428, 32),
		(5886, 'SABU TENGAH', 428, 32),
		(5887, 'SABU TIMUR', 428, 32),
		(5888, 'ALOK', 429, 32),
		(5889, 'ALOK BARAT', 429, 32),
		(5890, 'ALOK TIMUR', 429, 32),
		(5891, 'DORENG', 429, 32),
		(5892, 'HEWOKLOANG', 429, 32),
		(5893, 'KANGAE', 429, 32),
		(5894, 'KEWAPANTE', 429, 32),
		(5895, 'KOTING', 429, 32),
		(5896, 'LELA', 429, 32),
		(5897, 'MAGEPANDA', 429, 32),
		(5898, 'MAPITARA', 429, 32),
		(5899, 'MEGO', 429, 32),
		(5900, 'NELLE (MAUMEREI)', 429, 32),
		(5901, 'NITA', 429, 32),
		(5902, 'PAGA', 429, 32),
		(5903, 'PALUE', 429, 32),
		(5904, 'TALIBURA', 429, 32),
		(5905, 'TANA WAWO', 429, 32),
		(5906, 'WAIBLAMA', 429, 32),
		(5907, 'WAIGETE', 429, 32),
		(5908, 'KOTA WAIKABUBAK', 430, 32),
		(5909, 'LAMBOYA', 430, 32),
		(5910, 'LAMBOYA BARAT', 430, 32),
		(5911, 'LOLI', 430, 32),
		(5912, 'TANA RIGHU', 430, 32),
		(5913, 'WANOKAKA', 430, 32),
		(5914, 'KODI', 431, 32),
		(5915, 'KODI BALAGHAR', 431, 32),
		(5916, 'KODI BANGEDO', 431, 32),
		(5917, 'KODI UTARA', 431, 32),
		(5918, 'KOTA TAMBOLAKA', 431, 32),
		(5919, 'LOURA (LAURA)', 431, 32),
		(5920, 'WEWEWA BARAT', 431, 32),
		(5921, 'WEWEWA SELATAN', 431, 32),
		(5922, 'WEWEWA TENGAH (WEWERA TENGAH)', 431, 32),
		(5923, 'WEWEWA TIMUR', 431, 32),
		(5924, 'WEWEWA UTARA', 431, 32),
		(5925, 'KATIKUTANA', 432, 32),
		(5926, 'KATIKUTANA SELATAN', 432, 32),
		(5927, 'MAMBORO', 432, 32),
		(5928, 'UMBU RATU NGGAY', 432, 32),
		(5929, 'UMBU RATU NGGAY BARAT', 432, 32),
		(5930, 'HAHARU', 433, 32),
		(5931, 'KAHAUNGUWETI (KAHAUNGU ETI)', 433, 32),
		(5932, 'KAMBATA MAPAMBUHANG', 433, 32),
		(5933, 'KAMBERA', 433, 32),
		(5934, 'KANATANG', 433, 32),
		(5935, 'KARERA', 433, 32),
		(5936, 'KATALA HAMU LINGU', 433, 32),
		(5937, 'KOTA WAINGAPU', 433, 32),
		(5938, 'LEWA', 433, 32),
		(5939, 'LEWA TIDAHU', 433, 32),
		(5940, 'MAHU', 433, 32),
		(5941, 'MATAWAI LAPPAU (LA PAWU)', 433, 32),
		(5942, 'NGADU NGALA', 433, 32),
		(5943, 'NGGAHA ORIANGU', 433, 32),
		(5944, 'PABERIWAI', 433, 32),
		(5945, 'PAHUNGA LODU', 433, 32),
		(5946, 'PANDAWAI', 433, 32),
		(5947, 'PINUPAHAR (PIRAPAHAR)', 433, 32),
		(5948, 'RINDI', 433, 32),
		(5949, 'TABUNDUNG', 433, 32),
		(5950, 'UMALULU', 433, 32),
		(5951, 'WULA WAIJELU', 433, 32),
		(5952, 'AMANATUN SELATAN', 434, 32),
		(5953, 'AMANATUN UTARA', 434, 32),
		(5954, 'AMANUBAN BARAT', 434, 32),
		(5955, 'AMANUBAN SELATAN', 434, 32),
		(5956, 'AMANUBAN TENGAH', 434, 32),
		(5957, 'AMANUBAN TIMUR', 434, 32),
		(5958, 'BOKING', 434, 32),
		(5959, 'FATUKOPA', 434, 32),
		(5960, 'FATUMNASI', 434, 32),
		(5961, 'FAUTMOLO', 434, 32),
		(5962, 'KIE (KI\\E)', 434, 32),
		(5963, 'KOK BAUN', 434, 32),
		(5964, 'KOLBANO', 434, 32),
		(5965, 'KOT OLIN', 434, 32),
		(5966, 'KOTA SOE', 434, 32),
		(5967, 'KUALIN', 434, 32),
		(5968, 'KUANFATU', 434, 32),
		(5969, 'KUATNANA', 434, 32),
		(5970, 'MOLLO BARAT', 434, 32),
		(5971, 'MOLLO SELATAN', 434, 32),
		(5972, 'MOLLO TENGAH', 434, 32),
		(5973, 'MOLLO UTARA', 434, 32),
		(5974, 'NOEBANA', 434, 32),
		(5975, 'NOEBEBA', 434, 32),
		(5976, 'NUNBENA', 434, 32),
		(5977, 'NUNKOLO', 434, 32),
		(5978, 'OENINO', 434, 32),
		(5979, 'POLEN', 434, 32),
		(5980, 'SANTIAN', 434, 32),
		(5981, 'TOBU', 434, 32),
		(5982, 'TOIANAS', 434, 32),
		(5983, 'BIBOKI ANLEU', 435, 32),
		(5984, 'BIBOKI FEOTLEU', 435, 32),
		(5985, 'BIBOKI MOENLEU', 435, 32),
		(5986, 'BIBOKI SELATAN', 435, 32),
		(5987, 'BIBOKI TAN PAH', 435, 32),
		(5988, 'BIBOKI UTARA', 435, 32),
		(5989, 'BIKOMI NILULAT', 435, 32),
		(5990, 'BIKOMI SELATAN', 435, 32),
		(5991, 'BIKOMI TENGAH', 435, 32),
		(5992, 'BIKOMI UTARA', 435, 32),
		(5993, 'INSANA', 435, 32),
		(5994, 'INSANA BARAT', 435, 32),
		(5995, 'INSANA FAFINESU', 435, 32),
		(5996, 'INSANA TENGAH', 435, 32),
		(5997, 'INSANA UTARA', 435, 32),
		(5998, 'KOTA KEFAMENANU', 435, 32),
		(5999, 'MIOMAFFO BARAT', 435, 32),
		(6000, 'MIOMAFFO TENGAH', 435, 32)
	`)
}
