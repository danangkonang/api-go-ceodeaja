package kel10

import "github.com/danangkonang/ceodeaja-go/config"

func Kel18() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
		(18001, 'PEDAGANGAN', 1437, 76, 6),
		(18002, 'SEMBUNG', 1437, 76, 6),
		(18003, 'SOKO (SOOKO)', 1437, 76, 6),
		(18004, 'SUMBERAME', 1437, 76, 6),
		(18005, 'SUMBERGEDE', 1437, 76, 6),
		(18006, 'SUMBERWARU', 1437, 76, 6),
		(18007, 'SUMENGKO', 1437, 76, 6),
		(18008, 'WATESTANJUNG', 1437, 76, 6),
		(18009, 'WRINGINANOM', 1437, 76, 6),
		(18010, 'AJUNG', 1438, 77, 6),
		(18011, 'KLOMPANGAN', 1438, 77, 6),
		(18012, 'MANGARAN', 1438, 77, 6),
		(18013, 'PANCAKARYA', 1438, 77, 6),
		(18014, 'ROWO INDAH', 1438, 77, 6),
		(18015, 'SUKAMAKMUR', 1438, 77, 6),
		(18016, 'WIROWONGSO', 1438, 77, 6),
		(18017, 'AMBULU', 1439, 77, 6),
		(18018, 'ANDONGSARI', 1439, 77, 6),
		(18019, 'KARANG ANYAR', 1439, 77, 6),
		(18020, 'PONTANG', 1439, 77, 6),
		(18021, 'SABRANG', 1439, 77, 6),
		(18022, 'SUMBERREJO', 1439, 77, 6),
		(18023, 'TEGALSARI', 1439, 77, 6),
		(18024, 'BALUNG KIDUL', 1440, 77, 6),
		(18025, 'BALUNG KULON', 1440, 77, 6),
		(18026, 'BALUNG LOR', 1440, 77, 6),
		(18027, 'CURAHLELE', 1440, 77, 6),
		(18028, 'GUMELAR', 1440, 77, 6),
		(18029, 'KARANG DUREN', 1440, 77, 6),
		(18030, 'KARANG SEMANDING', 1440, 77, 6),
		(18031, 'TUTUL', 1440, 77, 6),
		(18032, 'BADEAN', 1441, 77, 6),
		(18033, 'BANGSALSARI', 1441, 77, 6),
		(18034, 'BANJARSARI', 1441, 77, 6),
		(18035, 'CURAH KALONG', 1441, 77, 6),
		(18036, 'GAMBIRONO', 1441, 77, 6),
		(18037, 'KARANGSONO', 1441, 77, 6),
		(18038, 'LANGKAP', 1441, 77, 6),
		(18039, 'PETUNG', 1441, 77, 6),
		(18040, 'SUKOREJO', 1441, 77, 6),
		(18041, 'TISNOGAMBAR', 1441, 77, 6),
		(18042, 'TUGUSARI', 1441, 77, 6),
		(18043, 'BAGOREJO', 1442, 77, 6),
		(18044, 'GUMUKMAS', 1442, 77, 6),
		(18045, 'KARANG REJO', 1442, 77, 6),
		(18046, 'KEPANJEN', 1442, 77, 6),
		(18047, 'MAYANGAN', 1442, 77, 6),
		(18048, 'MENAMPU', 1442, 77, 6),
		(18049, 'PURWOASRI', 1442, 77, 6),
		(18050, 'TEMBOKREJO', 1442, 77, 6),
		(18051, 'JELBUK', 1443, 77, 6),
		(18052, 'PANDUMAN', 1443, 77, 6),
		(18053, 'SUCO PANGEPOK', 1443, 77, 6),
		(18054, 'SUGER KIDUL', 1443, 77, 6),
		(18055, 'SUKO JEMBER', 1443, 77, 6),
		(18056, 'SUKOWIRYO', 1443, 77, 6),
		(18057, 'CANGKRING', 1444, 77, 6),
		(18058, 'JATIMULYO', 1444, 77, 6),
		(18059, 'JATISARI', 1444, 77, 6),
		(18060, 'JENGGAWAH', 1444, 77, 6),
		(18061, 'KEMUNING SARI KIDUL', 1444, 77, 6),
		(18062, 'KERTONEGORO', 1444, 77, 6),
		(18063, 'SERUNI', 1444, 77, 6),
		(18064, 'WONOJATI', 1444, 77, 6),
		(18065, 'BANJAR DOWO', 1445, 77, 6),
		(18066, 'CANDI MULYO', 1445, 77, 6),
		(18067, 'DAPUR KEJAMBON', 1445, 77, 6),
		(18068, 'DENANYAR', 1445, 77, 6),
		(18069, 'GEDONG DALEM', 1445, 77, 6),
		(18070, 'JABON', 1445, 77, 6),
		(18071, 'JELAKOMBO', 1445, 77, 6),
		(18072, 'JOMBANG', 1445, 77, 6),
		(18073, 'JOMBANG', 1445, 77, 6),
		(18074, 'JOMBANG WETAN', 1445, 77, 6),
		(18075, 'JOMBATAN', 1445, 77, 6),
		(18076, 'KALIWUNGU', 1445, 77, 6),
		(18077, 'KEPANJEN', 1445, 77, 6),
		(18078, 'KEPATIHAN', 1445, 77, 6),
		(18079, 'KETING', 1445, 77, 6),
		(18080, 'MASIGIT', 1445, 77, 6),
		(18081, 'MOJONGAPIT', 1445, 77, 6),
		(18082, 'NGAMPELREJO', 1445, 77, 6),
		(18083, 'PADOMASAN', 1445, 77, 6),
		(18084, 'PANGGUNG RAWI', 1445, 77, 6),
		(18085, 'PLANDI', 1445, 77, 6),
		(18086, 'PLOSOGENENG', 1445, 77, 6),
		(18087, 'PULO LOR', 1445, 77, 6),
		(18088, 'SAMBONG DUKUH', 1445, 77, 6),
		(18089, 'SARIMULYO', 1445, 77, 6),
		(18090, 'SENGON', 1445, 77, 6),
		(18091, 'SUKMAJAYA', 1445, 77, 6),
		(18092, 'SUMBERJO', 1445, 77, 6),
		(18093, 'TAMBAK REJO', 1445, 77, 6),
		(18094, 'TUNGGORONO', 1445, 77, 6),
		(18095, 'WRINGIN AGUNG', 1445, 77, 6),
		(18096, 'AJUNG', 1446, 77, 6),
		(18097, 'GAMBIRAN', 1446, 77, 6),
		(18098, 'GLAGAHWERO', 1446, 77, 6),
		(18099, 'GUMUKSARI', 1446, 77, 6),
		(18100, 'KALISAT', 1446, 77, 6),
		(18101, 'PATEMPURAN', 1446, 77, 6),
		(18102, 'PLALANGAN', 1446, 77, 6),
		(18103, 'SEBANEN', 1446, 77, 6),
		(18104, 'SUKORENO', 1446, 77, 6),
		(18105, 'SUMBER JERUK', 1446, 77, 6),
		(18106, 'SUMBER KALONG', 1446, 77, 6),
		(18107, 'SUMBER KETEMPAH', 1446, 77, 6),
		(18108, 'JEMBER KIDUL', 1447, 77, 6),
		(18109, 'KALIWATES', 1447, 77, 6),
		(18110, 'KEBON AGUNG', 1447, 77, 6),
		(18111, 'KEPATIHAN', 1447, 77, 6),
		(18112, 'MANGLI', 1447, 77, 6),
		(18113, 'SEMPUSARI', 1447, 77, 6),
		(18114, 'TEGAL BESAR', 1447, 77, 6),
		(18115, 'CAKRU', 1448, 77, 6),
		(18116, 'KENCONG', 1448, 77, 6),
		(18117, 'KRATON', 1448, 77, 6),
		(18118, 'PASEBAN', 1448, 77, 6),
		(18119, 'WONOREJO', 1448, 77, 6),
		(18120, 'KARANG PAITON', 1449, 77, 6),
		(18121, 'LEDOKOMBO', 1449, 77, 6),
		(18122, 'LEMBENGAN', 1449, 77, 6),
		(18123, 'SLATENG', 1449, 77, 6),
		(18124, 'SUKOGIDRI', 1449, 77, 6),
		(18125, 'SUMBER ANGET', 1449, 77, 6),
		(18126, 'SUMBER BULUS', 1449, 77, 6),
		(18127, 'SUMBER LESUNG', 1449, 77, 6),
		(18128, 'SUMBER SALAK', 1449, 77, 6),
		(18129, 'SUREN', 1449, 77, 6),
		(18130, 'MAYANG', 1450, 77, 6),
		(18131, 'MRAWAN', 1450, 77, 6),
		(18132, 'SEPUTIH', 1450, 77, 6),
		(18133, 'SIDOMUKTI', 1450, 77, 6),
		(18134, 'SUMBER KEJAYAN', 1450, 77, 6),
		(18135, 'TEGAL WARU', 1450, 77, 6),
		(18136, 'TEGALREJO', 1450, 77, 6),
		(18137, 'KARANGKEDAWUNG', 1451, 77, 6),
		(18138, 'KAWANGREJO', 1451, 77, 6),
		(18139, 'LAMPEJI', 1451, 77, 6),
		(18140, 'LENGKONG', 1451, 77, 6),
		(18141, 'MUMBULSARI', 1451, 77, 6),
		(18142, 'SUCO', 1451, 77, 6),
		(18143, 'TAMANSARI', 1451, 77, 6),
		(18144, 'BEDADUNG', 1452, 77, 6),
		(18145, 'JATIAN', 1452, 77, 6),
		(18146, 'KERTOSARI', 1452, 77, 6),
		(18147, 'PAKUSARI', 1452, 77, 6),
		(18148, 'PATEMON', 1452, 77, 6),
		(18149, 'SUBO', 1452, 77, 6),
		(18150, 'SUMBER PINANG', 1452, 77, 6),
		(18151, 'GLAGAHWERO', 1453, 77, 6),
		(18152, 'KEMIRI', 1453, 77, 6),
		(18153, 'KEMUNINGSARI LOR', 1453, 77, 6),
		(18154, 'PAKIS', 1453, 77, 6),
		(18155, 'PANTI', 1453, 77, 6),
		(18156, 'PANTI', 1453, 77, 6),
		(18157, 'SERUT', 1453, 77, 6),
		(18158, 'SUCI', 1453, 77, 6),
		(18159, 'BANJAR SENGON', 1454, 77, 6),
		(18160, 'BARATAN', 1454, 77, 6),
		(18161, 'BINTORO', 1454, 77, 6),
		(18162, 'GEBANG', 1454, 77, 6),
		(18163, 'JEMBER LOR', 1454, 77, 6),
		(18164, 'JUMERTO', 1454, 77, 6),
		(18165, 'PATRANG', 1454, 77, 6),
		(18166, 'SLAWU', 1454, 77, 6),
		(18167, 'BAGON', 1455, 77, 6),
		(18168, 'GRENDEN', 1455, 77, 6),
		(18169, 'JAMBEARUM', 1455, 77, 6),
		(18170, 'KASIYAN', 1455, 77, 6),
		(18171, 'KASIYAN TIMUR', 1455, 77, 6),
		(18172, 'MLOKOREJO', 1455, 77, 6),
		(18173, 'MOJOMULYO', 1455, 77, 6),
		(18174, 'MOJOSARI', 1455, 77, 6),
		(18175, 'PUGER KULON', 1455, 77, 6),
		(18176, 'PUGER WETAN', 1455, 77, 6),
		(18177, 'WONOSARI', 1455, 77, 6),
		(18178, 'WRINGIN TELU', 1455, 77, 6),
		(18179, 'CURAHMALANG', 1456, 77, 6),
		(18180, 'GUGUT', 1456, 77, 6),
		(18181, 'KALIWINING', 1456, 77, 6),
		(18182, 'NOGOSARI', 1456, 77, 6),
		(18183, 'PECORO', 1456, 77, 6),
		(18184, 'RAMBIGUNDAM', 1456, 77, 6),
		(18185, 'RAMBIPUJI', 1456, 77, 6),
		(18186, 'ROWOTAMTU', 1456, 77, 6),
		(18187, 'PONDOK DALEM', 1457, 77, 6),
		(18188, 'PONDOK JOYO', 1457, 77, 6),
		(18189, 'REJO AGUNG', 1457, 77, 6),
		(18190, 'SEMBORO', 1457, 77, 6),
		(18191, 'SIDOMEKAR', 1457, 77, 6),
		(18192, 'SIDOMULYO', 1457, 77, 6),
		(18193, 'GARAHAN', 1458, 77, 6),
		(18194, 'HARJOMOLYO', 1458, 77, 6),
		(18195, 'KARANGHARJO', 1458, 77, 6),
		(18196, 'MULYOREJO', 1458, 77, 6),
		(18197, 'PACE', 1458, 77, 6),
		(18198, 'SEMPOLAN', 1458, 77, 6),
		(18199, 'SIDOMULYO', 1458, 77, 6),
		(18200, 'SILO', 1458, 77, 6),
		(18201, 'SUMBERJATI', 1458, 77, 6),
		(18202, 'DUKUH MENCEK', 1459, 77, 6),
		(18203, 'JUBUNG', 1459, 77, 6),
		(18204, 'KARANGPRING', 1459, 77, 6),
		(18205, 'KELUNGKUNG', 1459, 77, 6),
		(18206, 'SUKORAMBI', 1459, 77, 6),
		(18207, 'ARJASA', 1460, 77, 6),
		(18208, 'BALET BARU', 1460, 77, 6),
		(18209, 'DAWUHAN MANGLI', 1460, 77, 6),
		(18210, 'MOJOGENI', 1460, 77, 6),
		(18211, 'POCANGAN', 1460, 77, 6),
		(18212, 'SUKOKERTO', 1460, 77, 6),
		(18213, 'SUKOREJO', 1460, 77, 6),
		(18214, 'SUKOSARI', 1460, 77, 6),
		(18215, 'SUKOWONO', 1460, 77, 6),
		(18216, 'SUMBER WRINGIN', 1460, 77, 6),
		(18217, 'SUMBERDANTI', 1460, 77, 6),
		(18218, 'SUMBERWARU', 1460, 77, 6),
		(18219, 'GELANG', 1461, 77, 6),
		(18220, 'JAMBESARI', 1461, 77, 6),
		(18221, 'JAMINTORO', 1461, 77, 6),
		(18222, 'JATIROTO', 1461, 77, 6),
		(18223, 'KALIGLAGAH', 1461, 77, 6),
		(18224, 'KARANG BAYAT', 1461, 77, 6),
		(18225, 'PRINGGOWIRAWAN', 1461, 77, 6),
		(18226, 'ROWO TENGAH', 1461, 77, 6),
		(18227, 'SUMBER AGUNG', 1461, 77, 6),
		(18228, 'YOSORATI', 1461, 77, 6),
		(18229, 'CUMEDAK', 1462, 77, 6),
		(18230, 'GUNUNG MALANG', 1462, 77, 6),
		(18231, 'JAMBE ARUM', 1462, 77, 6),
		(18232, 'PLEREYAN', 1462, 77, 6),
		(18233, 'PRINGGONDANI', 1462, 77, 6),
		(18234, 'RANDU AGUNG', 1462, 77, 6),
		(18235, 'ROWOSARI', 1462, 77, 6),
		(18236, 'SUMBER PAKEM', 1462, 77, 6),
		(18237, 'SUMBERJAMBE', 1462, 77, 6),
		(18238, 'ANTIROGO', 1463, 77, 6),
		(18239, 'KARANGREJO', 1463, 77, 6),
		(18240, 'KEBONSARI', 1463, 77, 6),
		(18241, 'KERANJINGAN', 1463, 77, 6),
		(18242, 'SUMBERSARI', 1463, 77, 6),
		(18243, 'TEGAL GEDE', 1463, 77, 6),
		(18244, 'WIROLEGI', 1463, 77, 6),
		(18245, 'DARUNGAN', 1464, 77, 6),
		(18246, 'KLATAKAN', 1464, 77, 6),
		(18247, 'KRAMAT SUKOHARJO', 1464, 77, 6),
		(18248, 'MANGGISAN', 1464, 77, 6),
		(18249, 'PATEMON', 1464, 77, 6),
		(18250, 'SELODAKON', 1464, 77, 6),
		(18251, 'TANGGUL KULON', 1464, 77, 6),
		(18252, 'TANGGUL WETAN', 1464, 77, 6),
		(18253, 'ANDONGREJO', 1465, 77, 6),
		(18254, 'CURAHNONGKO', 1465, 77, 6),
		(18255, 'CURAHTAKIR', 1465, 77, 6),
		(18256, 'PONDOKREJO', 1465, 77, 6),
		(18257, 'SANENREJO', 1465, 77, 6),
		(18258, 'SIDODADI', 1465, 77, 6),
		(18259, 'TEMPUREJO', 1465, 77, 6),
		(18260, 'WONOASRI', 1465, 77, 6),
		(18261, 'GADINGREJO', 1466, 77, 6),
		(18262, 'GUNUNGSARI', 1466, 77, 6),
		(18263, 'MUNDUREJO', 1466, 77, 6),
		(18264, 'PALERAN', 1466, 77, 6),
		(18265, 'SIDOREJO', 1466, 77, 6),
		(18266, 'SUKORENO', 1466, 77, 6),
		(18267, 'TANJUNGSARI', 1466, 77, 6),
		(18268, 'TEGAL WANGI', 1466, 77, 6),
		(18269, 'UMBULREJO', 1466, 77, 6),
		(18270, 'UMBULSARI', 1466, 77, 6),
		(18271, 'AMPEL', 1467, 77, 6),
		(18272, 'DUKUH DEMPOK', 1467, 77, 6),
		(18273, 'GLUNDENGAN', 1467, 77, 6),
		(18274, 'KESILIR', 1467, 77, 6),
		(18275, 'LOJEJER', 1467, 77, 6),
		(18276, 'TAMANSARI', 1467, 77, 6),
		(18277, 'TANJUNG REJO', 1467, 77, 6),
		(18278, 'BANDAR KEDUNG MULYO', 1468, 78, 6),
		(18279, 'BANJAR SARI', 1468, 78, 6),
		(18280, 'BARONG SAWAHAN', 1468, 78, 6),
		(18281, 'BRANGKAL', 1468, 78, 6),
		(18282, 'BRODOT', 1468, 78, 6),
		(18283, 'GONDANG MANIS', 1468, 78, 6),
		(18284, 'KARANG DAGANGAN', 1468, 78, 6),
		(18285, 'KAYEN', 1468, 78, 6),
		(18286, 'MOJOKAMBANG', 1468, 78, 6),
		(18287, 'PUCANG SIMO', 1468, 78, 6),
		(18288, 'TINGGAR', 1468, 78, 6),
		(18289, 'BANJARAGUNG', 1469, 78, 6),
		(18290, 'BARENG', 1469, 78, 6),
		(18291, 'JENISGELARAN', 1469, 78, 6),
		(18292, 'KARANGAN', 1469, 78, 6),
		(18293, 'KEBONDALEM', 1469, 78, 6),
		(18294, 'MOJOTENGAH', 1469, 78, 6),
		(18295, 'MUNDUSEWU', 1469, 78, 6),
		(18296, 'NGAMPUNGAN', 1469, 78, 6),
		(18297, 'NGLEBAK', 1469, 78, 6),
		(18298, 'NGRIMBI', 1469, 78, 6),
		(18299, 'PAKEL', 1469, 78, 6),
		(18300, 'PULOSARI', 1469, 78, 6),
		(18301, 'TEBEL', 1469, 78, 6),
		(18302, 'BALONG BESUK', 1470, 78, 6),
		(18303, 'BANDUNG', 1470, 78, 6),
		(18304, 'BENDET', 1470, 78, 6),
		(18305, 'BRAMBANG', 1470, 78, 6),
		(18306, 'BULUREJO', 1470, 78, 6),
		(18307, 'CEWENG', 1470, 78, 6),
		(18308, 'CUKIR', 1470, 78, 6),
		(18309, 'DIWEK', 1470, 78, 6),
		(18310, 'GROGOL', 1470, 78, 6),
		(18311, 'JATIPELEM', 1470, 78, 6),
		(18312, 'JATIREJO', 1470, 78, 6),
		(18313, 'KAYANGAN', 1470, 78, 6),
		(18314, 'KEDAWONG', 1470, 78, 6),
		(18315, 'KERAS', 1470, 78, 6),
		(18316, 'KWARON', 1470, 78, 6),
		(18317, 'NGUDIREJO', 1470, 78, 6),
		(18318, 'PANDANWANGI', 1470, 78, 6),
		(18319, 'PUNDONG', 1470, 78, 6),
		(18320, 'PUTON', 1470, 78, 6),
		(18321, 'WATUGALUH', 1470, 78, 6),
		(18322, 'BLIMBING', 1471, 78, 6),
		(18323, 'BUGASUR KEDALEMAN', 1471, 78, 6),
		(18324, 'GEMPOL LEGUNDI', 1471, 78, 6),
		(18325, 'GODONG', 1471, 78, 6),
		(18326, 'GUDO', 1471, 78, 6),
		(18327, 'JAPANAN', 1471, 78, 6),
		(18328, 'KEDUNGTURI', 1471, 78, 6),
		(18329, 'KREMBANGAN', 1471, 78, 6),
		(18330, 'MEJOYOLOSARI', 1471, 78, 6),
		(18331, 'MENTAOS', 1471, 78, 6),
		(18332, 'PESANGGRAHAN', 1471, 78, 6),
		(18333, 'PLUMBON GAMBANG', 1471, 78, 6),
		(18334, 'PUCANGRO', 1471, 78, 6),
		(18335, 'SEPANYUL', 1471, 78, 6),
		(18336, 'SUKOIBER', 1471, 78, 6),
		(18337, 'SUKOPINGGIR', 1471, 78, 6),
		(18338, 'TANGGUNGAN', 1471, 78, 6),
		(18339, 'WANGKALKEPUH', 1471, 78, 6),
		(18340, 'ALANG ALANG CARUBAN', 1472, 78, 6),
		(18341, 'JANTI', 1472, 78, 6),
		(18342, 'JARAK KULON', 1472, 78, 6),
		(18343, 'JOGOROTO', 1472, 78, 6),
		(18344, 'MAYANGAN', 1472, 78, 6),
		(18345, 'NGUMPUL', 1472, 78, 6),
		(18346, 'SAMBIREJO', 1472, 78, 6),
		(18347, 'SAWIJI', 1472, 78, 6),
		(18348, 'SUKOSARI', 1472, 78, 6),
		(18349, 'SUMBER MULYO', 1472, 78, 6),
		(18350, 'TAMBAR', 1472, 78, 6),
		(18351, 'BANJAR DOWO', 1473, 78, 6),
		(18352, 'GENENGAN JASEM', 1473, 78, 6),
		(18353, 'KABUH', 1473, 78, 6),
		(18354, 'KARANG PAKIS', 1473, 78, 6),
		(18355, 'KAUMAN', 1473, 78, 6),
		(18356, 'KEDUNG JATI', 1473, 78, 6),
		(18357, 'MANDURO', 1473, 78, 6),
		(18358, 'MANGUNAN', 1473, 78, 6),
		(18359, 'MARMOYO', 1473, 78, 6),
		(18360, 'MUNUNG KEREP', 1473, 78, 6),
		(18361, 'PENGAMPON', 1473, 78, 6),
		(18362, 'SUKODADI', 1473, 78, 6),
		(18363, 'SUMBER GONDANG', 1473, 78, 6),
		(18364, 'SUMBERAJI', 1473, 78, 6),
		(18365, 'SUMBERRINGIN', 1473, 78, 6),
		(18366, 'TANJUNG WADUNG', 1473, 78, 6),
		(18367, 'BAKALANRAYUNG', 1474, 78, 6),
		(18368, 'BENDUNGAN', 1474, 78, 6),
		(18369, 'KATEMAS', 1474, 78, 6),
		(18370, 'KEPUHREJO', 1474, 78, 6),
		(18371, 'KUDUBANJAR', 1474, 78, 6),
		(18372, 'MADE', 1474, 78, 6),
		(18373, 'MENTURUS', 1474, 78, 6),
		(18374, 'RANDUWATANG', 1474, 78, 6),
		(18375, 'SIDOKATON', 1474, 78, 6),
		(18376, 'SUMBER TEGUH', 1474, 78, 6),
		(18377, 'TAPEN', 1474, 78, 6),
		(18378, 'BALONGGEMEK', 1475, 78, 6),
		(18379, 'BALONGSARI', 1475, 78, 6),
		(18380, 'DUKUH ARUM', 1475, 78, 6),
		(18381, 'GONGSENG', 1475, 78, 6),
		(18382, 'KEDUNG REJO', 1475, 78, 6),
		(18383, 'MEGALUH', 1475, 78, 6),
		(18384, 'NGOGRI', 1475, 78, 6),
		(18385, 'PACAR PELUK', 1475, 78, 6),
		(18386, 'SIDOMULYO', 1475, 78, 6),
		(18387, 'SUDIMORO', 1475, 78, 6),
		(18388, 'SUMBER AGUNG', 1475, 78, 6),
		(18389, 'SUMBERSARI', 1475, 78, 6),
		(18390, 'TURI PINGGIR', 1475, 78, 6),
		(18391, 'BETEK', 1476, 78, 6),
		(18392, 'DUKUH MOJO', 1476, 78, 6),
		(18393, 'DUKUHDIMORO', 1476, 78, 6),
		(18394, 'GAMBIRAN', 1476, 78, 6),
		(18395, 'JANTI', 1476, 78, 6),
		(18396, 'JOHOWINONG', 1476, 78, 6),
		(18397, 'KADEMANGAN', 1476, 78, 6),
		(18398, 'KARANGWINONGAN', 1476, 78, 6),
		(18399, 'KAROBELAH', 1476, 78, 6),
		(18400, 'KAUMAN', 1476, 78, 6),
		(18401, 'KEDUNGLUMPANG', 1476, 78, 6),
		(18402, 'MANCILAN', 1476, 78, 6),
		(18403, 'MIAGAN', 1476, 78, 6),
		(18404, 'MOJOTRISNO', 1476, 78, 6),
		(18405, 'MURUKAN', 1476, 78, 6),
		(18406, 'SEKETI', 1476, 78, 6),
		(18407, 'TANGGALREJO', 1476, 78, 6),
		(18408, 'TEJO', 1476, 78, 6),
		(18409, 'CATAK GAYAM', 1477, 78, 6),
		(18410, 'GEDANGAN', 1477, 78, 6),
		(18411, 'GONDEK', 1477, 78, 6),
		(18412, 'GROBOGAN', 1477, 78, 6),
		(18413, 'JAPANAN', 1477, 78, 6),
		(18414, 'KARANGLO', 1477, 78, 6),
		(18415, 'KEDUNGPARI', 1477, 78, 6),
		(18416, 'LATSARI', 1477, 78, 6),
		(18417, 'MENGANTO', 1477, 78, 6),
		(18418, 'MOJODUWUR', 1477, 78, 6),
		(18419, 'MOJOJEJER', 1477, 78, 6),
		(18420, 'MOJOWANGI', 1477, 78, 6),
		(18421, 'MOJOWARNO', 1477, 78, 6),
		(18422, 'PENGGARON', 1477, 78, 6),
		(18423, 'REJOSLAMET', 1477, 78, 6),
		(18424, 'SELOREJO', 1477, 78, 6),
		(18425, 'SIDOKERTO', 1477, 78, 6),
		(18426, 'SUKOMULYO', 1477, 78, 6),
		(18427, 'WRINGINPITU', 1477, 78, 6),
		(18428, 'BADANG', 1478, 78, 6),
		(18429, 'BANDARASRI', 1478, 78, 6),
		(18430, 'BANYUARANG', 1478, 78, 6),
		(18431, 'CANDIHARJO', 1478, 78, 6),
		(18432, 'GAJAH', 1478, 78, 6),
		(18433, 'GENUKWATU', 1478, 78, 6),
		(18434, 'JASEM', 1478, 78, 6),
		(18435, 'JOMBOK', 1478, 78, 6),
		(18436, 'KAUMAN', 1478, 78, 6),
		(18437, 'KEMBANGSRI', 1478, 78, 6),
		(18438, 'KERTOREJO', 1478, 78, 6),
		(18439, 'KESAMBEN', 1478, 78, 6),
		(18440, 'KESEMEN', 1478, 78, 6),
		(18441, 'KUNJOROWESI', 1478, 78, 6),
		(18442, 'KUTOGIRANG', 1478, 78, 6),
		(18443, 'LOLAWANG', 1478, 78, 6),
		(18444, 'MANDURO MANGGUNG GAJAH', 1478, 78, 6),
		(18445, 'NGORO', 1478, 78, 6),
		(18446, 'NGORO', 1478, 78, 6),
		(18447, 'PULOREJO', 1478, 78, 6),
		(18448, 'PURWOJATI', 1478, 78, 6),
		(18449, 'REJOAGUNG', 1478, 78, 6),
		(18450, 'SEDATI', 1478, 78, 6),
		(18451, 'SIDOWAREK', 1478, 78, 6),
		(18452, 'SRIGADING', 1478, 78, 6),
		(18453, 'SUGIHWARAS', 1478, 78, 6),
		(18454, 'SUKOANYAR', 1478, 78, 6),
		(18455, 'TAMBAKREJO', 1478, 78, 6),
		(18456, 'TANJANGRONO', 1478, 78, 6),
		(18457, 'WATESNEGORO', 1478, 78, 6),
		(18458, 'WONOSARI', 1478, 78, 6),
		(18459, 'WOTANMASJEDONG', 1478, 78, 6),
		(18460, 'ASEM GEDE', 1479, 78, 6),
		(18461, 'CUPAK', 1479, 78, 6),
		(18462, 'KEBOAN', 1479, 78, 6),
		(18463, 'KEDUNG BOGO', 1479, 78, 6),
		(18464, 'KETAPANG KUNING', 1479, 78, 6),
		(18465, 'KROMONG', 1479, 78, 6),
		(18466, 'MANUNGGAL', 1479, 78, 6),
		(18467, 'MOJODANU', 1479, 78, 6),
		(18468, 'NGAMPEL', 1479, 78, 6),
		(18469, 'NGUSIKAN', 1479, 78, 6),
		(18470, 'SUMBER NONGKO', 1479, 78, 6),
		(18471, 'CANGKRINGRANDU', 1480, 78, 6),
		(18472, 'GADINGMANGU', 1480, 78, 6),
		(18473, 'GLAGAHAN', 1480, 78, 6),
		(18474, 'JANTIGANGGONG', 1480, 78, 6),
		(18475, 'KALANG SEMANDING', 1480, 78, 6),
		(18476, 'KEPUHKAJANG', 1480, 78, 6),
		(18477, 'PAGERWOJO', 1480, 78, 6),
		(18478, 'PERAK', 1480, 78, 6),
		(18479, 'PLOSOGENUK', 1480, 78, 6),
		(18480, 'SEMBUNG', 1480, 78, 6),
		(18481, 'SUKOREJO', 1480, 78, 6),
		(18482, 'SUMBERAGUNG', 1480, 78, 6),
		(18483, 'TEMUWULAN', 1480, 78, 6),
		(18484, 'BONGKOT', 1481, 78, 6),
		(18485, 'DUKUH KLOPO', 1481, 78, 6),
		(18486, 'KEBONTEMU', 1481, 78, 6),
		(18487, 'KEPLAKSARI', 1481, 78, 6),
		(18488, 'KEPUHKEMBENG', 1481, 78, 6),
		(18489, 'MANCAR', 1481, 78, 6),
		(18490, 'MOROSUNGGINGAN', 1481, 78, 6),
		(18491, 'NGRANDU LOR', 1481, 78, 6),
		(18492, 'PETERONGAN', 1481, 78, 6),
		(18493, 'SENDEN', 1481, 78, 6),
		(18494, 'SUMBERAGUNG', 1481, 78, 6),
		(18495, 'TANJUNG GUNUNG', 1481, 78, 6),
		(18496, 'TENGARAN', 1481, 78, 6),
		(18497, 'TUGUSUMBERJO', 1481, 78, 6),
		(18498, 'BANGSRI', 1482, 78, 6),
		(18499, 'DARUREJO', 1482, 78, 6),
		(18500, 'GEBANG BUNDER', 1482, 78, 6),
		(18501, 'JATIMLEREK', 1482, 78, 6),
		(18502, 'JIPURAPAH', 1482, 78, 6),
		(18503, 'KAMPUNGBARU', 1482, 78, 6),
		(18504, 'KARANG MOJO', 1482, 78, 6),
		(18505, 'KLITIH', 1482, 78, 6),
		(18506, 'PLABUHAN', 1482, 78, 6),
		(18507, 'PLANDAAN', 1482, 78, 6),
		(18508, 'PURISEMANDING', 1482, 78, 6),
		(18509, 'SUMBERJO', 1482, 78, 6),
		(18510, 'TONDOWULAN', 1482, 78, 6),
		(18511, 'BAWANGAN', 1483, 78, 6),
		(18512, 'DADI TUNGGAL', 1483, 78, 6),
		(18513, 'GEDONGOMBO', 1483, 78, 6),
		(18514, 'JATIBANJAR', 1483, 78, 6),
		(18515, 'JATIGEDONG', 1483, 78, 6),
		(18516, 'KEBON AGUNG', 1483, 78, 6),
		(18517, 'KEDUNGDOWO', 1483, 78, 6),
		(18518, 'LOSARI', 1483, 78, 6),
		(18519, 'PAGER TANJUNG', 1483, 78, 6),
		(18520, 'PANDAN BLOLE', 1483, 78, 6),
		(18521, 'PLOSO', 1483, 78, 6),
		(18522, 'REJOAGUNG', 1483, 78, 6),
		(18523, 'TANGGUNG KRAMAT', 1483, 78, 6),
		(18524, 'BADAS', 1484, 78, 6),
		(18525, 'BAKALAN', 1484, 78, 6),
		(18526, 'BRUDU', 1484, 78, 6),
		(18527, 'BUDUGSIDOREJO', 1484, 78, 6),
		(18528, 'CURAH MALANG', 1484, 78, 6),
		(18529, 'GEDANGAN', 1484, 78, 6),
		(18530, 'JOGOLOYO', 1484, 78, 6),
		(18531, 'KEDUNGPAPAR', 1484, 78, 6),
		(18532, 'KENDALSARI', 1484, 78, 6),
		(18533, 'MADYO PURO', 1484, 78, 6),
		(18534, 'MENTORO', 1484, 78, 6),
		(18535, 'MLARAS', 1484, 78, 6),
		(18536, 'NGLELE', 1484, 78, 6),
		(18537, 'PALREJO', 1484, 78, 6),
		(18538, 'PLEMAHAN', 1484, 78, 6),
		(18539, 'PLOSOKEREP', 1484, 78, 6),
		(18540, 'SEBANI', 1484, 78, 6),
		(18541, 'SEGODOREJO', 1484, 78, 6),
		(18542, 'SUMOBITO', 1484, 78, 6),
		(18543, 'TALUN KIDUL', 1484, 78, 6),
		(18544, 'TRAWASAN', 1484, 78, 6),
		(18545, 'BEDAH LAWAK', 1485, 78, 6),
		(18546, 'GABUSBANARAN', 1485, 78, 6),
		(18547, 'JATI WATES', 1485, 78, 6),
		(18548, 'KALI KEJAMBON', 1485, 78, 6),
		(18549, 'KEDUNG LOSARI', 1485, 78, 6),
		(18550, 'KEDUNG OTOK', 1485, 78, 6),
		(18551, 'KEPUH DOKO', 1485, 78, 6),
		(18552, 'MOJOKRAPAK', 1485, 78, 6),
		(18553, 'PESANTREN', 1485, 78, 6),
		(18554, 'PULO GEDANG', 1485, 78, 6),
		(18555, 'PULOREJO', 1485, 78, 6),
		(18556, 'REJOSO PINGGIR', 1485, 78, 6),
		(18557, 'SENTUL', 1485, 78, 6),
		(18558, 'TAMPING MOJO', 1485, 78, 6),
		(18559, 'TEMBELANG', 1485, 78, 6),
		(18560, 'BOTOREJO', 1486, 78, 6),
		(18561, 'BUNDERAN', 1486, 78, 6),
		(18562, 'CARANGWULUNG', 1486, 78, 6),
		(18563, 'DORENG', 1486, 78, 6),
		(18564, 'GALENGDOWO', 1486, 78, 6),
		(18565, 'GETAS', 1486, 78, 6),
		(18566, 'JARAK', 1486, 78, 6),
		(18567, 'JOGOLOYO', 1486, 78, 6),
		(18568, 'KALIANYAR', 1486, 78, 6),
		(18569, 'KARANGREJO', 1486, 78, 6),
		(18570, 'KARANGROWO', 1486, 78, 6),
		(18571, 'KENDALDOYONG', 1486, 78, 6),
		(18572, 'KERANGKULON', 1486, 78, 6),
		(18573, 'KUNCIR', 1486, 78, 6),
		(18574, 'LEMPUYANG', 1486, 78, 6),
		(18575, 'MOJODEMAK', 1486, 78, 6),
		(18576, 'MRANAK', 1486, 78, 6),
		(18577, 'MRISEN', 1486, 78, 6),
		(18578, 'PANGLUNGAN', 1486, 78, 6),
		(18579, 'PILANGREJO', 1486, 78, 6),
		(18580, 'SAMBIREJO', 1486, 78, 6),
		(18581, 'SIDO MULYO', 1486, 78, 6),
		(18582, 'SUMBERJO', 1486, 78, 6),
		(18583, 'TLOGODOWO', 1486, 78, 6),
		(18584, 'TLOGOREJO', 1486, 78, 6),
		(18585, 'TRENGGULI', 1486, 78, 6),
		(18586, 'WONOKERTO', 1486, 78, 6),
		(18587, 'WONOMERTO', 1486, 78, 6),
		(18588, 'WONOSALAM', 1486, 78, 6),
		(18589, 'WONOSALAM', 1486, 78, 6),
		(18590, 'BADAS', 1487, 79, 6),
		(18591, 'BLARU', 1487, 79, 6),
		(18592, 'BRINGIN', 1487, 79, 6),
		(18593, 'CANGGU', 1487, 79, 6),
		(18594, 'KRECEK', 1487, 79, 6),
		(18595, 'LAMONG', 1487, 79, 6),
		(18596, 'SEKOTO', 1487, 79, 6),
		(18597, 'TUNGLUR', 1487, 79, 6),
		(18598, 'BANYAKAN', 1488, 79, 6),
		(18599, 'JABON', 1488, 79, 6),
		(18600, 'JATIREJO', 1488, 79, 6),
		(18601, 'MANYARAN', 1488, 79, 6),
		(18602, 'MARON', 1488, 79, 6),
		(18603, 'NGABLAK', 1488, 79, 6),
		(18604, 'PARANG', 1488, 79, 6),
		(18605, 'SENDANG', 1488, 79, 6),
		(18606, 'TIRON', 1488, 79, 6),
		(18607, 'GAMPENG (GAMPING)', 1489, 79, 6),
		(18608, 'JONGBIRU', 1489, 79, 6),
		(18609, 'KALIBELO', 1489, 79, 6),
		(18610, 'KEPUHREJO', 1489, 79, 6),
		(18611, 'NGEBRAK', 1489, 79, 6),
		(18612, 'PLOSOREJO', 1489, 79, 6),
		(18613, 'PUTIH', 1489, 79, 6),
		(18614, 'SAMBIREJO', 1489, 79, 6),
		(18615, 'SAMBIRESIK', 1489, 79, 6),
		(18616, 'TURUS', 1489, 79, 6),
		(18617, 'WANENGPATEN', 1489, 79, 6),
		(18618, 'BAKALAN', 1490, 79, 6),
		(18619, 'BANARAN', 1490, 79, 6),
		(18620, 'CEMANI', 1490, 79, 6),
		(18621, 'CERME', 1490, 79, 6),
		(18622, 'DATENGAN', 1490, 79, 6),
		(18623, 'GAMBYOK', 1490, 79, 6),
		(18624, 'GEDANGAN', 1490, 79, 6),
		(18625, 'GROGOL', 1490, 79, 6),
		(18626, 'GROGOL', 1490, 79, 6),
		(18627, 'KADOKAN', 1490, 79, 6),
		(18628, 'KALIPANG', 1490, 79, 6),
		(18629, 'KWARASAN', 1490, 79, 6),
		(18630, 'LANGENHARJO', 1490, 79, 6),
		(18631, 'MADEGONDO', 1490, 79, 6),
		(18632, 'MANANG', 1490, 79, 6),
		(18633, 'PANDEYAN', 1490, 79, 6),
		(18634, 'PARANGJORO', 1490, 79, 6),
		(18635, 'PONDOK', 1490, 79, 6),
		(18636, 'SANGGRAHAN', 1490, 79, 6),
		(18637, 'SONOREJO', 1490, 79, 6),
		(18638, 'SUMBEREJO', 1490, 79, 6),
		(18639, 'TELUKAN', 1490, 79, 6),
		(18640, 'WONOASRI', 1490, 79, 6),
		(18641, 'ADAN-ADAN', 1491, 79, 6),
		(18642, 'BANGKOK', 1491, 79, 6),
		(18643, 'BANYUANYAR', 1491, 79, 6),
		(18644, 'BESUK', 1491, 79, 6),
		(18645, 'BLIMBING', 1491, 79, 6),
		(18646, 'BOGEM', 1491, 79, 6),
		(18647, 'GABRU', 1491, 79, 6),
		(18648, 'GAYAM', 1491, 79, 6),
		(18649, 'GEMPOLAN', 1491, 79, 6),
		(18650, 'GURAH', 1491, 79, 6),
		(18651, 'KERKEP', 1491, 79, 6),
		(18652, 'KRANGGAN', 1491, 79, 6),
		(18653, 'NGASEM', 1491, 79, 6),
		(18654, 'NGLUMBANG', 1491, 79, 6),
		(18655, 'SUKOREJO', 1491, 79, 6),
		(18656, 'SUMBERCANGKRING', 1491, 79, 6),
		(18657, 'TAMBAKREJO', 1491, 79, 6),
		(18658, 'TIRU KIDUL', 1491, 79, 6),
		(18659, 'TIRU LOR', 1491, 79, 6),
		(18660, 'TURUS', 1491, 79, 6),
		(18661, 'WONOJOYO', 1491, 79, 6),
		(18662, 'BLABAK', 1492, 79, 6),
		(18663, 'CENDONO', 1492, 79, 6),
		(18664, 'KANDAT', 1492, 79, 6),
		(18665, 'KARANGREJO', 1492, 79, 6),
		(18666, 'NGLETIH', 1492, 79, 6),
		(18667, 'NGRECO', 1492, 79, 6),
		(18668, 'PULE', 1492, 79, 6),
		(18669, 'PURWOREJO', 1492, 79, 6),
		(18670, 'RINGINSARI', 1492, 79, 6),
		(18671, 'SELOSARI', 1492, 79, 6),
		(18672, 'SUMBEREJO', 1492, 79, 6),
		(18673, 'TEGALAN', 1492, 79, 6),
		(18674, 'BANGSONGAN', 1493, 79, 6),
		(18675, 'BAYE', 1493, 79, 6),
		(18676, 'JAMBU', 1493, 79, 6),
		(18677, 'KAYEN KIDUL', 1493, 79, 6),
		(18678, 'MUKUH', 1493, 79, 6),
		(18679, 'NANGGUNGAN', 1493, 79, 6),
		(18680, 'PADANGAN', 1493, 79, 6),
		(18681, 'SAMBIROBYONG', 1493, 79, 6),
		(18682, 'SEKARAN', 1493, 79, 6),
		(18683, 'SEMAMBUNG', 1493, 79, 6),
		(18684, 'SENDEN', 1493, 79, 6),
		(18685, 'SUKOHARJO', 1493, 79, 6),
		(18686, 'BALOWERTI', 1494, 79, 6),
		(18687, 'BANJARAN', 1494, 79, 6),
		(18688, 'DANDANGAN', 1494, 79, 6),
		(18689, 'JAGALAN', 1494, 79, 6),
		(18690, 'KALIOMBO', 1494, 79, 6),
		(18691, 'KAMPUNG DALEM', 1494, 79, 6),
		(18692, 'KEMASAN', 1494, 79, 6),
		(18693, 'MANISRENGGO', 1494, 79, 6),
		(18694, 'NGADIREJO', 1494, 79, 6),
		(18695, 'NGRONGGO', 1494, 79, 6),
		(18696, 'PAKELAN', 1494, 79, 6),
		(18697, 'POCANAN', 1494, 79, 6),
		(18698, 'REJOMULYO', 1494, 79, 6),
		(18699, 'RINGIN ANOM', 1494, 79, 6),
		(18700, 'SEMAMPIR', 1494, 79, 6),
		(18701, 'SETONO GEDONG', 1494, 79, 6),
		(18702, 'SETONO PANDE', 1494, 79, 6),
		(18703, 'BESOWO', 1495, 79, 6),
		(18704, 'BRUMBUNG', 1495, 79, 6),
		(18705, 'DAMARWULAN', 1495, 79, 6),
		(18706, 'KAMPUNGBARU', 1495, 79, 6),
		(18707, 'KEBONREJO', 1495, 79, 6),
		(18708, 'KELING', 1495, 79, 6),
		(18709, 'KENCONG', 1495, 79, 6),
		(18710, 'KEPUNG', 1495, 79, 6),
		(18711, 'KRENCENG', 1495, 79, 6),
		(18712, 'SIMAN', 1495, 79, 6),
		(18713, 'BANJARANYAR', 1496, 79, 6),
		(18714, 'BENDOSARI', 1496, 79, 6),
		(18715, 'BLEBER', 1496, 79, 6),
		(18716, 'BUTUH', 1496, 79, 6),
		(18717, 'JABANG', 1496, 79, 6),
		(18718, 'JAMBEAN', 1496, 79, 6),
		(18719, 'KANIGORO', 1496, 79, 6),
		(18720, 'KARANGTALUN', 1496, 79, 6),
		(18721, 'KRANDANG', 1496, 79, 6),
		(18722, 'KRAS', 1496, 79, 6),
		(18723, 'MOJOSARI', 1496, 79, 6),
		(18724, 'NYAWANGAN', 1496, 79, 6),
		(18725, 'PELAS', 1496, 79, 6),
		(18726, 'PURWODADI', 1496, 79, 6),
		(18727, 'REJOMULYO', 1496, 79, 6),
		(18728, 'SETONOREJO', 1496, 79, 6),
		(18729, 'BALONGJERUK (BALUNG JERUK)', 1497, 79, 6),
		(18730, 'DUNGUS', 1497, 79, 6),
		(18731, 'JUWET', 1497, 79, 6),
		(18732, 'KAPAS', 1497, 79, 6),
		(18733, 'KAPI', 1497, 79, 6),
		(18734, 'KLEPEK', 1497, 79, 6),
		(18735, 'KUNJANG', 1497, 79, 6),
		(18736, 'KUWIK', 1497, 79, 6),
		(18737, 'PAKIS', 1497, 79, 6),
		(18738, 'PARELOR', 1497, 79, 6),
		(18739, 'TENGGERLOR', 1497, 79, 6),
		(18740, 'WONOREJO', 1497, 79, 6),
		(18741, 'BLIMBING', 1498, 79, 6),
		(18742, 'JUGO', 1498, 79, 6),
		(18743, 'KEDAWUNG', 1498, 79, 6),
		(18744, 'KENITEN', 1498, 79, 6),
		(18745, 'KRANDING', 1498, 79, 6),
		(18746, 'KRATON', 1498, 79, 6),
		(18747, 'MAESAN', 1498, 79, 6),
		(18748, 'MLATI', 1498, 79, 6),
		(18749, 'MOJO', 1498, 79, 6),
		(18750, 'MONDO', 1498, 79, 6),
		(18751, 'NGADI', 1498, 79, 6),
		(18752, 'NGETREP', 1498, 79, 6),
		(18753, 'PAMONGAN', 1498, 79, 6),
		(18754, 'PETOK', 1498, 79, 6),
		(18755, 'PETUNGROTO', 1498, 79, 6),
		(18756, 'PLOSO', 1498, 79, 6),
		(18757, 'PONGGOK', 1498, 79, 6),
		(18758, 'SUKOANYAR', 1498, 79, 6),
		(18759, 'SURAT', 1498, 79, 6),
		(18760, 'TAMBIBENDO', 1498, 79, 6),
		(18761, 'BANDAR KIDUL', 1499, 79, 6),
		(18762, 'BANDAR LOR', 1499, 79, 6),
		(18763, 'BANJAR MELATI (BANJARMLATI)', 1499, 79, 6),
		(18764, 'BUJEL', 1499, 79, 6),
		(18765, 'CAMPUREJO', 1499, 79, 6),
		(18766, 'DERMO', 1499, 79, 6),
		(18767, 'GAYAM', 1499, 79, 6),
		(18768, 'LIRBOYO', 1499, 79, 6),
		(18769, 'MOJOROTO', 1499, 79, 6),
		(18770, 'MRICAN', 1499, 79, 6),
		(18771, 'NGAMPEL', 1499, 79, 6),
		(18772, 'POJOK', 1499, 79, 6),
		(18773, 'SUKORAME', 1499, 79, 6),
		(18774, 'TAMANAN', 1499, 79, 6),
		(18775, 'BADAL', 1500, 79, 6),
		(18776, 'BADALPANDEAN', 1500, 79, 6),
		(18777, 'BANGGLE', 1500, 79, 6),
		(18778, 'BANJAREJO', 1500, 79, 6),
		(18779, 'BEDUG', 1500, 79, 6),
		(18780, 'BRANGGAHAN', 1500, 79, 6),
		(18781, 'DUKUH', 1500, 79, 6),
		(18782, 'MANGUNREJO', 1500, 79, 6),
		(18783, 'NGADILUWIH', 1500, 79, 6),
		(18784, 'PURWOKERTO', 1500, 79, 6),
		(18785, 'REMBANG', 1500, 79, 6),
		(18786, 'REMBANGKEPUH', 1500, 79, 6),
		(18787, 'SEKETI', 1500, 79, 6),
		(18788, 'SLUMBUNG', 1500, 79, 6),
		(18789, 'TALES', 1500, 79, 6),
		(18790, 'WONOREJO', 1500, 79, 6),
		(18791, 'BABADAN', 1501, 79, 6),
		(18792, 'BEDALI', 1501, 79, 6),
		(18793, 'JAGUL', 1501, 79, 6),
		(18794, 'KUNJANG', 1501, 79, 6),
		(18795, 'MANGGIS', 1501, 79, 6),
		(18796, 'MARGOURIP', 1501, 79, 6),
		(18797, 'NGANCAR', 1501, 79, 6),
		(18798, 'PANDANTOYO', 1501, 79, 6),
		(18799, 'SEMPU', 1501, 79, 6),
		(18800, 'SUGIHWARAS', 1501, 79, 6),
		(18801, 'BANDUNGREJO', 1502, 79, 6),
		(18802, 'BARENG', 1502, 79, 6),
		(18803, 'BUTOH', 1502, 79, 6),
		(18804, 'DOKO', 1502, 79, 6),
		(18805, 'DUKOHKIDUL', 1502, 79, 6),
		(18806, 'GOGORANTE', 1502, 79, 6),
		(18807, 'JAMPET', 1502, 79, 6),
		(18808, 'JELU', 1502, 79, 6),
		(18809, 'KARANGREJO', 1502, 79, 6),
		(18810, 'KOLONG', 1502, 79, 6),
		(18811, 'KWADUNGAN', 1502, 79, 6),
		(18812, 'MEDIYUNAN', 1502, 79, 6),
		(18813, 'NAMBAAN', 1502, 79, 6),
		(18814, 'NGADILUWIH', 1502, 79, 6),
		(18815, 'NGANTRU', 1502, 79, 6),
		(18816, 'NGASEM', 1502, 79, 6),
		(18817, 'NGASEM', 1502, 79, 6),
		(18818, 'PARON', 1502, 79, 6),
		(18819, 'SAMBONG', 1502, 79, 6),
		(18820, 'SENDANGHARJO', 1502, 79, 6),
		(18821, 'SETREN', 1502, 79, 6),
		(18822, 'SUKOREJO', 1502, 79, 6),
		(18823, 'SUMBERJO', 1502, 79, 6),
		(18824, 'TENGGER', 1502, 79, 6),
		(18825, 'TOYORESMI', 1502, 79, 6),
		(18826, 'TRENGGULUNAN', 1502, 79, 6),
		(18827, 'TUGUREJO', 1502, 79, 6),
		(18828, 'WADANG', 1502, 79, 6),
		(18829, 'WONOCATUR', 1502, 79, 6),
		(18830, 'BENDO', 1503, 79, 6),
		(18831, 'BULUPASAR', 1503, 79, 6),
		(18832, 'JAGUNG', 1503, 79, 6),
		(18833, 'KAMBINGAN', 1503, 79, 6),
		(18834, 'MENANG', 1503, 79, 6),
		(18835, 'PAGU', 1503, 79, 6),
		(18836, 'SEMANDING', 1503, 79, 6),
		(18837, 'SEMEN', 1503, 79, 6),
		(18838, 'SITIMERT0', 1503, 79, 6),
		(18839, 'TANJUNG', 1503, 79, 6),
		(18840, 'TENGGER KIDUL', 1503, 79, 6),
		(18841, 'WATES', 1503, 79, 6),
		(18842, 'WONOSARI', 1503, 79, 6),
		(18843, 'DAWUHAN KIDUL', 1504, 79, 6),
		(18844, 'JAMBANGAN', 1504, 79, 6),
		(18845, 'JANTI', 1504, 79, 6),
		(18846, 'KEDUNGMALANG', 1504, 79, 6),
		(18847, 'KEPUH', 1504, 79, 6),
		(18848, 'KWARON', 1504, 79, 6),
		(18849, 'MADURETNO', 1504, 79, 6),
		(18850, 'MINGGIRAN', 1504, 79, 6),
		(18851, 'NGAMPEL', 1504, 79, 6),
		(18852, 'PAPAR', 1504, 79, 6),
		(18853, 'PEHKULON', 1504, 79, 6),
		(18854, 'PEHWETAN', 1504, 79, 6),
		(18855, 'PUHJAJAR', 1504, 79, 6),
		(18856, 'PURWOTENGAH', 1504, 79, 6),
		(18857, 'SRIKATON', 1504, 79, 6),
		(18858, 'SUKOMORO', 1504, 79, 6),
		(18859, 'TANON', 1504, 79, 6),
		(18860, 'BENDO', 1505, 79, 6),
		(18861, 'DARUNGAN', 1505, 79, 6),
		(18862, 'GEDANGSEWU', 1505, 79, 6),
		(18863, 'PARE', 1505, 79, 6),
		(18864, 'PELEM', 1505, 79, 6),
		(18865, 'SAMBIREJO', 1505, 79, 6),
		(18866, 'SIDOREJO', 1505, 79, 6),
		(18867, 'SUMBERBENDO', 1505, 79, 6),
		(18868, 'TERTEK', 1505, 79, 6),
		(18869, 'TULUNGREJO', 1505, 79, 6),
		(18870, 'BANARAN', 1506, 79, 6),
		(18871, 'BANGSAL', 1506, 79, 6),
		(18872, 'BAWANG', 1506, 79, 6),
		(18873, 'BETET', 1506, 79, 6),
		(18874, 'BLABAK', 1506, 79, 6),
		(18875, 'BURENGAN', 1506, 79, 6),
		(18876, 'JAMSAREN', 1506, 79, 6),
		(18877, 'KETAMI', 1506, 79, 6),
		(18878, 'NGLETIH', 1506, 79, 6),
		(18879, 'PAKUNDEN', 1506, 79, 6),
		(18880, 'PESANTREN', 1506, 79, 6),
		(18881, 'SINGONEGARAN', 1506, 79, 6),
		(18882, 'TEMPUREJO', 1506, 79, 6),
		(18883, 'TINALAN', 1506, 79, 6),
		(18884, 'TOSAREN', 1506, 79, 6),
		(18885, 'BANJAREJO', 1507, 79, 6),
		(18886, 'BOGOKIDUL', 1507, 79, 6),
		(18887, 'KAYEN LOR', 1507, 79, 6),
		(18888, 'LANGENHARJO', 1507, 79, 6),
		(18889, 'MEJONO', 1507, 79, 6),
		(18890, 'MOJOAYU', 1507, 79, 6),
		(18891, 'MOJOKEREP', 1507, 79, 6),
		(18892, 'NGINO', 1507, 79, 6),
		(18893, 'PAYAMAN', 1507, 79, 6),
		(18894, 'PLEMAHAN', 1507, 79, 6),
		(18895, 'PUHJARAK', 1507, 79, 6),
		(18896, 'RINGINPITU', 1507, 79, 6),
		(18897, 'SEBET', 1507, 79, 6),
		(18898, 'SIDOWAREK', 1507, 79, 6),
		(18899, 'SUKOHARJO', 1507, 79, 6),
		(18900, 'TEGOWANGI', 1507, 79, 6),
		(18901, 'WONOKERTO', 1507, 79, 6),
		(18902, 'BRENGGOLO', 1508, 79, 6),
		(18903, 'DONGANTI', 1508, 79, 6),
		(18904, 'GONDANG', 1508, 79, 6),
		(18905, 'JARAK', 1508, 79, 6),
		(18906, 'KAWEDUSAN', 1508, 79, 6),
		(18907, 'KAYUNAN', 1508, 79, 6),
		(18908, 'KLANDERAN', 1508, 79, 6),
		(18909, 'PANJER', 1508, 79, 6),
		(18910, 'PLOSOKIDUL', 1508, 79, 6),
		(18911, 'PLOSOLOR', 1508, 79, 6),
		(18912, 'PRANGGANG', 1508, 79, 6),
		(18913, 'PUNJUL', 1508, 79, 6),
		(18914, 'SEPAWON', 1508, 79, 6),
		(18915, 'SUMBERAGUNG', 1508, 79, 6),
		(18916, 'WONOREJO TRISULO', 1508, 79, 6),
		(18917, 'ASMOROBANGUN', 1509, 79, 6),
		(18918, 'GADUNGAN', 1509, 79, 6),
		(18919, 'MANGGIS', 1509, 79, 6),
		(18920, 'PUNCU', 1509, 79, 6),
		(18921, 'SATAK', 1509, 79, 6),
		(18922, 'SIDOMULYO', 1509, 79, 6),
		(18923, 'WATUGEDE', 1509, 79, 6),
		(18924, 'WONOREJO', 1509, 79, 6),
		(18925, 'BELOR', 1510, 79, 6),
		(18926, 'BLAWE', 1510, 79, 6),
		(18927, 'BULU', 1510, 79, 6),
		(18928, 'DAWUHAN', 1510, 79, 6),
		(18929, 'DAYU', 1510, 79, 6),
		(18930, 'JANTOK', 1510, 79, 6),
		(18931, 'KARANGPAKIS', 1510, 79, 6),
		(18932, 'KEMPLENG', 1510, 79, 6),
		(18933, 'KETAWANG', 1510, 79, 6),
		(18934, 'KLAMPITAN', 1510, 79, 6),
		(18935, 'MEKIKIS', 1510, 79, 6),
		(18936, 'MERJOYO', 1510, 79, 6),
		(18937, 'MRANGGEN', 1510, 79, 6),
		(18938, 'MUNENG', 1510, 79, 6),
		(18939, 'PANDANSARI', 1510, 79, 6),
		(18940, 'PESING', 1510, 79, 6),
		(18941, 'PURWOASRI', 1510, 79, 6),
		(18942, 'PURWODADI', 1510, 79, 6),
		(18943, 'SIDOMULYO', 1510, 79, 6),
		(18944, 'SUMBERJO', 1510, 79, 6),
		(18945, 'TUGU', 1510, 79, 6),
		(18946, 'WONOTENGAH', 1510, 79, 6),
		(18947, 'WOROMARTO', 1510, 79, 6),
		(18948, 'BATUAJI', 1511, 79, 6),
		(18949, 'DAWUNG', 1511, 79, 6),
		(18950, 'DEYENG', 1511, 79, 6),
		(18951, 'JEMEKAN', 1511, 79, 6),
		(18952, 'NAMBAAN', 1511, 79, 6),
		(18953, 'PURWODADI', 1511, 79, 6),
		(18954, 'RINGINREJO', 1511, 79, 6),
		(18955, 'SAMBI', 1511, 79, 6),
		(18956, 'SELODONO', 1511, 79, 6),
		(18957, 'SRIKATON', 1511, 79, 6),
		(18958, 'SUSUHBANGO', 1511, 79, 6),
		(18959, 'BOBANG', 1512, 79, 6),
		(18960, 'BULU', 1512, 79, 6),
		(18961, 'JOHO', 1512, 79, 6),
		(18962, 'KANYORAN', 1512, 79, 6),
		(18963, 'KEDAK', 1512, 79, 6),
		(18964, 'PAGUNG', 1512, 79, 6),
		(18965, 'PUHRUBUH', 1512, 79, 6),
		(18966, 'PUHSARANG', 1512, 79, 6),
		(18967, 'SELOPANGGUNG', 1512, 79, 6),
		(18968, 'SEMEN', 1512, 79, 6),
		(18969, 'SIDOMULYO', 1512, 79, 6),
		(18970, 'TITIK', 1512, 79, 6),
		(18971, 'BLIMBING', 1513, 79, 6),
		(18972, 'BULUSARI', 1513, 79, 6),
		(18973, 'CENGKOK', 1513, 79, 6),
		(18974, 'JATI', 1513, 79, 6),
		(18975, 'KALIBOTO', 1513, 79, 6),
		(18976, 'KALIRONG', 1513, 79, 6),
		(18977, 'KEDUNGSARI', 1513, 79, 6),
		(18978, 'KEREP', 1513, 79, 6),
		(18979, 'SUMBERDUREN', 1513, 79, 6),
		(18980, 'TAROKAN', 1513, 79, 6),
		(18981, 'BABAT', 1514, 80, 6),
		(18982, 'BANARAN', 1514, 80, 6),
		(18983, 'BEDAHAN', 1514, 80, 6),
		(18984, 'BULUMARGI', 1514, 80, 6),
		(18985, 'DATINAWONG', 1514, 80, 6),
		(18986, 'GEMBONG', 1514, 80, 6),
		(18987, 'GENDONG KULON', 1514, 80, 6),
		(18988, 'KARANGKEMBANG', 1514, 80, 6),
		(18989, 'KEBALANDONO', 1514, 80, 6),
		(18990, 'KEBALANPELANG', 1514, 80, 6),
		(18991, 'KEBONAGUNG', 1514, 80, 6),
		(18992, 'KEYONGAN', 1514, 80, 6),
		(18993, 'KURIPAN', 1514, 80, 6),
		(18994, 'MOROPELANG', 1514, 80, 6),
		(18995, 'PATIHAN', 1514, 80, 6),
		(18996, 'PLAOSAN', 1514, 80, 6),
		(18997, 'PUCAKWANGI', 1514, 80, 6),
		(18998, 'SAMBANGAN', 1514, 80, 6),
		(18999, 'SOGO', 1514, 80, 6),
		(19000, 'SUMUR GENUK', 1514, 80, 6),
	`)
}
