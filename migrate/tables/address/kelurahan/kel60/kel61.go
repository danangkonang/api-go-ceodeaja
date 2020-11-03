package kel60

import "github.com/danangkonang/ceodeaja-go/config"

func Kel61() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
	(61001, 'BOKIN', 4719, 342, 24),
	(61002, 'BUANGIN', 4719, 342, 24),
	(61003, 'MAKUAN PARE', 4719, 342, 24),
	(61004, 'PITUNG PENANIAN', 4719, 342, 24),
	(61005, 'RANTEBUA', 4719, 342, 24),
	(61006, 'RANTEBUA SANGGALAGI', 4719, 342, 24),
	(61007, 'RANTEBUA SUMALU', 4719, 342, 24),
	(61008, 'KARASSIK', 4720, 342, 24),
	(61009, 'LAANG TANDUK', 4720, 342, 24),
	(61010, 'LIMBONG', 4720, 342, 24),
	(61011, 'MALANGO', 4720, 342, 24),
	(61012, 'MENTIRO TIKU', 4720, 342, 24),
	(61013, 'PASELE', 4720, 342, 24),
	(61014, 'PENANIAN', 4720, 342, 24),
	(61015, 'RANTE PASELE', 4720, 342, 24),
	(61016, 'RANTEPAO', 4720, 342, 24),
	(61017, 'SALOSO', 4720, 342, 24),
	(61018, 'SINGKI', 4720, 342, 24),
	(61019, 'AMPANG BATU (APPANG)', 4721, 342, 24),
	(61020, 'BUNTU BATU', 4721, 342, 24),
	(61021, 'LEMPOPOTON', 4721, 342, 24),
	(61022, 'LIMBONG MALTING', 4721, 342, 24),
	(61023, 'LO KO URU TANETE BATU', 4721, 342, 24),
	(61024, 'MAITING', 4721, 342, 24),
	(61025, 'PANGALA', 4721, 342, 24),
	(61026, 'PANGALA UTARA', 4721, 342, 24),
	(61027, 'RINDINGALLO', 4721, 342, 24),
	(61028, 'SA\\ DAN ANDULAN', 4722, 342, 24),
	(61029, 'SA\\ DAN BALLOPASANGE', 4722, 342, 24),
	(61030, 'SA\\ DAN LIKU LAMBE', 4722, 342, 24),
	(61031, 'SA\\ DAN MALIMBONG', 4722, 342, 24),
	(61032, 'SA\\ DAN MATALLO', 4722, 342, 24),
	(61033, 'SA\\ DAN PEBULIAN', 4722, 342, 24),
	(61034, 'SA\\ DAN PESONDONGAN', 4722, 342, 24),
	(61035, 'SA\\ DAN SANGKAROPI', 4722, 342, 24),
	(61036, 'SA\\ DAN TIROALLO', 4722, 342, 24),
	(61037, 'SA\\ DAN ULUSALU', 4722, 342, 24),
	(61038, 'BUNTU LA\\BO', 4723, 342, 24),
	(61039, 'LA\\BO', 4723, 342, 24),
	(61040, 'PAEPALEAN', 4723, 342, 24),
	(61041, 'PATA\\ PADANG', 4723, 342, 24),
	(61042, 'TALLUNG PENANIAN', 4723, 342, 24),
	(61043, 'TANDUNG LA\\BO\\', 4723, 342, 24),
	(61044, 'BORI', 4724, 342, 24),
	(61045, 'BORI LOMBONGAN', 4724, 342, 24),
	(61046, 'BORI RANTE LETOK', 4724, 342, 24),
	(61047, 'BUNTULOBO', 4724, 342, 24),
	(61048, 'DERI', 4724, 342, 24),
	(61049, 'PALAWA', 4724, 342, 24),
	(61050, 'PANGLI PALAWA', 4724, 342, 24),
	(61051, 'PANGLI SELATAN', 4724, 342, 24),
	(61052, 'PARINDING', 4724, 342, 24),
	(61053, 'LANDORUNDUN', 4725, 342, 24),
	(61054, 'LEMPO', 4725, 342, 24),
	(61055, 'SESEAN MATALLO', 4725, 342, 24),
	(61056, 'SULOARA', 4725, 342, 24),
	(61057, 'TONGA RIU', 4725, 342, 24),
	(61058, 'LANGDA', 4726, 342, 24),
	(61059, 'MARANTE', 4726, 342, 24),
	(61060, 'NONONGAN SELATAN', 4726, 342, 24),
	(61061, 'NONONGAN UTARA', 4726, 342, 24),
	(61062, 'SALU', 4726, 342, 24),
	(61063, 'SALU RANTE', 4726, 342, 24),
	(61064, 'SALU SOPAI', 4726, 342, 24),
	(61065, 'TOMBANG LANGDA', 4726, 342, 24),
	(61066, 'BUNTU TALLUNGLIPU', 4727, 342, 24),
	(61067, 'RANTEPAKU TALLUNGLIPU', 4727, 342, 24),
	(61068, 'TAGARI TALLUNGLIPU', 4727, 342, 24),
	(61069, 'TALLUNGLIPU', 4727, 342, 24),
	(61070, 'TALLUNGLIPU MATALLO', 4727, 342, 24),
	(61071, 'TAMPO TALLUNGLIPU', 4727, 342, 24),
	(61072, 'TANTANAN TALLUNGLIPU', 4727, 342, 24),
	(61073, 'TONDON', 4728, 342, 24),
	(61074, 'TONDON LANGI', 4728, 342, 24),
	(61075, 'TONDON MATALLO', 4728, 342, 24),
	(61076, 'TONDON SIBA TA', 4728, 342, 24),
	(61077, 'BELAWA', 4729, 343, 24),
	(61078, 'LAUTANG', 4729, 343, 24),
	(61079, 'LEPPANGENG', 4729, 343, 24),
	(61080, 'LIMPO RILAU', 4729, 343, 24),
	(61081, 'MACERO', 4729, 343, 24),
	(61082, 'MALANGKE (MALAKKE)', 4729, 343, 24),
	(61083, 'ONGKOE', 4729, 343, 24),
	(61084, 'SAPPA', 4729, 343, 24),
	(61085, 'WELE', 4729, 343, 24),
	(61086, 'BALIELO', 4730, 343, 24),
	(61087, 'BOLA', 4730, 343, 24),
	(61088, 'BOLA', 4730, 343, 24),
	(61089, 'HOKOR', 4730, 343, 24),
	(61090, 'IPIR', 4730, 343, 24),
	(61091, 'LATTIMU', 4730, 343, 24),
	(61092, 'LEMPONG', 4730, 343, 24),
	(61093, 'MANURUNG', 4730, 343, 24),
	(61094, 'PASIR PUTIH', 4730, 343, 24),
	(61095, 'PATTANGNGAE (PATANGAE)', 4730, 343, 24),
	(61096, 'RAJAMAWELANG (RAJAMAWELLANG)', 4730, 343, 24),
	(61097, 'SANRESENG ADE', 4730, 343, 24),
	(61098, 'SOLO', 4730, 343, 24),
	(61099, 'UJUNG TANAH', 4730, 343, 24),
	(61100, 'UMAUTA', 4730, 343, 24),
	(61101, 'WOLOKOLI', 4730, 343, 24),
	(61102, 'WOLONWALU', 4730, 343, 24),
	(61103, 'ABBATIRENG', 4731, 343, 24),
	(61104, 'ALAUSALO', 4731, 343, 24),
	(61105, 'ARAJANG', 4731, 343, 24),
	(61106, 'GILIRENG', 4731, 343, 24),
	(61107, 'LAMATA', 4731, 343, 24),
	(61108, 'MAMMINASAE (MAMINASAE)', 4731, 343, 24),
	(61109, 'PASELLORENG', 4731, 343, 24),
	(61110, 'POLEONRO', 4731, 343, 24),
	(61111, 'POLEWALIE', 4731, 343, 24),
	(61112, 'AWO', 4732, 343, 24),
	(61113, 'AWOTA', 4732, 343, 24),
	(61114, 'BALLAERE (BALLERE)', 4732, 343, 24),
	(61115, 'CIROMANIE', 4732, 343, 24),
	(61116, 'INRELLO', 4732, 343, 24),
	(61117, 'KEERA', 4732, 343, 24),
	(61118, 'LABAWANG', 4732, 343, 24),
	(61119, 'LALLISENG', 4732, 343, 24),
	(61120, 'PAOJEPE', 4732, 343, 24),
	(61121, 'PATTIROLOKKA', 4732, 343, 24),
	(61122, 'BOTTO BENTENG', 4733, 343, 24),
	(61123, 'BOTTO PENNO', 4733, 343, 24),
	(61124, 'BOTTO TANRE', 4733, 343, 24),
	(61125, 'CINNONG TABI', 4733, 343, 24),
	(61126, 'LAERUNG', 4733, 343, 24),
	(61127, 'LAMIKU', 4733, 343, 24),
	(61128, 'LIMPO MAJANG', 4733, 343, 24),
	(61129, 'LIU', 4733, 343, 24),
	(61130, 'MACANANG', 4733, 343, 24),
	(61131, 'PARIA', 4733, 343, 24),
	(61132, 'RUMPIA', 4733, 343, 24),
	(61133, 'TAJO', 4733, 343, 24),
	(61134, 'TELLU LIMPOE', 4733, 343, 24),
	(61135, 'TENGNGA', 4733, 343, 24),
	(61136, 'TOSORA', 4733, 343, 24),
	(61137, 'TUA', 4733, 343, 24),
	(61138, 'URAIANG', 4733, 343, 24),
	(61139, 'WATAN RUMPIA', 4733, 343, 24),
	(61140, 'ABBANUANGE', 4734, 343, 24),
	(61141, 'ANABANUA', 4734, 343, 24),
	(61142, 'DUA LIMPOE', 4734, 343, 24),
	(61143, 'KALOLA', 4734, 343, 24),
	(61144, 'MATTIROWALIE', 4734, 343, 24),
	(61145, 'MINANGA TELLUE', 4734, 343, 24),
	(61146, 'SOGI', 4734, 343, 24),
	(61147, 'TANGKOLI', 4734, 343, 24),
	(61148, 'ABBANUANGE', 4735, 343, 24),
	(61149, 'CINA', 4735, 343, 24),
	(61150, 'KAMPIRI', 4735, 343, 24),
	(61151, 'LAGOSI', 4735, 343, 24),
	(61152, 'LAMPULUNG', 4735, 343, 24),
	(61153, 'LAPAUKKE', 4735, 343, 24),
	(61154, 'LEMPA', 4735, 343, 24),
	(61155, 'PALLAWARUKKA', 4735, 343, 24),
	(61156, 'PAMMANA', 4735, 343, 24),
	(61157, 'PATILA', 4735, 343, 24),
	(61158, 'SIMPURSIA (SIMPURUSIA)', 4735, 343, 24),
	(61159, 'TADANG PALIE', 4735, 343, 24),
	(61160, 'TOBATANG', 4735, 343, 24),
	(61161, 'WATAMPANUA', 4735, 343, 24),
	(61162, 'WECUDAI', 4735, 343, 24),
	(61163, 'BENTENG', 4736, 343, 24),
	(61164, 'DOPING', 4736, 343, 24),
	(61165, 'LAWESSO', 4736, 343, 24),
	(61166, 'MAKMUR', 4736, 343, 24),
	(61167, 'PADAELO', 4736, 343, 24),
	(61168, 'PENRANG', 4736, 343, 24),
	(61169, 'RADDAE', 4736, 343, 24),
	(61170, 'TADANG PALIE', 4736, 343, 24),
	(61171, 'TEMMABARANG', 4736, 343, 24),
	(61172, 'WALANGA', 4736, 343, 24),
	(61173, 'ABBANDERANGNGE', 4737, 343, 24),
	(61174, 'ALESILURUNGE', 4737, 343, 24),
	(61175, 'BATU', 4737, 343, 24),
	(61176, 'BENTENG', 4737, 343, 24),
	(61177, 'BULETE', 4737, 343, 24),
	(61178, 'LAUWA', 4737, 343, 24),
	(61179, 'LOMPOLOANG', 4737, 343, 24),
	(61180, 'MARANNU', 4737, 343, 24),
	(61181, 'SIMPELLU', 4737, 343, 24),
	(61182, 'SIWA', 4737, 343, 24),
	(61183, 'TANGKORO', 4737, 343, 24),
	(61184, 'TANRONGI', 4737, 343, 24),
	(61185, 'TELLESANG', 4737, 343, 24),
	(61186, 'TOBARAKKA', 4737, 343, 24),
	(61187, 'BENTENG LOMPOE', 4738, 343, 24),
	(61188, 'BILA', 4738, 343, 24),
	(61189, 'LIU', 4738, 343, 24),
	(61190, 'MALLUSESALO', 4738, 343, 24),
	(61191, 'PALLIMAE', 4738, 343, 24),
	(61192, 'PASAKA', 4738, 343, 24),
	(61193, 'SALOTENGNGA', 4738, 343, 24),
	(61194, 'SOMPE', 4738, 343, 24),
	(61195, 'TADANGPALIE', 4738, 343, 24),
	(61196, 'TOLOTENRENG (TALOTENRENG)', 4738, 343, 24),
	(61197, 'UGI', 4738, 343, 24),
	(61198, 'UJUNGPERO', 4738, 343, 24),
	(61199, 'WAGE', 4738, 343, 24),
	(61200, 'WALENNAE (WALENAE)', 4738, 343, 24),
	(61201, 'WORONGNGE', 4738, 343, 24),
	(61202, 'AKKAJENG', 4739, 343, 24),
	(61203, 'AKKOTENGENG', 4739, 343, 24),
	(61204, 'ALEWADENG', 4739, 343, 24),
	(61205, 'ASSORAJANGE', 4739, 343, 24),
	(61206, 'BARANGMAMASE', 4739, 343, 24),
	(61207, 'MINANGAE', 4739, 343, 24),
	(61208, 'SAKKOLI', 4739, 343, 24),
	(61209, 'SALOBULO', 4739, 343, 24),
	(61210, 'TOWALIDA', 4739, 343, 24),
	(61211, 'AJURAJA', 4740, 343, 24),
	(61212, 'ALUPPANG', 4740, 343, 24),
	(61213, 'BOCCO', 4740, 343, 24),
	(61214, 'BOTTO', 4740, 343, 24),
	(61215, 'CEPPAGA', 4740, 343, 24),
	(61216, 'LAGOARI', 4740, 343, 24),
	(61217, 'LAMARUA', 4740, 343, 24),
	(61218, 'LEWENG', 4740, 343, 24),
	(61219, 'MANYILI', 4740, 343, 24),
	(61220, 'PANTAI TIMUR', 4740, 343, 24),
	(61221, 'PARIGI', 4740, 343, 24),
	(61222, 'PENEKI', 4740, 343, 24),
	(61223, 'SORO', 4740, 343, 24),
	(61224, 'ASSORAJANG', 4741, 343, 24),
	(61225, 'BARU TANCUNG', 4741, 343, 24),
	(61226, 'INALIPUE', 4741, 343, 24),
	(61227, 'LOWA', 4741, 343, 24),
	(61228, 'MANNAGAE', 4741, 343, 24),
	(61229, 'MAPPADAELO', 4741, 343, 24),
	(61230, 'MARIO', 4741, 343, 24),
	(61231, 'NEPO', 4741, 343, 24),
	(61232, 'PAJALELE', 4741, 343, 24),
	(61233, 'PAKKANNA', 4741, 343, 24),
	(61234, 'PALIPPU', 4741, 343, 24),
	(61235, 'PINCENGPUTE', 4741, 343, 24),
	(61236, 'TANCUNG', 4741, 343, 24),
	(61237, 'TONRALIPUE', 4741, 343, 24),
	(61238, 'UJUNG BARU', 4741, 343, 24),
	(61239, 'UJUNGE', 4741, 343, 24),
	(61240, 'WAETUWO', 4741, 343, 24),
	(61241, 'WAJORIAJA', 4741, 343, 24),
	(61242, 'WEWANGREWU', 4741, 343, 24),
	(61243, 'ATAKKAE', 4742, 343, 24),
	(61244, 'BULUPABBULU', 4742, 343, 24),
	(61245, 'CEMPALAGI', 4742, 343, 24),
	(61246, 'LAELO', 4742, 343, 24),
	(61247, 'LAPONGKODA', 4742, 343, 24),
	(61248, 'MADDUKKELLENG', 4742, 343, 24),
	(61249, 'MATTIROTAPPARENG', 4742, 343, 24),
	(61250, 'PADDUPPA', 4742, 343, 24),
	(61251, 'PATTIROSOMPE', 4742, 343, 24),
	(61252, 'SALOMENRALENG', 4742, 343, 24),
	(61253, 'SIENGKANG', 4742, 343, 24),
	(61254, 'SITAMPAE', 4742, 343, 24),
	(61255, 'TEDDAOPU', 4742, 343, 24),
	(61256, 'TEMPE', 4742, 343, 24),
	(61257, 'WATALLIPUE', 4742, 343, 24),
	(61258, 'WIRINGPALENNAE', 4742, 343, 24),



	
	(61259, 'BONE-BONE', 4743, 344, 25),
	(61260, 'KAOBULA', 4743, 344, 25),
	(61261, 'LANTO', 4743, 344, 25),
	(61262, 'NGANGANAUMALA', 4743, 344, 25),
	(61263, 'TARAFU', 4743, 344, 25),
	(61264, 'WAMEO', 4743, 344, 25),
	(61265, 'KATOBENGKE', 4744, 344, 25),
	(61266, 'LABALAWA', 4744, 344, 25),
	(61267, 'LIPU', 4744, 344, 25),
	(61268, 'SULAA', 4744, 344, 25),
	(61269, 'WABOROBO', 4744, 344, 25),
	(61270, 'KAMPEONAHO', 4745, 344, 25),
	(61271, 'LIA BUKU', 4745, 344, 25),
	(61272, 'NGKARING-KARING (NGKARI-NGKARI)', 4745, 344, 25),
	(61273, 'TAMPUNA', 4745, 344, 25),
	(61274, 'WALIA BUKU', 4745, 344, 25),
	(61275, 'KADOLO', 4746, 344, 25),
	(61276, 'KADOLOMOKO', 4746, 344, 25),
	(61277, 'LAKOLOGOU (LAKOLOGUO)', 4746, 344, 25),
	(61278, 'LIWUTO', 4746, 344, 25),
	(61279, 'SUKANAYO (SUKANAEYO)', 4746, 344, 25),
	(61280, 'WARURUMA', 4746, 344, 25),
	(61281, 'KALIA LIA', 4747, 344, 25),
	(61282, 'KANTALAI', 4747, 344, 25),
	(61283, 'KOLESE', 4747, 344, 25),
	(61284, 'LOWU-LOWU', 4747, 344, 25),
	(61285, 'PALABUSA', 4747, 344, 25),
	(61286, 'BAADIA', 4748, 344, 25),
	(61287, 'LAMANGGA', 4748, 344, 25),
	(61288, 'MELAI', 4748, 344, 25),
	(61289, 'TANGANAPADA', 4748, 344, 25),
	(61290, 'WAJO', 4748, 344, 25),
	(61291, 'BUGI', 4749, 344, 25),
	(61292, 'GONDA BARU', 4749, 344, 25),
	(61293, 'KAISABU BARU', 4749, 344, 25),
	(61294, 'KARYA BARU', 4749, 344, 25),
	(61295, 'BATARAGURU', 4750, 344, 25),
	(61296, 'BATULO', 4750, 344, 25),
	(61297, 'BUKIT WOLIO INDAH', 4750, 344, 25),
	(61298, 'KADOLO KATAPI', 4750, 344, 25),
	(61299, 'TOMBA', 4750, 344, 25),
	(61300, 'WALE', 4750, 344, 25),
	(61301, 'WANGKANAPI', 4750, 344, 25),
	(61302, 'RAHADOPI', 4751, 345, 25),
	(61303, 'RAHAMPUU', 4751, 345, 25),
	(61304, 'TEOMOKOLE', 4751, 345, 25),
	(61305, 'TIRONGKOTUA', 4751, 345, 25),
	(61306, 'BALIARA', 4752, 345, 25),
	(61307, 'BALIARA KEPULAUAN', 4752, 345, 25),
	(61308, 'BALIARA SELATAN', 4752, 345, 25),
	(61309, 'RAHANTARI', 4752, 345, 25),
	(61310, 'SIKELI', 4752, 345, 25),
	(61311, 'BATUAWU', 4753, 345, 25),
	(61312, 'LANGKEMA', 4753, 345, 25),
	(61313, 'PONGKALAERO', 4753, 345, 25),
	(61314, 'PUU NUNU', 4753, 345, 25),
	(61315, 'ENANO', 4754, 345, 25),
	(61316, 'LAMONGGI', 4754, 345, 25),
	(61317, 'LENGORA', 4754, 345, 25),
	(61318, 'LENGORA PANTAI', 4754, 345, 25),
	(61319, 'LENGORA SELATAN', 4754, 345, 25),
	(61320, 'TANGKENO', 4754, 345, 25),
	(61321, 'ULUNGKURA', 4754, 345, 25),
	(61322, 'BALO', 4755, 345, 25),
	(61323, 'BUNGI-BUNGI', 4755, 345, 25),
	(61324, 'DONGKALA', 4755, 345, 25),
	(61325, 'LAMBALE', 4755, 345, 25),
	(61326, 'TAPUHAKA', 4755, 345, 25),
	(61327, 'TOLI-TOLI', 4755, 345, 25),
	(61328, 'WUMBUBURA', 4755, 345, 25),
	(61329, 'EEMOKOLO (EEOMOKOLA)', 4756, 345, 25),
	(61330, 'LAROLANU (LAROMANU)', 4756, 345, 25),
	(61331, 'MAPILA', 4756, 345, 25),
	(61332, 'SANGIA MAKMUR', 4756, 345, 25),
	(61333, 'TEDUBARA', 4756, 345, 25),
	(61334, 'WUMBULASA', 4756, 345, 25),
	(61335, 'BATU LAMBURI', 4757, 345, 25),
	(61336, 'MASALOKA', 4757, 345, 25),
	(61337, 'MASALOKA BARAT (TENGAH)', 4757, 345, 25),
	(61338, 'MASALOKA SELATAN', 4757, 345, 25),
	(61339, 'MASALOKA TIMUR', 4757, 345, 25),
	(61340, 'ANUGERAH', 4758, 345, 25),
	(61341, 'KALAERO', 4758, 345, 25),
	(61342, 'LANGKADUE', 4758, 345, 25),
	(61343, 'LANGKOWALA', 4758, 345, 25),
	(61344, 'LANTARI', 4758, 345, 25),
	(61345, 'LOMBA KASIH', 4758, 345, 25),
	(61346, 'PASARE APUA', 4758, 345, 25),
	(61347, 'RARONGKEU', 4758, 345, 25),
	(61348, 'TINABITE', 4758, 345, 25),
	(61349, 'WATU-WATU', 4758, 345, 25),
	(61350, 'BATUSAME INDAH', 4759, 345, 25),
	(61351, 'HAMBAWA INDAH', 4759, 345, 25),
	(61352, 'LALOA', 4759, 345, 25),
	(61353, 'LIANO', 4759, 345, 25),
	(61354, 'LORA', 4759, 345, 25),
	(61355, 'MAWAR', 4759, 345, 25),
	(61356, 'PEMONTORO (POMONTORO)', 4759, 345, 25),
	(61357, 'PULAU TAMBAKO', 4759, 345, 25),
	(61358, 'TAJUNCU', 4759, 345, 25),
	(61359, 'TOLI-TOLI', 4759, 345, 25),
	(61360, 'KOLOMBI MATAUSU', 4760, 345, 25),
	(61361, 'LAMURU', 4760, 345, 25),
	(61362, 'MORENGKE', 4760, 345, 25),
	(61363, 'TOTOLE', 4760, 345, 25),
	(61364, 'WIA-WIA', 4760, 345, 25),
	(61365, 'BARANGGA', 4761, 345, 25),
	(61366, 'BOEARA', 4761, 345, 25),
	(61367, 'BOEPINANG', 4761, 345, 25),
	(61368, 'BOEPINANG BARAT', 4761, 345, 25),
	(61369, 'KASABOLO', 4761, 345, 25),
	(61370, 'KASTARIB', 4761, 345, 25),
	(61371, 'MATIRO WALIE', 4761, 345, 25),
	(61372, 'PALLIMAE', 4761, 345, 25),
	(61373, 'PUKURUMBA (POKORUMBA)', 4761, 345, 25),
	(61374, 'SALOSA', 4761, 345, 25),
	(61375, 'BABA MOLINGKU', 4762, 345, 25),
	(61376, 'BALASARI', 4762, 345, 25),
	(61377, 'BULUMANAI', 4762, 345, 25),
	(61378, 'LAMEONG-MEONG', 4762, 345, 25),
	(61379, 'PABIRING', 4762, 345, 25),
	(61380, 'RAKADUA', 4762, 345, 25),
	(61381, 'RANO KOMEA', 4762, 345, 25),
	(61382, 'TIMBALA', 4762, 345, 25),
	(61383, 'TOARI BUTON', 4762, 345, 25),
	(61384, 'AKACIPONG', 4763, 345, 25),
	(61385, 'BATU PUTIH', 4763, 345, 25),
	(61386, 'KALI BARU', 4763, 345, 25),
	(61387, 'LA EA', 4763, 345, 25),
	(61388, 'WAEMPUTANG', 4763, 345, 25),
	(61389, 'LEBO EA', 4764, 345, 25),
	(61390, 'MULAENO', 4764, 345, 25),
	(61391, 'PARIA', 4764, 345, 25),
	(61392, 'POLEONDRO', 4764, 345, 25),
	(61393, 'LAMOARE', 4765, 345, 25),
	(61394, 'LARETE', 4765, 345, 25),
	(61395, 'MARAMPUKA', 4765, 345, 25),
	(61396, 'TERAPUNG', 4765, 345, 25),
	(61397, 'BAMBAEA', 4766, 345, 25),
	(61398, 'BIRU', 4766, 345, 25),
	(61399, 'MAMBO', 4766, 345, 25),
	(61400, 'PUU LEMO', 4766, 345, 25),
	(61401, 'TEPPOE', 4766, 345, 25),
	(61402, 'KARYA BARU', 4767, 345, 25),
	(61403, 'LAWATU EA (ES)', 4767, 345, 25),
	(61404, 'PULAU EA (ES)', 4767, 345, 25),
	(61405, 'ROMPU ROMPU', 4767, 345, 25),
	(61406, 'TAMPAWULU (TANPABULU)', 4767, 345, 25),
	(61407, 'TANAH POLEANG', 4767, 345, 25),
	(61408, 'TOBURI', 4767, 345, 25),
	(61409, 'WAMBAREMA', 4767, 345, 25),
	(61410, 'LADUMPI', 4768, 345, 25),
	(61411, 'LAKOMEA', 4768, 345, 25),
	(61412, 'LAMPEANTANI', 4768, 345, 25),
	(61413, 'PANGKURI', 4768, 345, 25),
	(61414, 'RAROWATU', 4768, 345, 25),
	(61415, 'RAU-RAU', 4768, 345, 25),
	(61416, 'TAUBONTO', 4768, 345, 25),
	(61417, 'WATU KALANGKARI', 4768, 345, 25),
	(61418, 'ANEKA MARGA', 4769, 345, 25),
	(61419, 'HUKAEA', 4769, 345, 25),
	(61420, 'LANTOWUA', 4769, 345, 25),
	(61421, 'MARGA JAYA', 4769, 345, 25),
	(61422, 'TEMBE', 4769, 345, 25),
	(61423, 'TUNAS BARU', 4769, 345, 25),
	(61424, 'WATU MENTADE', 4769, 345, 25),
	(61425, 'WUMBU BANGKA', 4769, 345, 25),
	(61426, 'BINA KARYA BUANA', 4770, 345, 25),
	(61427, 'BINAKARYA PUTRA', 4770, 345, 25),
	(61428, 'BONTO CINI', 4770, 345, 25),
	(61429, 'BONTOMANAI', 4770, 345, 25),
	(61430, 'BONTOTIRO', 4770, 345, 25),
	(61431, 'DOULE', 4770, 345, 25),
	(61432, 'JENETALLASA', 4770, 345, 25),
	(61433, 'KASIPUTE', 4770, 345, 25),
	(61434, 'KASSI', 4770, 345, 25),
	(61435, 'LAMERORO', 4770, 345, 25),
	(61436, 'LAMPOPALA', 4770, 345, 25),
	(61437, 'LANTO WONUA (LANTAWONUA)', 4770, 345, 25),
	(61438, 'LEBANG MANAI', 4770, 345, 25),
	(61439, 'LEBANG MANAI UTARA', 4770, 345, 25),
	(61440, 'LOKA', 4770, 345, 25),
	(61441, 'PALLANTIKANG', 4770, 345, 25),
	(61442, 'REKSO BINANGUN', 4770, 345, 25),
	(61443, 'RENO BASUKI', 4770, 345, 25),
	(61444, 'RESTU BARU', 4770, 345, 25),
	(61445, 'RESTU BUANA', 4770, 345, 25),
	(61446, 'RUKTI BASUKI', 4770, 345, 25),
	(61447, 'RUMBIA', 4770, 345, 25),
	(61448, 'TELUK DALEM ILIR', 4770, 345, 25),
	(61449, 'TOMPOBULU', 4770, 345, 25),
	(61450, 'UJUNG BULU', 4770, 345, 25),
	(61451, 'KAMPUNG BARU', 4771, 345, 25),
	(61452, 'LAMPATA', 4771, 345, 25),
	(61453, 'LAURU', 4771, 345, 25),
	(61454, 'POEA', 4771, 345, 25),
	(61455, 'TAPUHAHI', 4771, 345, 25),
	(61456, 'PUU WONUA', 4772, 345, 25),
	(61457, 'TETE HAKA', 4772, 345, 25),
	(61458, 'TONGKOSENG', 4772, 345, 25),
	(61459, 'TONTONUNU', 4772, 345, 25),
	(61460, 'WATU MELOMBA', 4772, 345, 25),
	(61461, 'BANDAR BATAUGA', 4773, 346, 25),
	(61462, 'BOLA', 4773, 346, 25),
	(61463, 'BUSOA', 4773, 346, 25),
	(61464, 'LAKAMBAU', 4773, 346, 25),
	(61465, 'LAOMPO', 4773, 346, 25),
	(61466, 'LAPANAIRI (LAMPANAIRI)', 4773, 346, 25),
	(61467, 'LAWELA', 4773, 346, 25),
	(61468, 'LAWELA SELATAN', 4773, 346, 25),
	(61469, 'MAJAPAHIT', 4773, 346, 25),
	(61470, 'MASIRI', 4773, 346, 25),
	(61471, 'MOLAGINA', 4773, 346, 25),
	(61472, 'POGALAMPA (POOGALAMPA)', 4773, 346, 25),
	(61473, 'BATU ATAS LIWU', 4774, 346, 25),
	(61474, 'BATUATAS BARAT', 4774, 346, 25),
	(61475, 'BATUATAS TIMUR', 4774, 346, 25),
	(61476, 'TADUASA', 4774, 346, 25),
	(61477, 'TOLANDO JAYA', 4774, 346, 25),
	(61478, 'WACUALA', 4774, 346, 25),
	(61479, 'WAMBONGI', 4774, 346, 25),
	(61480, 'BANTEA', 4775, 346, 25),
	(61481, 'BOMBONAWULU', 4775, 346, 25),
	(61482, 'KAMAMA MEKAR', 4775, 346, 25),
	(61483, 'KOLOWA', 4775, 346, 25),
	(61484, 'LAKAPERA', 4775, 346, 25),
	(61485, 'LOWU-LOWU', 4775, 346, 25),
	(61486, 'RAHIA', 4775, 346, 25),
	(61487, 'WADIABERO', 4775, 346, 25),
	(61488, 'WAKEA KEA', 4775, 346, 25),
	(61489, 'WALANDO', 4775, 346, 25),
	(61490, 'WALIKO', 4775, 346, 25),
	(61491, 'WATULEA', 4775, 346, 25),
	(61492, 'BANABUNGI', 4776, 346, 25),
	(61493, 'BANABUNGI SELATAN', 4776, 346, 25),
	(61494, 'KAOFE', 4776, 346, 25),
	(61495, 'KAPOA', 4776, 346, 25),
	(61496, 'KAPOA BARAT', 4776, 346, 25),
	(61497, 'LIPU', 4776, 346, 25),
	(61498, 'MARAWALI', 4776, 346, 25),
	(61499, 'MAWAMBUNGA', 4776, 346, 25),
	(61500, 'UWEMAASI', 4776, 346, 25),
	(61501, 'WAONU', 4776, 346, 25),
	(61502, 'BARANGKA', 4777, 346, 25),
	(61503, 'BONEATIRO', 4777, 346, 25),
	(61504, 'BONEATIRO BARAT', 4777, 346, 25),
	(61505, 'BUKIT ASRI', 4777, 346, 25),
	(61506, 'KAMELANTA', 4777, 346, 25),
	(61507, 'LAMBUSANGO', 4777, 346, 25),
	(61508, 'LAMBUSANGO TIMUR', 4777, 346, 25),
	(61509, 'MAMBULUGO', 4777, 346, 25),
	(61510, 'TODANGA', 4777, 346, 25),
	(61511, 'TUANGILA', 4777, 346, 25),
	(61512, 'TUMADA', 4777, 346, 25),
	(61513, 'WAKALAMBE', 4777, 346, 25),
	(61514, 'WAKANGKA', 4777, 346, 25),
	(61515, 'WAKULI', 4777, 346, 25),
	(61516, 'WAMBULU', 4777, 346, 25),
	(61517, 'WAONDO WOLIO', 4777, 346, 25),
	(61518, 'WATUMOTOBE', 4777, 346, 25),
	(61519, 'BONEOGE', 4778, 346, 25),
	(61520, 'LAKUDO', 4778, 346, 25),
	(61521, 'LOLIBU', 4778, 346, 25),
	(61522, 'MADONGKA', 4778, 346, 25),
	(61523, 'MATAWINE', 4778, 346, 25),
	(61524, 'METERE', 4778, 346, 25),
	(61525, 'MOKO', 4778, 346, 25),
	(61526, 'MONE', 4778, 346, 25),
	(61527, 'NEPA MEKAR', 4778, 346, 25),
	(61528, 'ONE WAARA', 4778, 346, 25),
	(61529, 'TELUKLASONGKO', 4778, 346, 25),
	(61530, 'WAARA', 4778, 346, 25),
	(61531, 'WAJOGU', 4778, 346, 25),
	(61532, 'WANEPA NEPA', 4778, 346, 25),
	(61533, 'WONGKO LAKUDO', 4778, 346, 25),
	(61534, 'BURANGASI', 4779, 346, 25),
	(61535, 'BURANGASI RUMBIA', 4779, 346, 25),
	(61536, 'GAYABARU', 4779, 346, 25),
	(61537, 'LAPANDEWA', 4779, 346, 25),
	(61538, 'LAPANDEWA JAYA', 4779, 346, 25),
	(61539, 'LAPANDEWA KAIDEA', 4779, 346, 25),
	(61540, 'LAPANDEWA MAKMUR', 4779, 346, 25),
	(61541, 'BENTENG', 4780, 346, 25),
	(61542, 'BONELALO', 4780, 346, 25),
	(61543, 'KAKENAUWE', 4780, 346, 25),
	(61544, 'KAMARU', 4780, 346, 25),
	(61545, 'LASEMBANGI', 4780, 346, 25),
	(61546, 'LAWELE', 4780, 346, 25),
	(61547, 'NAMBO', 4780, 346, 25),
	(61548, 'SRIBATARA', 4780, 346, 25),
	(61549, 'SUANDALA', 4780, 346, 25),
	(61550, 'TALAGA BARU', 4780, 346, 25),
	(61551, 'TOGO MANGURA', 4780, 346, 25),
	(61552, 'WAGARI', 4780, 346, 25),
	(61553, 'WAOLEONA', 4780, 346, 25),
	(61554, 'WASAMBAA', 4780, 346, 25),
	(61555, 'WASUAMBA', 4780, 346, 25),
	(61556, 'AMBUAU INDAH', 4781, 346, 25),
	(61557, 'AMBUAU TOGO', 4781, 346, 25),
	(61558, 'BALIMU', 4781, 346, 25),
	(61559, 'HARAPAN JAYA', 4781, 346, 25),
	(61560, 'KINAPANI MAKMUR', 4781, 346, 25),
	(61561, 'LASALIMU', 4781, 346, 25),
	(61562, 'MEGABAHARI', 4781, 346, 25),
	(61563, 'MOPAANO/MOPANO', 4781, 346, 25),
	(61564, 'MULIA JAYA', 4781, 346, 25),
	(61565, 'REJO SARI', 4781, 346, 25),
	(61566, 'SANGIA ARANO', 4781, 346, 25),
	(61567, 'SIOMANURU', 4781, 346, 25),
	(61568, 'SIONTAPINA (SIOTAPINA)', 4781, 346, 25),
	(61569, 'SUMBER AGUNG', 4781, 346, 25),
	(61570, 'UMALAOGE', 4781, 346, 25),
	(61571, 'WAJAHJAYA', 4781, 346, 25),
	(61572, 'AIR BAJO', 4782, 346, 25),
	(61573, 'BALOBONE', 4782, 346, 25),
	(61574, 'BANGA', 4782, 346, 25),
	(61575, 'DARIANGO', 4782, 346, 25),
	(61576, 'GUMAMANO', 4782, 346, 25),
	(61577, 'KANAPA-NAPA', 4782, 346, 25),
	(61578, 'KANCE BUNGI', 4782, 346, 25),
	(61579, 'MATARA (MATARA)', 4782, 346, 25),
	(61580, 'MAWASANGKA', 4782, 346, 25),
	(61581, 'NAPA', 4782, 346, 25),
	(61582, 'ONGKOLAKI (OENGKOLAKI)', 4782, 346, 25),
	(61583, 'POLINDU', 4782, 346, 25),
	(61584, 'TANAILANDU', 4782, 346, 25),
	(61585, 'TERAPUNG', 4782, 346, 25),
	(61586, 'WAKAMBANGURA', 4782, 346, 25),
	(61587, 'WAKAMBANGURA II', 4782, 346, 25),
	(61588, 'WASILOMATA I', 4782, 346, 25),
	(61589, 'WASILOMATA II', 4782, 346, 25),
	(61590, 'WATOLO (WATALO)', 4782, 346, 25),
	(61591, 'GUNDU GUNDU', 4783, 346, 25),
	(61592, 'KATUKOBARI', 4783, 346, 25),
	(61593, 'LAKUROA (LAKORUA)', 4783, 346, 25),
	(61594, 'LALIBO', 4783, 346, 25),
	(61595, 'LANGKOMU', 4783, 346, 25),
	(61596, 'LANTO', 4783, 346, 25),
	(61597, 'LANTONGAU', 4783, 346, 25),
	(61598, 'MORIKANA', 4783, 346, 25),
	(61599, 'WATORAMBE', 4783, 346, 25),
	(61600, 'WATORAMBE BATA', 4783, 346, 25),
	(61601, 'BATUBANAWA', 4784, 346, 25),
	(61602, 'BONE MARAMBE', 4784, 346, 25),
	(61603, 'BUNGI', 4784, 346, 25),
	(61604, 'INULU', 4784, 346, 25),
	(61605, 'LAGILI', 4784, 346, 25),
	(61606, 'LASORI', 4784, 346, 25),
	(61607, 'WAMBULOLI', 4784, 346, 25),
	(61608, 'WANTOPI', 4784, 346, 25),
	(61609, 'AWAINULU', 4785, 346, 25),
	(61610, 'BANABUNGI', 4785, 346, 25),
	(61611, 'DONGKALA', 4785, 346, 25),
	(61612, 'HOLIMOMBO JAYA', 4785, 346, 25),
	(61613, 'KABAWAKOLE', 4785, 346, 25),
	(61614, 'KAHULUNGAYA', 4785, 346, 25),
	(61615, 'KAMBULABULANA', 4785, 346, 25),
	(61616, 'KANCINAA', 4785, 346, 25),
	(61617, 'KAONGKEONGKEA', 4785, 346, 25),
	(61618, 'KOMBELI', 4785, 346, 25),
	(61619, 'KONDOWA', 4785, 346, 25),
	(61620, 'LABURUNCI', 4785, 346, 25),
	(61621, 'LAPODI', 4785, 346, 25),
	(61622, 'MANTOWU', 4785, 346, 25),
	(61623, 'PASARWAJO', 4785, 346, 25),
	(61624, 'SARAGI', 4785, 346, 25),
	(61625, 'TAKIMPO', 4785, 346, 25),
	(61626, 'WAGOLA', 4785, 346, 25),
	(61627, 'WANGU ANGU', 4785, 346, 25),
	(61628, 'WARINTA', 4785, 346, 25),
	(61629, 'WASAGA', 4785, 346, 25),
	(61630, 'WINNING', 4785, 346, 25),
	(61631, 'BAHARI', 4786, 346, 25),
	(61632, 'BAHARI DUA', 4786, 346, 25),
	(61633, 'BAHARI TIGA', 4786, 346, 25),
	(61634, 'BANGUN', 4786, 346, 25),
	(61635, 'GERAK MAKMUR', 4786, 346, 25),
	(61636, 'GUNUNG SEJUK', 4786, 346, 25),
	(61637, 'HENDEA', 4786, 346, 25),
	(61638, 'JAYA BAKTI', 4786, 346, 25),
	(61639, 'KATILOMBU', 4786, 346, 25),
	(61640, 'LIPUMANGAU', 4786, 346, 25),
	(61641, 'SANDANG PANGAN', 4786, 346, 25),
	(61642, 'TIRA', 4786, 346, 25),
	(61643, 'TODOMBULU', 4786, 346, 25),
	(61644, 'WATIGINANDA', 4786, 346, 25),
	(61645, 'WAWOANGI', 4786, 346, 25),
	(61646, 'WINDU MAKMUR', 4786, 346, 25),
	(61647, 'BARUTA', 4787, 346, 25),
	(61648, 'BARUTA ANALAKI', 4787, 346, 25),
	(61649, 'BARUTA LESTARI', 4787, 346, 25),
	(61650, 'DODA BAHARI', 4787, 346, 25),
	(61651, 'TOLANDONA', 4787, 346, 25),
	(61652, 'TOLANDONA MATANAEO', 4787, 346, 25),
	(61653, 'BATUAWU', 4788, 346, 25),
	(61654, 'BIWINAPADA', 4788, 346, 25),
	(61655, 'KAIMBULAWA', 4788, 346, 25),
	(61656, 'KARAE', 4788, 346, 25),
	(61657, 'LANTOI (LONTOI)', 4788, 346, 25),
	(61658, 'LAPARA', 4788, 346, 25),
	(61659, 'NGGULANGGULA', 4788, 346, 25),
	(61660, 'TONGALI', 4788, 346, 25),
	(61661, 'WAINDAWULA', 4788, 346, 25),
	(61662, 'WAKINAMBORO', 4788, 346, 25),
	(61663, 'KAMOALI', 4789, 346, 25),
	(61664, 'KATAMPE', 4789, 346, 25),
	(61665, 'LALOLE', 4789, 346, 25),
	(61666, 'LAMANINGGARA', 4789, 346, 25),
	(61667, 'MBANUA', 4789, 346, 25),
	(61668, 'MOKOBEAU', 4789, 346, 25),
	(61669, 'MOLONA', 4789, 346, 25),
	(61670, 'WATUAMPARA', 4789, 346, 25),
	(61671, 'BAHARI MAKMUR', 4790, 346, 25),
	(61672, 'GUNUNG JAYA', 4790, 346, 25),
	(61673, 'KARYA JAYA', 4790, 346, 25),
	(61674, 'KUMBEWAHA', 4790, 346, 25),
	(61675, 'KURAA', 4790, 346, 25),
	(61676, 'LABUANDIRI', 4790, 346, 25),
	(61677, 'MANURU', 4790, 346, 25),
	(61678, 'MATANAUWE', 4790, 346, 25),
	(61679, 'SAMPUABALO', 4790, 346, 25),
	(61680, 'SUMBER SARI', 4790, 346, 25),
	(61681, 'WALOMPO', 4790, 346, 25),
	(61682, 'KOKOE', 4791, 346, 25),
	(61683, 'LIWU LOMPONA', 4791, 346, 25),
	(61684, 'PANGILIA', 4791, 346, 25),
	(61685, 'TELAGA/TALAGA BESAR', 4791, 346, 25),
	(61686, 'TELAGA/TALAGA I', 4791, 346, 25),
	(61687, 'TELAGA/TALAGA II', 4791, 346, 25),
	(61688, 'WULU (WULUH)', 4791, 346, 25),
	(61689, 'BAJO BAHARI', 4792, 346, 25),
	(61690, 'HOLIMOMBO', 4792, 346, 25),
	(61691, 'KOHOLIMOMBONA', 4792, 346, 25),
	(61692, 'WABULA', 4792, 346, 25),
	(61693, 'WABULA SATU', 4792, 346, 25),
	(61694, 'WASAMPELA', 4792, 346, 25),
	(61695, 'WASUEMBA', 4792, 346, 25),
	(61696, 'BUNGI', 4793, 346, 25),
	(61697, 'GALANTI', 4793, 346, 25),
	(61698, 'KAUMBU', 4793, 346, 25),
	(61699, 'MATAWIA', 4793, 346, 25),
	(61700, 'SUKAMAJU', 4793, 346, 25),
	(61701, 'WOLOWA', 4793, 346, 25),
	(61702, 'WOLOWA BARU', 4793, 346, 25),
	(61703, 'BONEGUNU', 4794, 347, 25),
	(61704, 'BURANGA', 4794, 347, 25),
	(61705, 'DAMAI LABORONA', 4794, 347, 25),
	(61706, 'EEN SUMALA', 4794, 347, 25),
	(61707, 'GUNUNG SARI', 4794, 347, 25),
	(61708, 'KOBORUNO', 4794, 347, 25),
	(61709, 'KOEPISINO', 4794, 347, 25),
	(61710, 'LAANOIPI', 4794, 347, 25),
	(61711, 'LANGERE', 4794, 347, 25),
	(61712, 'NGAPA\\EA', 4794, 347, 25),
	(61713, 'RANTE GOLA', 4794, 347, 25),
	(61714, 'RONTA', 4794, 347, 25),
	(61715, 'TATOMBULI', 4794, 347, 25),
	(61716, 'WAODE ANGKALO', 4794, 347, 25),
	(61717, 'WAODE KALOWO', 4794, 347, 25),
	(61718, 'BALUARA', 4795, 347, 25),
	(61719, 'BENTE', 4795, 347, 25),
	(61720, 'BUBU', 4795, 347, 25),
	(61721, 'BUBU BARAT', 4795, 347, 25),
	(61722, 'KAMBOWA', 4795, 347, 25),
	(61723, 'KONDO', 4795, 347, 25),
	(61724, 'LAGUNDI', 4795, 347, 25),
	(61725, 'LAHAMOKO JAYA', 4795, 347, 25),
	(61726, 'MATA', 4795, 347, 25),
	(61727, 'MORINDINO', 4795, 347, 25),
	(61728, 'PONGKAWULU', 4795, 347, 25),
	(61729, 'BANGKUDU (EREKE)', 4796, 347, 25),
	(61730, 'BANU-BANUA JAYA', 4796, 347, 25),
	(61731, 'BONELIPU', 4796, 347, 25),
	(61732, 'EELAHAJI', 4796, 347, 25),
	(61733, 'JAMPAKA', 4796, 347, 25),
	(61734, 'KADACUA', 4796, 347, 25),
	(61735, 'KALIBU', 4796, 347, 25),
	(61736, 'LAANGKE', 4796, 347, 25),
	(61737, 'LAKONEA', 4796, 347, 25),
	(61738, 'LANTAGI', 4796, 347, 25),
	(61739, 'LEMO (BONEROMBO)', 4796, 347, 25),
	(61740, 'LEMO EA', 4796, 347, 25),
	(61741, 'LINSOWU', 4796, 347, 25),
	(61742, 'LIPU', 4796, 347, 25),
	(61743, 'LOJI', 4796, 347, 25),
	(61744, 'MALALANDA', 4796, 347, 25),
	(61745, 'ROMBO', 4796, 347, 25),
	(61746, 'SARA\\EA', 4796, 347, 25),
	(61747, 'TOMOAHI', 4796, 347, 25),
	(61748, 'TRI WACU-WACU', 4796, 347, 25),
	(61749, 'WACULAEA', 4796, 347, 25),
	(61750, 'WANDAKA', 4796, 347, 25),
	(61751, 'WASALABOSE', 4796, 347, 25),
	(61752, 'BUMI LAPERO/LAPERA', 4797, 347, 25),
	(61753, 'KARYA BAKTI', 4797, 347, 25),
	(61754, 'KARYA MULYA', 4797, 347, 25),
	(61755, 'KASULATOMBI', 4797, 347, 25),
	(61756, 'KOTAWO (KOTAWA)', 4797, 347, 25),
	(61757, 'LABULANDA', 4797, 347, 25),
	(61758, 'LAMBALE', 4797, 347, 25),
	(61759, 'LAMPALA JAYA (DAMPALA JAYA)', 4797, 347, 25),
	(61760, 'LAPANDEWA', 4797, 347, 25),
	(61761, 'LAUKI', 4797, 347, 25),
	(61762, 'MARGA KARYA', 4797, 347, 25),
	(61763, 'MEKAR JAYA', 4797, 347, 25),
	(61764, 'RAHMAT BARU', 4797, 347, 25),
	(61765, 'SOLOY/SOIOI AGUNG', 4797, 347, 25),
	(61766, 'BIRA', 4798, 347, 25),
	(61767, 'E\\ERINERE', 4798, 347, 25),
	(61768, 'KUROLABU', 4798, 347, 25),
	(61769, 'LAANOSANGIA (LANOSANGIA)', 4798, 347, 25),
	(61770, 'LABELETE', 4798, 347, 25),
	(61771, 'LAMOAHI', 4798, 347, 25),
	(61772, 'LELAMO', 4798, 347, 25),
	(61773, 'PEBAOA', 4798, 347, 25),
	(61774, 'PETETEAA', 4798, 347, 25),
	(61775, 'TOROMBIA (TORUMBIA)', 4798, 347, 25),
	(61776, 'ULUNAMBO', 4798, 347, 25),
	(61777, 'WAMBOULE', 4798, 347, 25),
	(61778, 'WAODE BURI', 4798, 347, 25),
	(61779, 'WOWONGA JAYA', 4798, 347, 25),
	(61780, 'LABA JAYA', 4799, 347, 25),
	(61781, 'LABARAGA', 4799, 347, 25),
	(61782, 'LABUAN', 4799, 347, 25),
	(61783, 'LABUAN BAJO', 4799, 347, 25),
	(61784, 'LABUAN WALIO', 4799, 347, 25),
	(61785, 'LABUKO', 4799, 347, 25),
	(61786, 'LAEA', 4799, 347, 25),
	(61787, 'LASIWA', 4799, 347, 25),
	(61788, 'MATALAGI', 4799, 347, 25),
	(61789, 'OENGKAPALA', 4799, 347, 25),
	(61790, 'SUMAMPENO', 4799, 347, 25),
	(61791, 'WAMORAPA', 4799, 347, 25),
	(61792, 'WANTULASI', 4799, 347, 25),
	(61793, 'ABELI', 4800, 348, 25),
	(61794, 'ANGGALOMELAI (ANGGOLOMELAI)', 4800, 348, 25),
	(61795, 'BENUA NIRAE', 4800, 348, 25),
	(61796, 'BUNGKUTOKO', 4800, 348, 25),
	(61797, 'LAPULU', 4800, 348, 25),
	(61798, 'NAMBO', 4800, 348, 25),
	(61799, 'PETOAHA', 4800, 348, 25),
	(61800, 'POASIA', 4800, 348, 25),
	(61801, 'PUDAY (PUDAI)', 4800, 348, 25),
	(61802, 'SAMBULI', 4800, 348, 25),
	(61803, 'TALIA', 4800, 348, 25),
	(61804, 'TOBIMEITA', 4800, 348, 25),
	(61805, 'TONDONGGEU', 4800, 348, 25),
	(61806, 'BARUGA', 4801, 348, 25),
	(61807, 'LEPO-LEPO', 4801, 348, 25),
	(61808, 'WATUBANGGA', 4801, 348, 25),
	(61809, 'WUNDUPOPI', 4801, 348, 25),
	(61810, 'ANAIWOI', 4802, 348, 25),
	(61811, 'BENDE', 4802, 348, 25),
	(61812, 'KADIA', 4802, 348, 25),
	(61813, 'PONDAMBEA', 4802, 348, 25),
	(61814, 'WOWAWANGGU', 4802, 348, 25),
	(61815, 'KAMBU', 4803, 348, 25),
	(61816, 'LALOLARA', 4803, 348, 25),
	(61817, 'MOKOAU', 4803, 348, 25),
	(61818, 'PADALEU', 4803, 348, 25),
	(61819, 'GUNUNG JATI', 4804, 348, 25),
	(61820, 'JATI MEKAR', 4804, 348, 25),
	(61821, 'KAMPUNG SALO', 4804, 348, 25),
	(61822, 'KANDAI', 4804, 348, 25),
	(61823, 'KENDARI CADDI', 4804, 348, 25),
	(61824, 'KESSILAMPE', 4804, 348, 25),
	(61825, 'MANGGA DUA', 4804, 348, 25),
	(61826, 'MATA', 4804, 348, 25),
	(61827, 'PURIRANO', 4804, 348, 25),
	(61828, 'BENU-BENUA', 4805, 348, 25),
	(61829, 'DAPU-DAPURA', 4805, 348, 25),
	(61830, 'KEMARAYA', 4805, 348, 25),
	(61831, 'LAHUNDAPE', 4805, 348, 25),
	(61832, 'PUNGGALOBA', 4805, 348, 25),
	(61833, 'SANUA', 4805, 348, 25),
	(61834, 'SODOHOA', 4805, 348, 25),
	(61835, 'TIPULU', 4805, 348, 25),
	(61836, 'WATU-WATU', 4805, 348, 25),
	(61837, 'ALOLAMA', 4806, 348, 25),
	(61838, 'ANGGILOWU', 4806, 348, 25),
	(61839, 'KORUMBA', 4806, 348, 25),
	(61840, 'LABIBIA', 4806, 348, 25),
	(61841, 'MANDONGA', 4806, 348, 25),
	(61842, 'WAWOMBALATA', 4806, 348, 25),
	(61843, 'ANDUONOHU', 4807, 348, 25),
	(61844, 'ANGGOEYA', 4807, 348, 25),
	(61845, 'MATABUBU', 4807, 348, 25),
	(61846, 'RAHANDOUNA', 4807, 348, 25),
	(61847, 'ABELI DALAM', 4808, 348, 25),
	(61848, 'LALODATI', 4808, 348, 25),
	(61849, 'PUNGGOLAKA', 4808, 348, 25),
	(61850, 'PUUWATU (PUWATU)', 4808, 348, 25),
	(61851, 'TOBUUHA', 4808, 348, 25),
	(61852, 'WATULONDO', 4808, 348, 25),
	(61853, 'ANAWAI', 4809, 348, 25),
	(61854, 'BONGGOEYA', 4809, 348, 25),
	(61855, 'MATAIWOI', 4809, 348, 25),
	(61856, 'WUA-WUA', 4809, 348, 25),
	(61857, 'BAULA', 4810, 349, 25),
	(61858, 'LONGORI / LANGORI', 4810, 349, 25),
	(61859, 'PEWUTAA', 4810, 349, 25),
	(61860, 'PUUBENUA', 4810, 349, 25),
	(61861, 'PUUBUNGA', 4810, 349, 25),
	(61862, 'PUULEMO', 4810, 349, 25),
	(61863, 'PUUNDOHO', 4810, 349, 25),
	(61864, 'PUURODA', 4810, 349, 25),
	(61865, 'WATALARA', 4810, 349, 25),
	(61866, 'BALANDETE', 4811, 349, 25),
	(61867, 'LALOEHA', 4811, 349, 25),
	(61868, 'LALOMBAA', 4811, 349, 25),
	(61869, 'LAMOKATO', 4811, 349, 25),
	(61870, 'SABILAMBO', 4811, 349, 25),
	(61871, 'TAHOA', 4811, 349, 25),
	(61872, 'WATULIANDU', 4811, 349, 25),
	(61873, 'ANGGALOOSI', 4812, 349, 25),
	(61874, 'ATULA', 4812, 349, 25),
	(61875, 'DANGIA', 4812, 349, 25),
	(61876, 'GUNUNG JAYA', 4812, 349, 25),
	(61877, 'LADONGI (JAYA)', 4812, 349, 25),
	(61878, 'LALOWOSULA/LALOWOSUKA', 4812, 349, 25),
	(61879, 'LEMBAH SUBUR', 4812, 349, 25),
	(61880, 'POMBEYOHA', 4812, 349, 25),
	(61881, 'PUTEMATA', 4812, 349, 25),
	(61882, 'RARAA', 4812, 349, 25),
	(61883, 'WANDE', 4812, 349, 25),
	(61884, 'WELALA', 4812, 349, 25),
	(61885, 'WUNGGOLOKO/WUNGGOLOLEO', 4812, 349, 25),
	(61886, 'KEISIO', 4813, 349, 25),
	(61887, 'LALOLAE', 4813, 349, 25),
	(61888, 'LALOSINGI', 4813, 349, 25),
	(61889, 'TALODO', 4813, 349, 25),
	(61890, 'WESALO', 4813, 349, 25),
	(61891, 'AERE', 4814, 349, 25),
	(61892, 'ALADADIO', 4814, 349, 25),
	(61893, 'ATOLANU', 4814, 349, 25),
	(61894, 'AWIU', 4814, 349, 25),
	(61895, 'BOU', 4814, 349, 25),
	(61896, 'INOTU', 4814, 349, 25),
	(61897, 'IWOIMEA JAYA', 4814, 349, 25),
	(61898, 'IWOIMENGGURA', 4814, 349, 25),
	(61899, 'LALOLERE', 4814, 349, 25),
	(61900, 'LAMBANDIA', 4814, 349, 25),
	(61901, 'LERE JAYA', 4814, 349, 25),
	(61902, 'LOWA', 4814, 349, 25),
	(61903, 'MOKUPA', 4814, 349, 25),
	(61904, 'PEKOREA', 4814, 349, 25),
	(61905, 'PENANGGO JAYA', 4814, 349, 25),
	(61906, 'PENANGGOOSI', 4814, 349, 25),
	(61907, 'POMBUREA', 4814, 349, 25),
	(61908, 'TAORE', 4814, 349, 25),
	(61909, 'TINETE', 4814, 349, 25),
	(61910, 'WONUAMBUTEO', 4814, 349, 25),
	(61911, 'INDUHA', 4815, 349, 25),
	(61912, 'KOLAKAASIH', 4815, 349, 25),
	(61913, 'LATAMBAGA', 4815, 349, 25),
	(61914, 'MANGOLO', 4815, 349, 25),
	(61915, 'SAKULI', 4815, 349, 25),
	(61916, 'SEA', 4815, 349, 25),
	(61917, 'ULUNGGOLAKA', 4815, 349, 25),
	(61918, 'IWOIKONDO', 4816, 349, 25),
	(61919, 'LALOWURA', 4816, 349, 25),
	(61920, 'LAMOARE', 4816, 349, 25),
	(61921, 'LOEA', 4816, 349, 25),
	(61922, 'MATAIWOI', 4816, 349, 25),
	(61923, 'PEATOA', 4816, 349, 25),
	(61924, 'SIMBALAI/SIMBALAE', 4816, 349, 25),
	(61925, 'TEPOSUA', 4816, 349, 25),
	(61926, 'HORODOPI', 4817, 349, 25),
	(61927, 'INEBENGGI', 4817, 349, 25),
	(61928, 'LAPANGISI', 4817, 349, 25),
	(61929, 'NELOMBU', 4817, 349, 25),
	(61930, 'PUOSU', 4817, 349, 25),
	(61931, 'ULU MOWEWE', 4817, 349, 25),
	(61932, 'WATUPUTE', 4817, 349, 25),
	(61933, 'WOITOMBO', 4817, 349, 25),
	(61934, 'ANDOWENGGA', 4818, 349, 25),
	(61935, 'PANGI-PANGI', 4818, 349, 25),
	(61936, 'POLEMAJU JAYA', 4818, 349, 25),
	(61937, 'POLENGA JAYA', 4818, 349, 25),
	(61938, 'POLI POLIA', 4818, 349, 25),
	(61939, 'TAOSU', 4818, 349, 25),
	(61940, 'TOKAI', 4818, 349, 25),
	(61941, 'WIA-WIA', 4818, 349, 25),
	(61942, 'PLASMA JAYA', 4819, 349, 25),
	(61943, 'POLINGGONA', 4819, 349, 25),
	(61944, 'PONDOUWAE', 4819, 349, 25),
	(61945, 'PUUDONGI', 4819, 349, 25),
	(61946, 'TANGGEAU', 4819, 349, 25),
	(61947, 'WULONGGERE', 4819, 349, 25),
	(61948, 'DAWI-DAWI', 4820, 349, 25),
	(61949, 'HAKATUTOBU', 4820, 349, 25),
	(61950, 'HUKO HUKO', 4820, 349, 25),
	(61951, 'KUMORO', 4820, 349, 25),
	(61952, 'OKO OKO', 4820, 349, 25),
	(61953, 'PELAMBUA', 4820, 349, 25),
	(61954, 'PESOUHA', 4820, 349, 25),
	(61955, 'POMALAA', 4820, 349, 25),
	(61956, 'SOPURA', 4820, 349, 25),
	(61957, 'TAMBEA', 4820, 349, 25),
	(61958, 'TONGGONI', 4820, 349, 25),
	(61959, 'TOTOBO', 4820, 349, 25),
	(61960, 'AMOMUTU', 4821, 349, 25),
	(61961, 'AWA', 4821, 349, 25),
	(61962, 'KALOLOA', 4821, 349, 25),
	(61963, 'KONAWEHA', 4821, 349, 25),
	(61964, 'LAMBOLEMO', 4821, 349, 25),
	(61965, 'LATUO', 4821, 349, 25),
	(61966, 'LAWULO', 4821, 349, 25),
	(61967, 'MALAHA', 4821, 349, 25),
	(61968, 'PUU TAMBOLI', 4821, 349, 25),
	(61969, 'SANI SANI', 4821, 349, 25),
	(61970, 'TAMBOLI', 4821, 349, 25),
	(61971, 'TONGANAPO', 4821, 349, 25),
	(61972, 'TOSIBA', 4821, 349, 25),
	(61973, 'ULU KONAWEHA', 4821, 349, 25),
	(61974, 'WOWA TAMBOLI', 4821, 349, 25),
	(61975, 'ANAIWOI', 4822, 349, 25),
	(61976, 'LALONGGOLOSUA', 4822, 349, 25),
	(61977, 'LAMEDAI', 4822, 349, 25),
	(61978, 'LAMOIKO', 4822, 349, 25),
	(61979, 'ONEEHA', 4822, 349, 25),
	(61980, 'PALEWAI', 4822, 349, 25),
	(61981, 'PETUDUA', 4822, 349, 25),
	(61982, 'POPALIA', 4822, 349, 25),
	(61983, 'POWISOA JAYA (PEWISOA JAYA)', 4822, 349, 25),
	(61984, 'PUUNDAIPA', 4822, 349, 25),
	(61985, 'RAHANGGADA', 4822, 349, 25),
	(61986, 'TANGGETADA', 4822, 349, 25),
	(61987, 'TONDOWOLIO', 4822, 349, 25),
	(61988, 'AMBAPA', 4823, 349, 25),
	(61989, 'AMERORO', 4823, 349, 25),
	(61990, 'LAMUNDE', 4823, 349, 25),
	(61991, 'SOLEWATU', 4823, 349, 25),
	(61992, 'TALATA', 4823, 349, 25),
	(61993, 'TAWAROMBADAKA', 4823, 349, 25),
	(61994, 'TINONDO', 4823, 349, 25),
	(61995, 'WEAMO', 4823, 349, 25),
	(61996, 'LALINGATO', 4824, 349, 25),
	(61997, 'LARA', 4824, 349, 25),
	(61998, 'LOKA', 4824, 349, 25),
	(61999, 'ORAWA', 4824, 349, 25),
	(62000, 'PONI PONIKI', 4824, 349, 25)
	`)
}
