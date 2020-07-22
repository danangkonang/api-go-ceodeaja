package kel70

import "github.com/danangkonang/rest-api/config"

func Kel79() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
	(79001, 'TELI', 6472, 467, 34),
	(79002, 'TURWE', 6472, 467, 34),
	(79003, 'ARINA', 6473, 467, 34),
	(79004, 'ARINGGON', 6473, 467, 34),
	(79005, 'BORDAMBAN', 6473, 467, 34),
	(79006, 'BORME', 6473, 467, 34),
	(79007, 'BUKAM', 6473, 467, 34),
	(79008, 'CANGDAMBAN', 6473, 467, 34),
	(79009, 'KOLGIR', 6473, 467, 34),
	(79010, 'KWIME', 6473, 467, 34),
	(79011, 'LAYDAMBAN', 6473, 467, 34),
	(79012, 'OMBAN', 6473, 467, 34),
	(79013, 'ONYA', 6473, 467, 34),
	(79014, 'SEBAN', 6473, 467, 34),
	(79015, 'SIKIBUR', 6473, 467, 34),
	(79016, 'BARAMA', 6474, 467, 34),
	(79017, 'BASIRINGE', 6474, 467, 34),
	(79018, 'BUNYIRYE', 6474, 467, 34),
	(79019, 'EIPUMEK / ELPOMEK (ELPUNA)', 6474, 467, 34),
	(79020, 'KWEREDALA', 6474, 467, 34),
	(79021, 'LALAKON', 6474, 467, 34),
	(79022, 'LONDININ', 6474, 467, 34),
	(79023, 'MALINGDAM', 6474, 467, 34),
	(79024, 'MUNGKONA', 6474, 467, 34),
	(79025, 'SERABUM', 6474, 467, 34),
	(79026, 'SUPLEYU', 6474, 467, 34),
	(79027, 'TALEMU', 6474, 467, 34),
	(79028, 'TANIME', 6474, 467, 34),
	(79029, 'WAKIDAM', 6474, 467, 34),
	(79030, 'DINMOT ARIM', 6475, 467, 34),
	(79031, 'DIPOL', 6475, 467, 34),
	(79032, 'EWENKATOP', 6475, 467, 34),
	(79033, 'IWUR', 6475, 467, 34),
	(79034, 'KAMYOIM', 6475, 467, 34),
	(79035, 'KURUKIM (KURUMKIM/KURUMKLIN)', 6475, 467, 34),
	(79036, 'NARNGER', 6475, 467, 34),
	(79037, 'NENGINUM', 6475, 467, 34),
	(79038, 'ULKUBI', 6475, 467, 34),
	(79039, 'WALAPKUBUN', 6475, 467, 34),
	(79040, 'BRIS', 6476, 467, 34),
	(79041, 'JETFA', 6476, 467, 34),
	(79042, 'KALIMBU', 6476, 467, 34),
	(79043, 'LULIS', 6476, 467, 34),
	(79044, 'TUPALMA DUA', 6476, 467, 34),
	(79045, 'TUPALMA SATU', 6476, 467, 34),
	(79046, 'ARINKOP', 6477, 467, 34),
	(79047, 'DABOLDING', 6477, 467, 34),
	(79048, 'IMIK', 6477, 467, 34),
	(79049, 'KALOMDOL (KALONDOI)', 6477, 467, 34),
	(79050, 'TULO', 6477, 467, 34),
	(79051, 'ARINTAP', 6478, 467, 34),
	(79052, 'ATER', 6478, 467, 34),
	(79053, 'KAWOR', 6478, 467, 34),
	(79054, 'NANUM', 6478, 467, 34),
	(79055, 'SAKUP', 6478, 467, 34),
	(79056, 'TARNGOP', 6478, 467, 34),
	(79057, 'UMDING', 6478, 467, 34),
	(79058, 'ASUA', 6479, 467, 34),
	(79059, 'BERUSAHA', 6479, 467, 34),
	(79060, 'DELPEM', 6479, 467, 34),
	(79061, 'DIIP', 6479, 467, 34),
	(79062, 'KIWI', 6479, 467, 34),
	(79063, 'KUKIHIL', 6479, 467, 34),
	(79064, 'LOLIM', 6479, 467, 34),
	(79065, 'MANGOLDOKI', 6479, 467, 34),
	(79066, 'OKNANGGUL', 6479, 467, 34),
	(79067, 'PELEBIB', 6479, 467, 34),
	(79068, 'POMDING', 6479, 467, 34),
	(79069, 'SOPAMIKMA', 6479, 467, 34),
	(79070, 'DIKDON', 6480, 467, 34),
	(79071, 'EHIPTEM (EHIPTEN/OHIPTEM)', 6480, 467, 34),
	(79072, 'OKETUR', 6480, 467, 34),
	(79073, 'OKHIK (OKLIP)', 6480, 467, 34),
	(79074, 'OKYAKO', 6480, 467, 34),
	(79075, 'OKYAOP', 6480, 467, 34),
	(79076, 'OKYOP (OKYOB/OKYAOP)', 6480, 467, 34),
	(79077, 'TATAM', 6480, 467, 34),
	(79078, 'WANTEM (WANTEN)', 6480, 467, 34),
	(79079, 'MILKI', 6481, 467, 34),
	(79080, 'MOT / MOOT', 6481, 467, 34),
	(79081, 'MUARA ASBI', 6481, 467, 34),
	(79082, 'TUAL', 6481, 467, 34),
	(79083, 'YUBU', 6481, 467, 34),
	(79084, 'BIAS', 6482, 467, 34),
	(79085, 'BUMI', 6482, 467, 34),
	(79086, 'DELEMO', 6482, 467, 34),
	(79087, 'TERO', 6482, 467, 34),
	(79088, 'CANGPALLY', 6483, 467, 34),
	(79089, 'KWARBAN', 6483, 467, 34),
	(79090, 'NONGME', 6483, 467, 34),
	(79091, 'OMTAMUR', 6483, 467, 34),
	(79092, 'RUBOL', 6483, 467, 34),
	(79093, 'YARIGON', 6483, 467, 34),
	(79094, 'YOKOM', 6483, 467, 34),
	(79095, 'APLIM', 6484, 467, 34),
	(79096, 'BULANGKOP', 6484, 467, 34),
	(79097, 'KUNGULDING', 6484, 467, 34),
	(79098, 'LIMARUM', 6484, 467, 34),
	(79099, 'OKDO', 6484, 467, 34),
	(79100, 'YUMAKOT', 6484, 467, 34),
	(79101, 'ATEMBABOL', 6485, 467, 34),
	(79102, 'BORBAN', 6485, 467, 34),
	(79103, 'DUMPASI', 6485, 467, 34),
	(79104, 'KIRIMU', 6485, 467, 34),
	(79105, 'MAKSUM', 6485, 467, 34),
	(79106, 'MARKOM', 6485, 467, 34),
	(79107, 'OMLIOM', 6485, 467, 34),
	(79108, 'PEDAM', 6485, 467, 34),
	(79109, 'PENELI', 6485, 467, 34),
	(79110, 'SABIN', 6485, 467, 34),
	(79111, 'TUPOPLYOM', 6485, 467, 34),
	(79112, 'YAPIL', 6485, 467, 34),
	(79113, 'AKMER', 6486, 467, 34),
	(79114, 'BAPE', 6486, 467, 34),
	(79115, 'IBOT', 6486, 467, 34),
	(79116, 'KASAWI', 6486, 467, 34),
	(79117, 'MASIM', 6486, 467, 34),
	(79118, 'TAPASIK DUA', 6486, 467, 34),
	(79119, 'ATANG DOKI', 6487, 467, 34),
	(79120, 'BEMHIMIKU', 6487, 467, 34),
	(79121, 'BILIP BAYO', 6487, 467, 34),
	(79122, 'KAEP', 6487, 467, 34),
	(79123, 'OKALUT KUMAL', 6487, 467, 34),
	(79124, 'OKBEM', 6487, 467, 34),
	(79125, 'OKNGAM', 6487, 467, 34),
	(79126, 'OKTARU', 6487, 467, 34),
	(79127, 'OKTAU', 6487, 467, 34),
	(79128, 'ABMISIBIL (APMISIBIL)', 6488, 467, 34),
	(79129, 'ATOLDOL (ATOLBOL)', 6488, 467, 34),
	(79130, 'IRIDING', 6488, 467, 34),
	(79131, 'MANUNGGAL (OKIPUR)', 6488, 467, 34),
	(79132, 'OKAPLO', 6488, 467, 34),
	(79133, 'OKBIFISIL (OKBIFISIS)', 6488, 467, 34),
	(79134, 'OKSEMAR', 6488, 467, 34),
	(79135, 'OKTANGLAP', 6488, 467, 34),
	(79136, 'KOTYOBAKON', 6489, 467, 34),
	(79137, 'OKELWEL', 6489, 467, 34),
	(79138, 'OKTENENG', 6489, 467, 34),
	(79139, 'TENGNONG', 6489, 467, 34),
	(79140, 'KOMOK', 6490, 467, 34),
	(79141, 'OKAMIN', 6490, 467, 34),
	(79142, 'OKBUMUL', 6490, 467, 34),
	(79143, 'OKHIM', 6490, 467, 34),
	(79144, 'OKLIP', 6490, 467, 34),
	(79145, 'OKTEM', 6490, 467, 34),
	(79146, 'OKTUMI', 6490, 467, 34),
	(79147, 'AUTPAHIK', 6491, 467, 34),
	(79148, 'BOMDING', 6491, 467, 34),
	(79149, 'HONKUDING', 6491, 467, 34),
	(79150, 'OKDILAM', 6491, 467, 34),
	(79151, 'OKDUNAM', 6491, 467, 34),
	(79152, 'OKHAKA', 6491, 467, 34),
	(79153, 'OKMA', 6491, 467, 34),
	(79154, 'OKPA', 6491, 467, 34),
	(79155, 'OKTAE', 6491, 467, 34),
	(79156, 'PAUNE', 6491, 467, 34),
	(79157, 'TINIBIL', 6491, 467, 34),
	(79158, 'TOMKA', 6491, 467, 34),
	(79159, 'KUBIPHKOP (KUBIBKOP)', 6492, 467, 34),
	(79160, 'MANGABIP', 6492, 467, 34),
	(79161, 'OKANO', 6492, 467, 34),
	(79162, 'SEBUL', 6492, 467, 34),
	(79163, 'AKMAKOT', 6493, 467, 34),
	(79164, 'ALDOM', 6493, 467, 34),
	(79165, 'BUNAMDOL', 6493, 467, 34),
	(79166, 'KABIDING (BETAABIB)', 6493, 467, 34),
	(79167, 'KUTDOL', 6493, 467, 34),
	(79168, 'MABILABOL', 6493, 467, 34),
	(79169, 'MOLBIB SILIBIB', 6493, 467, 34),
	(79170, 'POLSAM', 6493, 467, 34),
	(79171, 'ALUTBAKON', 6494, 467, 34),
	(79172, 'ATENOR / ATENAR', 6494, 467, 34),
	(79173, 'MIMIN', 6494, 467, 34),
	(79174, 'OKSOP', 6494, 467, 34),
	(79175, 'OKTUMI', 6494, 467, 34),
	(79176, 'BARAMIRYE', 6495, 467, 34),
	(79177, 'BARICE', 6495, 467, 34),
	(79178, 'IMDE', 6495, 467, 34),
	(79179, 'KALEK', 6495, 467, 34),
	(79180, 'LUMDAKNA', 6495, 467, 34),
	(79181, 'MANDALAK', 6495, 467, 34),
	(79182, 'MARIKLA', 6495, 467, 34),
	(79183, 'PAMEK', 6495, 467, 34),
	(79184, 'PINGGON', 6495, 467, 34),
	(79185, 'YABOSOROM', 6495, 467, 34),
	(79186, 'YOKUL', 6495, 467, 34),
	(79187, 'BON YAKWOL', 6496, 467, 34),
	(79188, 'DENOM', 6496, 467, 34),
	(79189, 'OK TELAB', 6496, 467, 34),
	(79190, 'OKBON', 6496, 467, 34),
	(79191, 'PEPERA', 6496, 467, 34),
	(79192, 'WOK BAKON', 6496, 467, 34),
	(79193, 'YUN MUKU', 6496, 467, 34),
	(79194, 'MODUSIT', 6497, 467, 34),
	(79195, 'OKATEM', 6497, 467, 34),
	(79196, 'PARIM', 6497, 467, 34),
	(79197, 'SERAMKATOP', 6497, 467, 34),
	(79198, 'SIMINBUK', 6497, 467, 34),
	(79199, 'WANBAKON', 6497, 467, 34),
	(79200, 'YAKMOR', 6497, 467, 34),
	(79201, 'YAPIMAKOT', 6497, 467, 34),
	(79202, 'BETEN DUA', 6498, 467, 34),
	(79203, 'BITIPDING', 6498, 467, 34),
	(79204, 'IMSIN', 6498, 467, 34),
	(79205, 'MARANG TIKING', 6498, 467, 34),
	(79206, 'OMOR', 6498, 467, 34),
	(79207, 'ONKOR', 6498, 467, 34),
	(79208, 'TARUP', 6498, 467, 34),
	(79209, 'BAUTME', 6499, 467, 34),
	(79210, 'MAIGAME', 6499, 467, 34),
	(79211, 'MURME', 6499, 467, 34),
	(79212, 'SINAMI / SINANI (YUABAN)', 6499, 467, 34),
	(79213, 'TEIRAPLU', 6499, 467, 34),
	(79214, 'TERAPDEI', 6499, 467, 34),
	(79215, 'TERIAME', 6499, 467, 34),
	(79216, 'YITARGET', 6499, 467, 34),
	(79217, 'YUABAN DUA', 6499, 467, 34),
	(79218, 'YUABAN SATU', 6499, 467, 34),
	(79219, 'DALUBAN', 6500, 467, 34),
	(79220, 'LIMREPASIKNE', 6500, 467, 34),
	(79221, 'MEKDAMGON', 6500, 467, 34),
	(79222, 'MERPASIKNE', 6500, 467, 34),
	(79223, 'MERYANG', 6500, 467, 34),
	(79224, 'NOMTEREN', 6500, 467, 34),
	(79225, 'TARAMLU', 6500, 467, 34),
	(79226, 'WEIME', 6500, 467, 34),
	(79227, 'YOULBAN (YOLBAN)', 6500, 467, 34),
	(79228, 'AGADUGUME', 6501, 468, 34),
	(79229, 'JIWOT', 6501, 468, 34),
	(79230, 'TUPUT', 6501, 468, 34),
	(79231, 'BABE', 6502, 468, 34),
	(79232, 'DAMBET (DANGBET/KALEMOGOM)', 6502, 468, 34),
	(79233, 'JAMBUL', 6502, 468, 34),
	(79234, 'KELANDIRUMA', 6502, 468, 34),
	(79235, 'MILAWAK', 6502, 468, 34),
	(79236, 'NUNGAI', 6502, 468, 34),
	(79237, 'OGAMKI', 6502, 468, 34),
	(79238, 'PUBET / GIMORIT', 6502, 468, 34),
	(79239, 'PULUK', 6502, 468, 34),
	(79240, 'TINGGILBET', 6502, 468, 34),
	(79241, 'WAMKI (DENGKIBUMA)', 6502, 468, 34),
	(79242, 'YULOKOMA', 6502, 468, 34),
	(79243, 'DOUFU/DOVO', 6503, 468, 34),
	(79244, 'FAISAU', 6503, 468, 34),
	(79245, 'IRATOI', 6503, 468, 34),
	(79246, 'KORDEI (KORDESI)', 6503, 468, 34),
	(79247, 'TAYAI (TAYA)', 6503, 468, 34),
	(79248, 'AGIYOME', 6504, 468, 34),
	(79249, 'AMUNGKALPIA (AMUNGKALIPIA)', 6504, 468, 34),
	(79250, 'BELA', 6504, 468, 34),
	(79251, 'ERELMEKAWIA', 6504, 468, 34),
	(79252, 'GOME', 6504, 468, 34),
	(79253, 'MISIMAGA', 6504, 468, 34),
	(79254, 'MUNDIDOK', 6504, 468, 34),
	(79255, 'TOANGGI', 6504, 468, 34),
	(79256, 'UPAGA', 6504, 468, 34),
	(79257, 'WAKO', 6504, 468, 34),
	(79258, 'YAIKIMAKI', 6504, 468, 34),
	(79259, 'YENGGRENOK (YANGGERENOK)', 6504, 468, 34),
	(79260, 'BOLOGOBAK (POLOGOBAK/POLOGOBALE)', 6505, 468, 34),
	(79261, 'BUBET', 6505, 468, 34),
	(79262, 'EROMAGA (EROMOGA)', 6505, 468, 34),
	(79263, 'GILINI', 6505, 468, 34),
	(79264, 'KAGO', 6505, 468, 34),
	(79265, 'KIMAK (KIMALE)', 6505, 468, 34),
	(79266, 'KUNGA (KUIGA)', 6505, 468, 34),
	(79267, 'MAKI', 6505, 468, 34),
	(79268, 'MAYUBERI', 6505, 468, 34),
	(79269, 'MINDIBA (MIANDIWA/MUDIBA)', 6505, 468, 34),
	(79270, 'OGNANIM (OKNANIM)', 6505, 468, 34),
	(79271, 'PALUGA', 6505, 468, 34),
	(79272, 'PINAPA', 6505, 468, 34),
	(79273, 'PINGGIL (PIGGIL)', 6505, 468, 34),
	(79274, 'WULONI', 6505, 468, 34),
	(79275, 'AGUIT (AGUWIT)', 6506, 468, 34),
	(79276, 'BACINI (BAKCINI)', 6506, 468, 34),
	(79277, 'BINA', 6506, 468, 34),
	(79278, 'GAGAMA', 6506, 468, 34),
	(79279, 'GOLU', 6506, 468, 34),
	(79280, 'KEMBRU (YIMUK)', 6506, 468, 34),
	(79281, 'MOLU', 6506, 468, 34),
	(79282, 'POGOMA (BOKOMA)', 6506, 468, 34),
	(79283, 'WAKME', 6506, 468, 34),
	(79284, 'WIHA', 6506, 468, 34),
	(79285, 'AKENGGEN', 6507, 468, 34),
	(79286, 'AMULAME (AMULUME)', 6507, 468, 34),
	(79287, 'GIGOBAK (GOGOBAK)', 6507, 468, 34),
	(79288, 'JIGIUNGGI (JIGINGGI)', 6507, 468, 34),
	(79289, 'KALIBUK (KALEBUT/KAWIBUT)', 6507, 468, 34),
	(79290, 'KELEMAME', 6507, 468, 34),
	(79291, 'KILUGAME', 6507, 468, 34),
	(79292, 'KUNIKOMO (KUMIKOMO)', 6507, 468, 34),
	(79293, 'NIKOLEME (NIGOLOME/NIKOAME)', 6507, 468, 34),
	(79294, 'PAMEBUT', 6507, 468, 34),
	(79295, 'WENI (ISWELI/WELI)', 6507, 468, 34),
	(79296, 'YAURIA (WAMBRU)', 6507, 468, 34),
	(79297, 'AILPALIN (ALPALIN)', 6508, 468, 34),
	(79298, 'DAUNG', 6508, 468, 34),
	(79299, 'ERONG BERONG/BERANG', 6508, 468, 34),
	(79300, 'HIKINAT (IKINAT)', 6508, 468, 34),
	(79301, 'JINDAK', 6508, 468, 34),
	(79302, 'MIRILAUKIN', 6508, 468, 34),
	(79303, 'PILOKOMA', 6508, 468, 34),
	(79304, 'PUNGKI', 6508, 468, 34),
	(79305, 'ULIPIA', 6508, 468, 34),
	(79306, 'WANGBE', 6508, 468, 34),
	(79307, 'WONALBE', 6508, 468, 34),
	(79308, 'AMBO', 6509, 469, 34),
	(79309, 'AMURINGGIK', 6509, 469, 34),
	(79310, 'ATOLI', 6509, 469, 34),
	(79311, 'BAKUSI', 6509, 469, 34),
	(79312, 'BIRICARE', 6509, 469, 34),
	(79313, 'DAGAI', 6509, 469, 34),
	(79314, 'DAGAI 2', 6509, 469, 34),
	(79315, 'DEBITE', 6509, 469, 34),
	(79316, 'DEIDE', 6509, 469, 34),
	(79317, 'DOREY', 6509, 469, 34),
	(79318, 'EVO', 6509, 469, 34),
	(79319, 'FARRIDE', 6509, 469, 34),
	(79320, 'FAWI', 6509, 469, 34),
	(79321, 'FII', 6509, 469, 34),
	(79322, 'GUERI', 6509, 469, 34),
	(79323, 'KAHO', 6509, 469, 34),
	(79324, 'KIYAGE', 6509, 469, 34),
	(79325, 'MANDARA LANI', 6509, 469, 34),
	(79326, 'MBOMBAN', 6509, 469, 34),
	(79327, 'SOI', 6509, 469, 34),
	(79328, 'TENGGABANGGWI', 6509, 469, 34),
	(79329, 'TURUMO', 6509, 469, 34),
	(79330, 'WANGGIBA', 6509, 469, 34),
	(79331, 'YEIHNERI', 6509, 469, 34),
	(79332, 'YEREI', 6509, 469, 34),
	(79333, 'ABERIAMBUT', 6510, 469, 34),
	(79334, 'AGOPAGA', 6510, 469, 34),
	(79335, 'AMBIT-MBIT', 6510, 469, 34),
	(79336, 'ANDIRON', 6510, 469, 34),
	(79337, 'ANGGUTARI', 6510, 469, 34),
	(79338, 'ARINGGAK', 6510, 469, 34),
	(79339, 'AULUKME', 6510, 469, 34),
	(79340, 'BELANTARA', 6510, 469, 34),
	(79341, 'BUNUME', 6510, 469, 34),
	(79342, 'DIGULOME', 6510, 469, 34),
	(79343, 'DOLINGGAME', 6510, 469, 34),
	(79344, 'GALONEGAME', 6510, 469, 34),
	(79345, 'GIRMOR', 6510, 469, 34),
	(79346, 'GUBULOME', 6510, 469, 34),
	(79347, 'GUMAWI', 6510, 469, 34),
	(79348, 'JEMBENERI', 6510, 469, 34),
	(79349, 'JIBONOK', 6510, 469, 34),
	(79350, 'JIGELO', 6510, 469, 34),
	(79351, 'JIGULUK', 6510, 469, 34),
	(79352, 'JIMBANIME', 6510, 469, 34),
	(79353, 'JURIA I', 6510, 469, 34),
	(79354, 'KALENGGA', 6510, 469, 34),
	(79355, 'KIRIGIMADUK', 6510, 469, 34),
	(79356, 'KOBARAK', 6510, 469, 34),
	(79357, 'KURIKPULOK', 6510, 469, 34),
	(79358, 'LAMBO', 6510, 469, 34),
	(79359, 'LAWOREGE', 6510, 469, 34),
	(79360, 'MAKA', 6510, 469, 34),
	(79361, 'MOBIGI', 6510, 469, 34),
	(79362, 'MOULO', 6510, 469, 34),
	(79363, 'NGGINIGUM', 6510, 469, 34),
	(79364, 'NINOM', 6510, 469, 34),
	(79365, 'NUMBUKENGGAWI', 6510, 469, 34),
	(79366, 'PEREYA', 6510, 469, 34),
	(79367, 'PINDEBAGA', 6510, 469, 34),
	(79368, 'PIYAPIGI', 6510, 469, 34),
	(79369, 'PULAU TIMUR', 6510, 469, 34),
	(79370, 'PULOGENGGA', 6510, 469, 34),
	(79371, 'TENOMANGGEN', 6510, 469, 34),
	(79372, 'TIRIGOI', 6510, 469, 34),
	(79373, 'WUGIWAGI', 6510, 469, 34),
	(79374, 'WURAK', 6510, 469, 34),
	(79375, 'YAGALUK', 6510, 469, 34),
	(79376, 'YAMONERI', 6510, 469, 34),
	(79377, 'AKWIBAGA', 6511, 469, 34),
	(79378, 'ANEBALUI', 6511, 469, 34),
	(79379, 'APELOME', 6511, 469, 34),
	(79380, 'GIBAGA', 6511, 469, 34),
	(79381, 'GILOME', 6511, 469, 34),
	(79382, 'GINILUME', 6511, 469, 34),
	(79383, 'GUBUME', 6511, 469, 34),
	(79384, 'GUBUNERI', 6511, 469, 34),
	(79385, 'GUNA', 6511, 469, 34),
	(79386, 'JIGONIKME', 6511, 469, 34),
	(79387, 'JINGGI', 6511, 469, 34),
	(79388, 'KENENDAGA', 6511, 469, 34),
	(79389, 'KERIN', 6511, 469, 34),
	(79390, 'KOMANGGI', 6511, 469, 34),
	(79391, 'LUGUBAGO', 6511, 469, 34),
	(79392, 'MABUK TIMUR', 6511, 469, 34),
	(79393, 'MANDURA', 6511, 469, 34),
	(79394, 'MELELA', 6511, 469, 34),
	(79395, 'MEPAR', 6511, 469, 34),
	(79396, 'MEYONGGA', 6511, 469, 34),
	(79397, 'MILINERI', 6511, 469, 34),
	(79398, 'MODU', 6511, 469, 34),
	(79399, 'MURUWI', 6511, 469, 34),
	(79400, 'NDUGWIR', 6511, 469, 34),
	(79401, 'NIOGA', 6511, 469, 34),
	(79402, 'NIYOLUK ABILI', 6511, 469, 34),
	(79403, 'NOBA NOBA', 6511, 469, 34),
	(79404, 'NOGI', 6511, 469, 34),
	(79405, 'NUMBOK', 6511, 469, 34),
	(79406, 'NUME', 6511, 469, 34),
	(79407, 'ONENDU', 6511, 469, 34),
	(79408, 'OUM', 6511, 469, 34),
	(79409, 'PAGOLOME', 6511, 469, 34),
	(79410, 'PALILOME', 6511, 469, 34),
	(79411, 'PAPAK', 6511, 469, 34),
	(79412, 'PUGUNOM', 6511, 469, 34),
	(79413, 'PUKI PAKI', 6511, 469, 34),
	(79414, 'TINGGIREGE', 6511, 469, 34),
	(79415, 'TOMBOK', 6511, 469, 34),
	(79416, 'WALON', 6511, 469, 34),
	(79417, 'WAMBAGALO', 6511, 469, 34),
	(79418, 'WANDURI', 6511, 469, 34),
	(79419, 'WONGI', 6511, 469, 34),
	(79420, 'WORALUK', 6511, 469, 34),
	(79421, 'WUMBIRI', 6511, 469, 34),
	(79422, 'WUTIME', 6511, 469, 34),
	(79423, 'YAWOR', 6511, 469, 34),
	(79424, 'YONGGI', 6511, 469, 34),
	(79425, 'BALINGGUP', 6512, 469, 34),
	(79426, 'BELABAGA', 6512, 469, 34),
	(79427, 'BIAK', 6512, 469, 34),
	(79428, 'DOLOGOBAK', 6512, 469, 34),
	(79429, 'GILIBE', 6512, 469, 34),
	(79430, 'GILILOME', 6512, 469, 34),
	(79431, 'GOLAPAGA', 6512, 469, 34),
	(79432, 'GUMBRU', 6512, 469, 34),
	(79433, 'KILILUMO', 6512, 469, 34),
	(79434, 'LUMO', 6512, 469, 34),
	(79435, 'MALOINGGEN', 6512, 469, 34),
	(79436, 'MEWOLUK', 6512, 469, 34),
	(79437, 'MEWUT', 6512, 469, 34),
	(79438, 'NGGININIK', 6512, 469, 34),
	(79439, 'NINGGINERI', 6512, 469, 34),
	(79440, 'OGOLUK', 6512, 469, 34),
	(79441, 'TIGIR', 6512, 469, 34),
	(79442, 'TIOLOME', 6512, 469, 34),
	(79443, 'UTIKME', 6512, 469, 34),
	(79444, 'WALIBA', 6512, 469, 34),
	(79445, 'WANUME', 6512, 469, 34),
	(79446, 'WURABAK', 6512, 469, 34),
	(79447, 'WURAMBURU', 6512, 469, 34),
	(79448, 'AMULOME', 6513, 469, 34),
	(79449, 'ANGGUTARE', 6513, 469, 34),
	(79450, 'BERELEMA', 6513, 469, 34),
	(79451, 'BIRAK AMBUT', 6513, 469, 34),
	(79452, 'DANGENPAGA', 6513, 469, 34),
	(79453, 'DOLIGOBAK', 6513, 469, 34),
	(79454, 'GINIGOM', 6513, 469, 34),
	(79455, 'ILAMBURAWI', 6513, 469, 34),
	(79456, 'JIGINIKIME', 6513, 469, 34),
	(79457, 'KARUBATE', 6513, 469, 34),
	(79458, 'KILUNGGWI', 6513, 469, 34),
	(79459, 'KULIRIK', 6513, 469, 34),
	(79460, 'KULOK ENGGAME', 6513, 469, 34),
	(79461, 'LULAME', 6513, 469, 34),
	(79462, 'LUNGGINERI', 6513, 469, 34),
	(79463, 'MONDU', 6513, 469, 34),
	(79464, 'MOULO', 6513, 469, 34),
	(79465, 'MULIAGAMBUT', 6513, 469, 34),
	(79466, 'MULIAMBUT', 6513, 469, 34),
	(79467, 'OMABUT', 6513, 469, 34),
	(79468, 'PAGALEME', 6513, 469, 34),
	(79469, 'PEPERA', 6513, 469, 34),
	(79470, 'PRULEME', 6513, 469, 34),
	(79471, 'PUNCAK SENYUM', 6513, 469, 34),
	(79472, 'TALILOME', 6513, 469, 34),
	(79473, 'TENOLOK', 6513, 469, 34),
	(79474, 'TOWEMARIB', 6513, 469, 34),
	(79475, 'TOWOGI', 6513, 469, 34),
	(79476, 'TRIKORA', 6513, 469, 34),
	(79477, 'URGELE', 6513, 469, 34),
	(79478, 'USIR', 6513, 469, 34),
	(79479, 'WANDURI', 6513, 469, 34),
	(79480, 'WONDENGGOBAK', 6513, 469, 34),
	(79481, 'WONOME', 6513, 469, 34),
	(79482, 'WUYUKWI', 6513, 469, 34),
	(79483, 'WUYUNERI', 6513, 469, 34),
	(79484, 'YALINGGUA', 6513, 469, 34),
	(79485, 'YAMBI', 6513, 469, 34),
	(79486, 'YAMBIDUGUN', 6513, 469, 34),
	(79487, 'YOBOLUK', 6513, 469, 34),
	(79488, 'YOGA', 6513, 469, 34),
	(79489, 'AGAPE', 6514, 469, 34),
	(79490, 'AULAKMU', 6514, 469, 34),
	(79491, 'AULAME', 6514, 469, 34),
	(79492, 'BEREM', 6514, 469, 34),
	(79493, 'BERIME', 6514, 469, 34),
	(79494, 'BIGIRAGE', 6514, 469, 34),
	(79495, 'BIME', 6514, 469, 34),
	(79496, 'BINGGELAKME', 6514, 469, 34),
	(79497, 'BONALUK', 6514, 469, 34),
	(79498, 'CILOME', 6514, 469, 34),
	(79499, 'DEGI', 6514, 469, 34),
	(79500, 'DIRALUK', 6514, 469, 34),
	(79501, 'GIBUME', 6514, 469, 34),
	(79502, 'GIGUME', 6514, 469, 34),
	(79503, 'GIMANGGEN', 6514, 469, 34),
	(79504, 'GINILOME', 6514, 469, 34),
	(79505, 'GINIPAGO', 6514, 469, 34),
	(79506, 'GREJA LAMA', 6514, 469, 34),
	(79507, 'GUBUPUR', 6514, 469, 34),
	(79508, 'GWENGGU', 6514, 469, 34),
	(79509, 'INIKIMALUK', 6514, 469, 34),
	(79510, 'JIBINGGAME', 6514, 469, 34),
	(79511, 'JIGOBAK', 6514, 469, 34),
	(79512, 'JINGGWI', 6514, 469, 34),
	(79513, 'JIRAMOK', 6514, 469, 34),
	(79514, 'JIWONE', 6514, 469, 34),
	(79515, 'JUGUMBLAWI', 6514, 469, 34),
	(79516, 'KALOME', 6514, 469, 34),
	(79517, 'KAYOGEBUR', 6514, 469, 34),
	(79518, 'KIBURU', 6514, 469, 34),
	(79519, 'LERAWERA', 6514, 469, 34),
	(79520, 'LONGGAWI', 6514, 469, 34),
	(79521, 'LUMBUK', 6514, 469, 34),
	(79522, 'LURA BARA', 6514, 469, 34),
	(79523, 'MELEKOM', 6514, 469, 34),
	(79524, 'MONDOGONERI', 6514, 469, 34),
	(79525, 'MONIA', 6514, 469, 34),
	(79526, 'NABURAGE', 6514, 469, 34),
	(79527, 'NAKONGWE', 6514, 469, 34),
	(79528, 'NENEGAME', 6514, 469, 34),
	(79529, 'NGGEGEN', 6514, 469, 34),
	(79530, 'NUWI', 6514, 469, 34),
	(79531, 'ORILUK', 6514, 469, 34),
	(79532, 'PAGARUGOM', 6514, 469, 34),
	(79533, 'PALUMAGI', 6514, 469, 34),
	(79534, 'PAPUA', 6514, 469, 34),
	(79535, 'PARALO', 6514, 469, 34),
	(79536, 'PELALO', 6514, 469, 34),
	(79537, 'PERNALUK', 6514, 469, 34),
	(79538, 'PILIA', 6514, 469, 34),
	(79539, 'PILIBUR', 6514, 469, 34),
	(79540, 'PONGGONAME', 6514, 469, 34),
	(79541, 'PURUGI', 6514, 469, 34),
	(79542, 'TENOLOK', 6514, 469, 34),
	(79543, 'TINGGINAMBUT', 6514, 469, 34),
	(79544, 'TINGGINERI', 6514, 469, 34),
	(79545, 'TINGGINIME', 6514, 469, 34),
	(79546, 'TOMBO', 6514, 469, 34),
	(79547, 'TORAGE', 6514, 469, 34),
	(79548, 'TOWOLUK', 6514, 469, 34),
	(79549, 'TUROTARO', 6514, 469, 34),
	(79550, 'URAGI', 6514, 469, 34),
	(79551, 'WIYAGE', 6514, 469, 34),
	(79552, 'WONWI', 6514, 469, 34),
	(79553, 'WUNDINI', 6514, 469, 34),
	(79554, 'WURABUME', 6514, 469, 34),
	(79555, 'WURUNGKIME', 6514, 469, 34),
	(79556, 'YALIBATE', 6514, 469, 34),
	(79557, 'YAMENGGA', 6514, 469, 34),
	(79558, 'YANENGGAWI', 6514, 469, 34),
	(79559, 'YANIRUK', 6514, 469, 34),
	(79560, 'YARMUKUM', 6514, 469, 34),
	(79561, 'YOGORINI', 6514, 469, 34),
	(79562, 'YONGGUN', 6514, 469, 34),
	(79563, 'YUNGGWI', 6514, 469, 34),
	(79564, 'AMBOK', 6515, 469, 34),
	(79565, 'ASUA', 6515, 469, 34),
	(79566, 'AWIYAM', 6515, 469, 34),
	(79567, 'DEGEI', 6515, 469, 34),
	(79568, 'DIGI', 6515, 469, 34),
	(79569, 'GUBUGANI', 6515, 469, 34),
	(79570, 'GUNUNG TAYOK', 6515, 469, 34),
	(79571, 'KELANDU', 6515, 469, 34),
	(79572, 'KIKUP', 6515, 469, 34),
	(79573, 'NALU', 6515, 469, 34),
	(79574, 'NAMBU', 6515, 469, 34),
	(79575, 'SIGOU', 6515, 469, 34),
	(79576, 'TUBUGIME', 6515, 469, 34),
	(79577, 'TUMBIWOLU', 6515, 469, 34),
	(79578, 'WARIRU', 6515, 469, 34),
	(79579, 'ALIGAPAGA', 6516, 469, 34),
	(79580, 'BINIME', 6516, 469, 34),
	(79581, 'DOKOME', 6516, 469, 34),
	(79582, 'DONDO', 6516, 469, 34),
	(79583, 'GIBAGA', 6516, 469, 34),
	(79584, 'GOYAGE', 6516, 469, 34),
	(79585, 'IMULINERI', 6516, 469, 34),
	(79586, 'JIGUNIKIME', 6516, 469, 34),
	(79587, 'KABILIMBUT', 6516, 469, 34),
	(79588, 'KIMIBUT', 6516, 469, 34),
	(79589, 'KWATINERI', 6516, 469, 34),
	(79590, 'MOLOBAK', 6516, 469, 34),
	(79591, 'NAMI', 6516, 469, 34),
	(79592, 'NIRUWI', 6516, 469, 34),
	(79593, 'NOWOME', 6516, 469, 34),
	(79594, 'NOWONERI', 6516, 469, 34),
	(79595, 'PULAU YAMO', 6516, 469, 34),
	(79596, 'PURBALO', 6516, 469, 34),
	(79597, 'SANOBA', 6516, 469, 34),
	(79598, 'TEMU', 6516, 469, 34),
	(79599, 'TIOGEME', 6516, 469, 34),
	(79600, 'TUKWI', 6516, 469, 34),
	(79601, 'WIYAKTINI', 6516, 469, 34),
	(79602, 'WOLAME', 6516, 469, 34),
	(79603, 'WUNAGELO', 6516, 469, 34),
	(79604, 'WUNDU', 6516, 469, 34),
	(79605, 'WURAGE', 6516, 469, 34),
	(79606, 'YALIKAMBI', 6516, 469, 34),
	(79607, 'YALIMAMBU', 6516, 469, 34),
	(79608, 'YAMO', 6516, 469, 34),
	(79609, 'YOGOLAWI', 6516, 469, 34),
	(79610, 'AIRORAN', 6517, 470, 34),
	(79611, 'AURIMI', 6517, 470, 34),
	(79612, 'BINA', 6517, 470, 34),
	(79613, 'KWAWITANIA', 6517, 470, 34),
	(79614, 'MANIWA', 6517, 470, 34),
	(79615, 'MURARA', 6517, 470, 34),
	(79616, 'SASAWAPECE', 6517, 470, 34),
	(79617, 'SURIMANIA (SYURIMANIA)', 6517, 470, 34),
	(79618, 'TAMAYA', 6517, 470, 34),
	(79619, 'WAMARIRI', 6517, 470, 34),
	(79620, 'ANUS', 6518, 470, 34),
	(79621, 'ARMOPA', 6518, 470, 34),
	(79622, 'BEBON JAYA', 6518, 470, 34),
	(79623, 'KIREN', 6518, 470, 34),
	(79624, 'KRIM PODENA', 6518, 470, 34),
	(79625, 'MAWESWARES (MAWESRES)', 6518, 470, 34),
	(79626, 'RIMSER SARI', 6518, 470, 34),
	(79627, 'ROTEA', 6518, 470, 34),
	(79628, 'TARONTA SRUM', 6518, 470, 34),
	(79629, 'TETOM', 6518, 470, 34),
	(79630, 'GWIN JAYA', 6519, 470, 34),
	(79631, 'KAPTIAU', 6519, 470, 34),
	(79632, 'MAWES MUKTI', 6519, 470, 34),
	(79633, 'MAWESDAY', 6519, 470, 34),
	(79634, 'TAMAS SARI', 6519, 470, 34),
	(79635, 'ARBAIS', 6520, 470, 34),
	(79636, 'ARUSWAR', 6520, 470, 34),
	(79637, 'BURGENA', 6520, 470, 34),
	(79638, 'KAMENAWARI', 6520, 470, 34),
	(79639, 'KAPESO', 6520, 470, 34),
	(79640, 'KARFASIA', 6520, 470, 34),
	(79641, 'MARTEWAR', 6520, 470, 34),
	(79642, 'MASEP', 6520, 470, 34),
	(79643, 'NISRO', 6520, 470, 34),
	(79644, 'NIWERAWAR', 6520, 470, 34),
	(79645, 'SAMORKENA', 6520, 470, 34),
	(79646, 'SIANTOA', 6520, 470, 34),
	(79647, 'SUBU', 6520, 470, 34),
	(79648, 'WAIM', 6520, 470, 34),
	(79649, 'WARI', 6520, 470, 34),
	(79650, 'WEBRO', 6520, 470, 34),
	(79651, 'ANSUDU (ANSUBI)', 6521, 470, 34),
	(79652, 'ANSUDU II (SREM)', 6521, 470, 34),
	(79653, 'BETAF', 6521, 470, 34),
	(79654, 'BETAF II (TAMNIR)', 6521, 470, 34),
	(79655, 'KOMRA', 6521, 470, 34),
	(79656, 'SUNUM (YAMNA)', 6521, 470, 34),
	(79657, 'YAMBEN (JAMBER/BENERAF)', 6521, 470, 34),
	(79658, 'AMPERA', 6522, 470, 34),
	(79659, 'ARARE', 6522, 470, 34),
	(79660, 'DABE', 6522, 470, 34),
	(79661, 'DABE II', 6522, 470, 34),
	(79662, 'FITO', 6522, 470, 34),
	(79663, 'KEDER II', 6522, 470, 34),
	(79664, 'NENGKE', 6522, 470, 34),
	(79665, 'NENGKE II (NENGKE BARU)', 6522, 470, 34),
	(79666, 'TABRAWAR', 6522, 470, 34),
	(79667, 'TIMRON', 6522, 470, 34),
	(79668, 'VINYABOR', 6522, 470, 34),
	(79669, 'ARMO', 6523, 470, 34),
	(79670, 'BAGAISERWAR', 6523, 470, 34),
	(79671, 'BAGAISEWAR II', 6523, 470, 34),
	(79672, 'LEMBAH NEIDAM', 6523, 470, 34),
	(79673, 'LIKI', 6523, 470, 34),
	(79674, 'MARARENA', 6523, 470, 34),
	(79675, 'PULAU ARMO', 6523, 470, 34),
	(79676, 'SARMI', 6523, 470, 34),
	(79677, 'SARMO', 6523, 470, 34),
	(79678, 'SAWAR', 6523, 470, 34),
	(79679, 'TEFAREWAR', 6523, 470, 34),
	(79680, 'AMSIRA', 6524, 470, 34),
	(79681, 'ANGKASA DUA', 6524, 470, 34),
	(79682, 'KASUKWE', 6524, 470, 34),
	(79683, 'MANUKANIA', 6524, 470, 34),
	(79684, 'SIARATESA', 6524, 470, 34),
	(79685, 'WAPO', 6524, 470, 34),
	(79686, 'BAGAISEWAR DUA / II', 6525, 470, 34),
	(79687, 'BINYER', 6525, 470, 34),
	(79688, 'BIYER', 6525, 470, 34),
	(79689, 'EBRAM', 6525, 470, 34),
	(79690, 'EBRAM', 6525, 470, 34),
	(79691, 'HOLMAFEN (HOLMAFEM)', 6525, 470, 34),
	(79692, 'SEWAN', 6525, 470, 34),
	(79693, 'TANJUNG BATU', 6525, 470, 34),
	(79694, 'WASKEY', 6525, 470, 34),
	(79695, 'WASKEY / TANJUNG BATU', 6525, 470, 34),
	(79696, 'BORA BORA', 6526, 470, 34),
	(79697, 'DENANDER', 6526, 470, 34),
	(79698, 'DENANDER DUA/II', 6526, 470, 34),
	(79699, 'KONDERJAN', 6526, 470, 34),
	(79700, 'OMTE (KWEPKE)', 6526, 470, 34),
	(79701, 'SAFRON TANE DUA/II', 6526, 470, 34),
	(79702, 'SAKRON TANE', 6526, 470, 34),
	(79703, 'SAMANENTE', 6526, 470, 34),
	(79704, 'TOGONFO (TEMINABOR)', 6526, 470, 34),
	(79705, 'WAAF', 6526, 470, 34),
	(79706, 'WAAF DUA/II', 6526, 470, 34),
	(79707, 'ARURI', 6527, 471, 34),
	(79708, 'IMBRISBARI (IMBIRSBARI)', 6527, 471, 34),
	(79709, 'INEKI', 6527, 471, 34),
	(79710, 'INSUMBREI', 6527, 471, 34),
	(79711, 'MANGGONSWAN (MANGOSWAN)', 6527, 471, 34),
	(79712, 'MBRUWANDI (BRUWANDI)', 6527, 471, 34),
	(79713, 'RAYORI', 6527, 471, 34),
	(79714, 'WONGGEINA (WONGKEINA)', 6527, 471, 34),
	(79715, 'YAMNAISU', 6527, 471, 34),
	(79716, 'AMYAS', 6528, 471, 34),
	(79717, 'KORYAKAM (KOIRYAKAM)', 6528, 471, 34),
	(79718, 'MAPIA', 6528, 471, 34),
	(79719, 'MASYAI', 6528, 471, 34),
	(79720, 'NAPISANDI', 6528, 471, 34),
	(79721, 'WARYEI (WARIYEI)', 6528, 471, 34),
	(79722, 'WAYORI (MAYORI)', 6528, 471, 34),
	(79723, 'AWAKI', 6529, 471, 34),
	(79724, 'BINIKI', 6529, 471, 34),
	(79725, 'DIDIABOLO', 6529, 471, 34),
	(79726, 'FANINDI', 6529, 471, 34),
	(79727, 'MARYAIDORI', 6529, 471, 34),
	(79728, 'ODORI', 6529, 471, 34),
	(79729, 'WABERFONDI (WARBEPONDI)', 6529, 471, 34),
	(79730, 'DOUWBO (DOUBO)', 6530, 471, 34),
	(79731, 'DUBER', 6530, 471, 34),
	(79732, 'MARSRAM', 6530, 471, 34),
	(79733, 'SAUYAS', 6530, 471, 34),
	(79734, 'SORENDIDORI (SORENDIWERI)', 6530, 471, 34),
	(79735, 'SYURDORI', 6530, 471, 34),
	(79736, 'WAFOR', 6530, 471, 34),
	(79737, 'WARYESI', 6530, 471, 34),
	(79738, 'WOMBONDA', 6530, 471, 34),
	(79739, 'YAWERMA', 6530, 471, 34),
	(79740, 'FANJUR', 6531, 471, 34),
	(79741, 'KOBARI JAYA', 6531, 471, 34),
	(79742, 'PUWERI', 6531, 471, 34),
	(79743, 'WARBOR', 6531, 471, 34),
	(79744, 'WARSA', 6531, 471, 34),
	(79745, 'KUBUR', 6532, 472, 34),
	(79746, 'LENGGUP', 6532, 472, 34),
	(79747, 'LIWESE', 6532, 472, 34),
	(79748, 'ONGGOKME', 6532, 472, 34),
	(79749, 'TABO WANIMBO', 6532, 472, 34),
	(79750, 'TINGER', 6532, 472, 34),
	(79751, 'WENDURI', 6532, 472, 34),
	(79752, 'WEU', 6532, 472, 34),
	(79753, 'ANAWI', 6533, 472, 34),
	(79754, 'ARIDUNDA', 6533, 472, 34),
	(79755, 'BIELEME', 6533, 472, 34),
	(79756, 'GINERI', 6533, 472, 34),
	(79757, 'IMURIK', 6533, 472, 34),
	(79758, 'KOTORI', 6533, 472, 34),
	(79759, 'LINGGIRA', 6533, 472, 34),
	(79760, 'LOGUK', 6533, 472, 34),
	(79761, 'YALIPURA', 6533, 472, 34),
	(79762, 'YALOKOBAK', 6533, 472, 34),
	(79763, 'AGIN', 6534, 472, 34),
	(79764, 'KOGAGI', 6534, 472, 34),
	(79765, 'KOLANGGUN', 6534, 472, 34),
	(79766, 'POSMAN', 6534, 472, 34),
	(79767, 'TIYONGGI', 6534, 472, 34),
	(79768, 'WAMIGI', 6534, 472, 34),
	(79769, 'WENGGUN', 6534, 472, 34),
	(79770, 'WULUK', 6534, 472, 34),
	(79771, 'YEBENA', 6534, 472, 34),
	(79772, 'YELLY', 6534, 472, 34),
	(79773, 'ABENA', 6535, 472, 34),
	(79774, 'ARELAM', 6535, 472, 34),
	(79775, 'BILUBAGA', 6535, 472, 34),
	(79776, 'BITILLABUR', 6535, 472, 34),
	(79777, 'DUMA', 6535, 472, 34),
	(79778, 'GABUNGGOBAK', 6535, 472, 34),
	(79779, 'GELALO', 6535, 472, 34),
	(79780, 'NOGOBUMBU', 6535, 472, 34),
	(79781, 'WANGGULAM', 6535, 472, 34),
	(79782, 'WANIA', 6535, 472, 34),
	(79783, 'WINDIK', 6535, 472, 34),
	(79784, 'WULURIK', 6535, 472, 34),
	(79785, 'YIBALO', 6535, 472, 34),
	(79786, 'YINAMA', 6535, 472, 34),
	(79787, 'BIUK', 6536, 472, 34),
	(79788, 'GALUBUP', 6536, 472, 34),
	(79789, 'GUBURINI', 6536, 472, 34),
	(79790, 'MBINIME/JINULIRA', 6536, 472, 34),
	(79791, 'PURUGI', 6536, 472, 34),
	(79792, 'TOMAGI/GUBAGI', 6536, 472, 34),
	(79793, 'TOMAGIPURA', 6536, 472, 34),
	(79794, 'WONABU', 6536, 472, 34),
	(79795, 'YILUK/KONDENGGUN', 6536, 472, 34),
	(79796, 'YIYOGOBAK/KIBUR', 6536, 472, 34),
	(79797, 'YUGU MABUR', 6536, 472, 34),
	(79798, 'ALIDUDA', 6537, 472, 34),
	(79799, 'ANDOMAK', 6537, 472, 34),
	(79800, 'BOGONUK', 6537, 472, 34),
	(79801, 'EWAN', 6537, 472, 34),
	(79802, 'LAURA', 6537, 472, 34),
	(79803, 'PABA', 6537, 472, 34),
	(79804, 'TALINAMBER', 6537, 472, 34),
	(79805, 'WALELO', 6537, 472, 34),
	(79806, 'WISMAN', 6537, 472, 34),
	(79807, 'WUMELAK', 6537, 472, 34),
	(79808, 'APIAM', 6538, 472, 34),
	(79809, 'BOKONDINI', 6538, 472, 34),
	(79810, 'DUNDUMA', 6538, 472, 34),
	(79811, 'GALALA', 6538, 472, 34),
	(79812, 'JAWALANE', 6538, 472, 34),
	(79813, 'KOLOGUME', 6538, 472, 34),
	(79814, 'LAMBOGO', 6538, 472, 34),
	(79815, 'MAIRINI', 6538, 472, 34),
	(79816, 'MINGGANGGO', 6538, 472, 34),
	(79817, 'TENGGAGAMA', 6538, 472, 34),
	(79818, 'UMAGA', 6538, 472, 34),
	(79819, 'ABIMBAK', 6539, 472, 34),
	(79820, 'BOKONERI', 6539, 472, 34),
	(79821, 'BOLLY', 6539, 472, 34),
	(79822, 'DONGGEM', 6539, 472, 34),
	(79823, 'DURIMA', 6539, 472, 34),
	(79824, 'KANERE', 6539, 472, 34),
	(79825, 'KANEWUNUK', 6539, 472, 34),
	(79826, 'KUREWUNUK', 6539, 472, 34),
	(79827, 'LEREWERE', 6539, 472, 34),
	(79828, 'MUNAGAME', 6539, 472, 34),
	(79829, 'NANGGURILIME', 6539, 472, 34),
	(79830, 'NUNGGALO', 6539, 472, 34),
	(79831, 'OMUK', 6539, 472, 34),
	(79832, 'TANABUME', 6539, 472, 34),
	(79833, 'WARINGGA', 6539, 472, 34),
	(79834, 'WERI', 6539, 472, 34),
	(79835, 'WONAGA', 6539, 472, 34),
	(79836, 'AMBENA', 6540, 472, 34),
	(79837, 'BUMBU', 6540, 472, 34),
	(79838, 'DELEGARI', 6540, 472, 34),
	(79839, 'GUNOMBO', 6540, 472, 34),
	(79840, 'MAWI', 6540, 472, 34),
	(79841, 'MILIPAA', 6540, 472, 34),
	(79842, 'NIAGALE', 6540, 472, 34),
	(79843, 'RIPA', 6540, 472, 34),
	(79844, 'TARAWI', 6540, 472, 34),
	(79845, 'WANIA', 6540, 472, 34),
	(79846, 'BIRE', 6541, 472, 34),
	(79847, 'DAGARI', 6541, 472, 34),
	(79848, 'DOW/BIJERE', 6541, 472, 34),
	(79849, 'ITOLI', 6541, 472, 34),
	(79850, 'PAKARE', 6541, 472, 34),
	(79851, 'PRAWA', 6541, 472, 34),
	(79852, 'SIGOU', 6541, 472, 34),
	(79853, 'TAKRI', 6541, 472, 34),
	(79854, 'TIGU', 6541, 472, 34),
	(79855, 'VOKUYO', 6541, 472, 34),
	(79856, 'WARKA', 6541, 472, 34),
	(79857, 'BIMO', 6542, 472, 34),
	(79858, 'DUGUNAGEP', 6542, 472, 34),
	(79859, 'DUNDU', 6542, 472, 34),
	(79860, 'KEMBU', 6542, 472, 34),
	(79861, 'KURUPU', 6542, 472, 34),
	(79862, 'NAKWI', 6542, 472, 34),
	(79863, 'NILOGABU', 6542, 472, 34),
	(79864, 'NINI', 6542, 472, 34),
	(79865, 'NUGINI', 6542, 472, 34),
	(79866, 'YIKU', 6542, 472, 34),
	(79867, 'EGIAM', 6543, 472, 34),
	(79868, 'KALIUNDI', 6543, 472, 34),
	(79869, 'KURBA', 6543, 472, 34),
	(79870, 'MURNI', 6543, 472, 34),
	(79871, 'PINDE', 6543, 472, 34),
	(79872, 'TABONAKME', 6543, 472, 34),
	(79873, 'WAYONGGA', 6543, 472, 34),
	(79874, 'WERI', 6543, 472, 34),
	(79875, 'YOKA', 6543, 472, 34),
	(79876, 'YONIRA', 6543, 472, 34),
	(79877, 'ALOBAGA', 6544, 472, 34),
	(79878, 'DIMBARA', 6544, 472, 34),
	(79879, 'GEYA', 6544, 472, 34),
	(79880, 'JELEPELE', 6544, 472, 34),
	(79881, 'KIBU', 6544, 472, 34),
	(79882, 'NAWU', 6544, 472, 34),
	(79883, 'TIMORI', 6544, 472, 34),
	(79884, 'TINAGOGA', 6544, 472, 34),
	(79885, 'WEYAMBI', 6544, 472, 34),
	(79886, 'WINALO', 6544, 472, 34),
	(79887, 'WITIPUR', 6544, 472, 34),
	(79888, 'WUNGGILIPUR', 6544, 472, 34),
	(79889, 'DIMBARA', 6545, 472, 34),
	(79890, 'GEKA', 6545, 472, 34),
	(79891, 'GELOK', 6545, 472, 34),
	(79892, 'KWA', 6545, 472, 34),
	(79893, 'MAKIDO', 6545, 472, 34),
	(79894, 'MEMBRAMONGGEN', 6545, 472, 34),
	(79895, 'TABUNAKME', 6545, 472, 34),
	(79896, 'WANUWI', 6545, 472, 34),
	(79897, 'WENIGI', 6545, 472, 34),
	(79898, 'YINUWANU', 6545, 472, 34),
	(79899, 'BAGUNI', 6546, 472, 34),
	(79900, 'EGONI', 6546, 472, 34),
	(79901, 'KULUTIN', 6546, 472, 34),
	(79902, 'LEREWERE', 6546, 472, 34),
	(79903, 'MARTELO', 6546, 472, 34),
	(79904, 'ORELUKBAN', 6546, 472, 34),
	(79905, 'TINGGOM', 6546, 472, 34),
	(79906, 'WELESI', 6546, 472, 34),
	(79907, 'YAKEP', 6546, 472, 34),
	(79908, 'YAMULO', 6546, 472, 34),
	(79909, 'ANGKASA', 6547, 472, 34),
	(79910, 'BENARI', 6547, 472, 34),
	(79911, 'BINI', 6547, 472, 34),
	(79912, 'BOPA (BOBE)', 6547, 472, 34),
	(79913, 'DIDELONIK', 6547, 472, 34),
	(79914, 'DOGE', 6547, 472, 34),
	(79915, 'DUGI', 6547, 472, 34),
	(79916, 'GILOK', 6547, 472, 34),
	(79917, 'GOYAGE', 6547, 472, 34),
	(79918, 'KUMBUR (KUMBU)', 6547, 472, 34),
	(79919, 'MAMPULAGA (NAMPULAGA)', 6547, 472, 34),
	(79920, 'PEKO (PEKU)', 6547, 472, 34),
	(79921, 'TAGIKUN (TIGIKUM)', 6547, 472, 34),
	(79922, 'TIDUR MABUK', 6547, 472, 34),
	(79923, 'TIGIR', 6547, 472, 34),
	(79924, 'TIRI', 6547, 472, 34),
	(79925, 'WIJAMURIK', 6547, 472, 34),
	(79926, 'WOJI', 6547, 472, 34),
	(79927, 'YEMARIMA (YEMARMA)', 6547, 472, 34),
	(79928, 'AWORERA', 6548, 472, 34),
	(79929, 'ENGGAWOGO', 6548, 472, 34),
	(79930, 'GINGGA', 6548, 472, 34),
	(79931, 'GUBUK', 6548, 472, 34),
	(79932, 'GUMBINI', 6548, 472, 34),
	(79933, 'KALARIN', 6548, 472, 34),
	(79934, 'KURIK', 6548, 472, 34),
	(79935, 'MURUNERI', 6548, 472, 34),
	(79936, 'NANGGA', 6548, 472, 34),
	(79937, 'OKER', 6548, 472, 34),
	(79938, 'PUNGGELAK', 6548, 472, 34),
	(79939, 'UMAR', 6548, 472, 34),
	(79940, 'WAMILI', 6548, 472, 34),
	(79941, 'WAMOLO', 6548, 472, 34),
	(79942, 'WINENGGA', 6548, 472, 34),
	(79943, 'WOBE', 6548, 472, 34),
	(79944, 'WORAGA', 6548, 472, 34),
	(79945, 'AGAIN', 6549, 472, 34),
	(79946, 'BAWI', 6549, 472, 34),
	(79947, 'KAIGA', 6549, 472, 34),
	(79948, 'KOTORAMBUR', 6549, 472, 34),
	(79949, 'KURBAYA', 6549, 472, 34),
	(79950, 'PIRALEME', 6549, 472, 34),
	(79951, 'TINA', 6549, 472, 34),
	(79952, 'TUNGGUNALE', 6549, 472, 34),
	(79953, 'WIYANGGER', 6549, 472, 34),
	(79954, 'WOLU', 6549, 472, 34),
	(79955, 'BEREMBANAK', 6550, 472, 34),
	(79956, 'HABAG', 6550, 472, 34),
	(79957, 'KALONIKI', 6550, 472, 34),
	(79958, 'KAMBONIKI', 6550, 472, 34),
	(79959, 'KEKOLI', 6550, 472, 34),
	(79960, 'MALTA', 6550, 472, 34),
	(79961, 'MARBUNA', 6550, 472, 34),
	(79962, 'TARI', 6550, 472, 34),
	(79963, 'AULANI', 6551, 472, 34),
	(79964, 'DUNDU', 6551, 472, 34),
	(79965, 'KAGIMALUK', 6551, 472, 34),
	(79966, 'KANGGIME', 6551, 472, 34),
	(79967, 'KERENA', 6551, 472, 34),
	(79968, 'LAWOR (LAWORI)', 6551, 472, 34),
	(79969, 'LIGIIBAK / LIGIMBAK (LAGIMBAK)', 6551, 472, 34),
	(79970, 'LOGON', 6551, 472, 34),
	(79971, 'MARLO', 6551, 472, 34),
	(79972, 'PURUGI', 6551, 472, 34),
	(79973, 'AMPERA', 6552, 472, 34),
	(79974, 'BELEME', 6552, 472, 34),
	(79975, 'DANGGULURIK', 6552, 472, 34),
	(79976, 'EBENHAISER', 6552, 472, 34),
	(79977, 'ELSADAI', 6552, 472, 34),
	(79978, 'GININGGADONAK', 6552, 472, 34),
	(79979, 'GURIKAGEWA', 6552, 472, 34),
	(79980, 'GURIKME', 6552, 472, 34),
	(79981, 'KARUBAGA', 6552, 472, 34),
	(79982, 'KIMOBUR', 6552, 472, 34),
	(79983, 'KIRANAGE', 6552, 472, 34),
	(79984, 'KOGIMAGI', 6552, 472, 34),
	(79985, 'KOLILAN', 6552, 472, 34),
	(79986, 'KULONAME', 6552, 472, 34),
	(79987, 'KURAGEPURA', 6552, 472, 34),
	(79988, 'LIRAK', 6552, 472, 34),
	(79989, 'LOSMEN', 6552, 472, 34),
	(79990, 'LUWIK', 6552, 472, 34),
	(79991, 'MOLERA', 6552, 472, 34),
	(79992, 'MUARA', 6552, 472, 34),
	(79993, 'NALORINI', 6552, 472, 34),
	(79994, 'PULANGGUN', 6552, 472, 34),
	(79995, 'YALIKALUK', 6552, 472, 34),
	(79996, 'AGIMDEK', 6553, 472, 34),
	(79997, 'AWORERA', 6553, 472, 34),
	(79998, 'GENANI', 6553, 472, 34),
	(79999, 'KABORI', 6553, 472, 34),
	(80000, 'KEMBU', 6553, 472, 34)
	`)
}
