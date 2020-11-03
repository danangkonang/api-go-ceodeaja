package kel70

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Kel76() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(76001, 'KAMISABE', 6126, 444, 33),
	(76002, 'MOSWAREN', 6126, 444, 33),
	(76003, 'TOKASS', 6126, 444, 33),
	(76004, 'BOTAIN', 6127, 444, 33),
	(76005, 'KAYABO', 6127, 444, 33),
	(76006, 'KENAYA', 6127, 444, 33),
	(76007, 'KOMANGGARET', 6127, 444, 33),
	(76008, 'KWOWOK', 6127, 444, 33),
	(76009, 'MANGGROHOLO', 6127, 444, 33),
	(76010, 'MLASWAT', 6127, 444, 33),
	(76011, 'SAYAL', 6127, 444, 33),
	(76012, 'SIRA', 6127, 444, 33),
	(76013, 'SISIR', 6127, 444, 33),
	(76014, 'ALMA', 6128, 444, 33),
	(76015, 'ELLES', 6128, 444, 33),
	(76016, 'KLAMIT', 6128, 444, 33),
	(76017, 'KOFALIT', 6128, 444, 33),
	(76018, 'MLABOLO', 6128, 444, 33),
	(76019, 'SASNEK', 6128, 444, 33),
	(76020, 'SAWIAT', 6128, 444, 33),
	(76021, 'SFAKYO', 6128, 444, 33),
	(76022, 'SODROFOYO', 6128, 444, 33),
	(76023, 'WEEN', 6128, 444, 33),
	(76024, 'WENDI', 6128, 444, 33),
	(76025, 'WENSLOLO', 6128, 444, 33),
	(76026, 'WENSOUGH', 6128, 444, 33),
	(76027, 'HAHA', 6129, 444, 33),
	(76028, 'KAKAS', 6129, 444, 33),
	(76029, 'KAMARO', 6129, 444, 33),
	(76030, 'KLAOGIN', 6129, 444, 33),
	(76031, 'SBIR', 6129, 444, 33),
	(76032, 'SRER', 6129, 444, 33),
	(76033, 'TOFOT', 6129, 444, 33),
	(76034, 'WOLOIN', 6129, 444, 33),
	(76035, 'AIBOBOR', 6130, 444, 33),
	(76036, 'ANI SESNA', 6130, 444, 33),
	(76037, 'GOROLO', 6130, 444, 33),
	(76038, 'KAIBUS', 6130, 444, 33),
	(76039, 'KEYEN', 6130, 444, 33),
	(76040, 'KOHOIN', 6130, 444, 33),
	(76041, 'MAGIS', 6130, 444, 33),
	(76042, 'NAMBRO', 6130, 444, 33),
	(76043, 'SEYOLO', 6130, 444, 33),
	(76044, 'SIRIBAU', 6130, 444, 33),
	(76045, 'TAPIRI', 6130, 444, 33),
	(76046, 'TEGIROLO', 6130, 444, 33),
	(76047, 'WEHALI', 6130, 444, 33),
	(76048, 'WERMIT', 6130, 444, 33),
	(76049, 'WERNAS', 6130, 444, 33),
	(76050, 'WERSAR', 6130, 444, 33),
	(76051, 'BAGRAGA (BAGARAGA)', 6131, 444, 33),
	(76052, 'BALDON (BOLDON)', 6131, 444, 33),
	(76053, 'SESOR', 6131, 444, 33),
	(76054, 'SUNGGUER', 6131, 444, 33),
	(76055, 'UNGI (UNGGI)', 6131, 444, 33),
	(76056, 'WAIGO', 6131, 444, 33),
	(76057, 'WARDIK', 6131, 444, 33),
	(76058, 'WAYER', 6131, 444, 33),
	(76059, 'SAUBEBA', 6132, 445, 33),
	(76060, 'WAIBEM (WAIBEN)', 6132, 445, 33),
	(76061, 'WARMANDI', 6132, 445, 33),
	(76062, 'WAU', 6132, 445, 33),
	(76063, 'WEYAF', 6132, 445, 33),
	(76064, 'BONDOPI', 6133, 445, 33),
	(76065, 'MANGGANEK (ARUPI)', 6133, 445, 33),
	(76066, 'SASUY (SASUI)', 6133, 445, 33),
	(76067, 'SERAYO', 6133, 445, 33),
	(76068, 'SOUKOREM (SAUKOREM)', 6133, 445, 33),
	(76069, 'WASARAK', 6133, 445, 33),
	(76070, 'WEFIANI', 6133, 445, 33),
	(76071, 'ASES', 6134, 445, 33),
	(76072, 'FEF', 6134, 445, 33),
	(76073, 'MAWOR', 6134, 445, 33),
	(76074, 'SIKOR', 6134, 445, 33),
	(76075, 'SYUBUN', 6134, 445, 33),
	(76076, 'WAYO', 6134, 445, 33),
	(76077, 'AJAMI', 6135, 445, 33),
	(76078, 'AKMURI', 6135, 445, 33),
	(76079, 'ATAI', 6135, 445, 33),
	(76080, 'INAM', 6135, 445, 33),
	(76081, 'INAMBUARI', 6135, 445, 33),
	(76082, 'INJAI', 6135, 445, 33),
	(76083, 'JANDURAU', 6135, 445, 33),
	(76084, 'NEKORI', 6135, 445, 33),
	(76085, 'BATDE (BADDEI)', 6136, 445, 33),
	(76086, 'HOPMARE', 6136, 445, 33),
	(76087, 'KRANFOTSU', 6136, 445, 33),
	(76088, 'KRISNOS', 6136, 445, 33),
	(76089, 'KWESEFO (KOSEFA)', 6136, 445, 33),
	(76090, 'KWOOR', 6136, 445, 33),
	(76091, 'SYUAU', 6136, 445, 33),
	(76092, 'SYUKWES', 6136, 445, 33),
	(76093, 'SYUMBAB', 6136, 445, 33),
	(76094, 'AIBOGIAR', 6137, 445, 33),
	(76095, 'AYAE', 6137, 445, 33),
	(76096, 'AYIAMANE', 6137, 445, 33),
	(76097, 'MEIS', 6137, 445, 33),
	(76098, 'MIRI/ IBIAH', 6137, 445, 33),
	(76099, 'REVEWES (RUVEWES)', 6137, 445, 33),
	(76100, 'RUF', 6137, 445, 33),
	(76101, 'SIAKWA', 6137, 445, 33),
	(76102, 'TABAMSERE', 6137, 445, 33),
	(76103, 'YABUOW (JABOUW)', 6137, 445, 33),
	(76104, 'DELA (DELLA)', 6138, 445, 33),
	(76105, 'MALAWARSAI (MALAWORSAI)', 6138, 445, 33),
	(76106, 'MEGA', 6138, 445, 33),
	(76107, 'SALEWOK (SELEWOK)', 6138, 445, 33),
	(76108, 'SELEKEBU (SELEKOPI)', 6138, 445, 33),
	(76109, 'SENGKEDUK (SAENGKEDUK)', 6138, 445, 33),
	(76110, 'SIWIS', 6138, 445, 33),
	(76111, 'ARFU', 6139, 445, 33),
	(76112, 'ATORI', 6139, 445, 33),
	(76113, 'BAWEY', 6139, 445, 33),
	(76114, 'BIJANFOU (BIJAMFOU/BUANFOU)', 6139, 445, 33),
	(76115, 'MERIAMBEKE (MERIAMBEKER)', 6139, 445, 33),
	(76116, 'WARU', 6139, 445, 33),
	(76117, 'WASNEBRI (WARSNEMBRI/WASNEMBRI)', 6139, 445, 33),
	(76118, 'BIKAR', 6140, 445, 33),
	(76119, 'EMAOS', 6140, 445, 33),
	(76120, 'JOKTE', 6140, 445, 33),
	(76121, 'SAUSAPOR', 6140, 445, 33),
	(76122, 'UIGWEM', 6140, 445, 33),
	(76123, 'WERTAM', 6140, 445, 33),
	(76124, 'WERUR', 6140, 445, 33),
	(76125, 'WERUR BESAR', 6140, 445, 33),
	(76126, 'WERWAF', 6140, 445, 33),
	(76127, 'AFRAWI', 6141, 445, 33),
	(76128, 'SENOPI', 6141, 445, 33),
	(76129, 'WAUSIN (ASITI)', 6141, 445, 33),
	(76130, 'BANSO', 6142, 445, 33),
	(76131, 'FRAFANE', 6142, 445, 33),
	(76132, 'IOF', 6142, 445, 33),
	(76133, 'SOON', 6142, 445, 33),
	(76134, 'SYUJAK (SUJAK)', 6142, 445, 33),
	(76135, 'BAMUS', 6143, 445, 33),
	(76136, 'BAMUS WAIMAN (BAMUSWAYMAN)', 6143, 445, 33),
	(76137, 'BAUN (BAUM)', 6143, 445, 33),
	(76138, 'MEFNANYAM (METNAYAM)', 6143, 445, 33),
	(76139, 'METBELUM', 6143, 445, 33),
	(76140, 'METBESA', 6143, 445, 33),
	(76141, 'SUMBEKAS', 6143, 445, 33),
	(76142, 'SYARWOM', 6143, 445, 33),
	(76143, 'ARANDAY (ARANDAI)', 6144, 446, 33),
	(76144, 'BARU (KAMPUNG BARU)', 6144, 446, 33),
	(76145, 'KECAP', 6144, 446, 33),
	(76146, 'MANUNGGAL KARYA/JAYA', 6144, 446, 33),
	(76147, 'AROBA', 6145, 446, 33),
	(76148, 'SANGGUAR', 6145, 446, 33),
	(76149, 'SIDO MAKMUR', 6145, 446, 33),
	(76150, 'WIMBRO (NELAYAN WIMBRO)', 6145, 446, 33),
	(76151, 'YARU', 6145, 446, 33),
	(76152, 'AMUTU', 6146, 446, 33),
	(76153, 'IRARUTU III', 6146, 446, 33),
	(76154, 'KASIRA', 6146, 446, 33),
	(76155, 'NUSEI', 6146, 446, 33),
	(76156, 'ARGOSIGEMERAI', 6147, 446, 33),
	(76157, 'BEIMES', 6147, 446, 33),
	(76158, 'BINTUNI BARAT', 6147, 446, 33),
	(76159, 'BINTUNI TIMUR', 6147, 446, 33),
	(76160, 'IGURIJI', 6147, 446, 33),
	(76161, 'MASINA', 6147, 446, 33),
	(76162, 'TUASAI', 6147, 446, 33),
	(76163, 'WESIRI', 6147, 446, 33),
	(76164, 'ENIBA', 6148, 446, 33),
	(76165, 'IBORI', 6148, 446, 33),
	(76166, 'JAHABRA', 6148, 446, 33),
	(76167, 'LAUDOHO', 6148, 446, 33),
	(76168, 'MENYEMBRUI (MENYEMBRU)', 6148, 446, 33),
	(76169, 'MEYORGA', 6148, 446, 33),
	(76170, 'MOWITKA', 6148, 446, 33),
	(76171, 'CUMNAJI', 6149, 446, 33),
	(76172, 'HORNA (HOMA)', 6149, 446, 33),
	(76173, 'HUSS', 6149, 446, 33),
	(76174, 'MENCI', 6149, 446, 33),
	(76175, 'SIR', 6149, 446, 33),
	(76176, 'UGDOHOP', 6149, 446, 33),
	(76177, 'FRUATA (FUATA/IRORUTU II)', 6150, 446, 33),
	(76178, 'MERYEDI (MARYEDI)', 6150, 446, 33),
	(76179, 'RIENDO', 6150, 446, 33),
	(76180, 'SARA', 6151, 446, 33),
	(76181, 'SUGA', 6151, 446, 33),
	(76182, 'TUGARAMA (TUGERAMA)', 6151, 446, 33),
	(76183, 'WARGA NUSA II', 6151, 446, 33),
	(76184, 'WARGANUSA I (WARBANUSA I)', 6151, 446, 33),
	(76185, 'BIBIRAM', 6152, 446, 33),
	(76186, 'KALITAMA I (KALITAMI I)', 6152, 446, 33),
	(76187, 'KALITAMA II (KALITAMI II)', 6152, 446, 33),
	(76188, 'KENARA', 6152, 446, 33),
	(76189, 'NARAMASA', 6153, 446, 33),
	(76190, 'OBO', 6153, 446, 33),
	(76191, 'REFIDESO', 6153, 446, 33),
	(76192, 'SARBE', 6153, 446, 33),
	(76193, 'WAGURA', 6153, 446, 33),
	(76194, 'ATIBO MANIMERI', 6154, 446, 33),
	(76195, 'BANJAR AUSOY', 6154, 446, 33),
	(76196, 'BUMI SANIARI', 6154, 446, 33),
	(76197, 'KORANO JAYA', 6154, 446, 33),
	(76198, 'PASAMAI', 6154, 446, 33),
	(76199, 'WARAITAMA', 6154, 446, 33),
	(76200, 'KALIBIRU', 6155, 446, 33),
	(76201, 'MAESTOFU (MESTOFU)', 6155, 446, 33),
	(76202, 'MASYETA', 6155, 446, 33),
	(76203, 'MESOMDA', 6155, 446, 33),
	(76204, 'BARMA', 6156, 446, 33),
	(76205, 'BARMA BARU (STENCOOL)', 6156, 446, 33),
	(76206, 'MEYADO (MAYADO/MEYABO)', 6156, 446, 33),
	(76207, 'VASCO DAMNEEN (VESCO DAMNEM/STEN)', 6156, 446, 33),
	(76208, 'ANAJERO', 6157, 446, 33),
	(76209, 'MEKIESEFEB', 6157, 446, 33),
	(76210, 'MENGGERBA', 6157, 446, 33),
	(76211, 'MERDEY', 6157, 446, 33),
	(76212, 'MERYEB', 6157, 446, 33),
	(76213, 'MEYETGA', 6157, 446, 33),
	(76214, 'MEYOM', 6157, 446, 33),
	(76215, 'MOGROMUS', 6157, 446, 33),
	(76216, 'MOROMBUY', 6157, 446, 33),
	(76217, 'ISTIWKEM (ISTIKEM)', 6158, 446, 33),
	(76218, 'MACOK', 6158, 446, 33),
	(76219, 'MEJNIC (MAJNIC/MAJNIK)', 6158, 446, 33),
	(76220, 'MEYERGA (MAYERGA)', 6158, 446, 33),
	(76221, 'BARMA BARAT', 6159, 446, 33),
	(76222, 'INGGOF', 6159, 446, 33),
	(76223, 'JAGIRO', 6159, 446, 33),
	(76224, 'MEYENDA', 6159, 446, 33),
	(76225, 'RAWARA', 6159, 446, 33),
	(76226, 'IGOMU', 6160, 446, 33),
	(76227, 'MESNA', 6160, 446, 33),
	(76228, 'SUMUY (SUMUI)', 6160, 446, 33),
	(76229, 'INOFINA', 6161, 446, 33),
	(76230, 'MERESTIM (MARISTIM)', 6161, 446, 33),
	(76231, 'MOSUM', 6161, 446, 33),
	(76232, 'MOYEBA', 6161, 446, 33),
	(76233, 'FORADA', 6162, 446, 33),
	(76234, 'MATERABU JAYA', 6162, 446, 33),
	(76235, 'SAENGGA', 6162, 446, 33),
	(76236, 'TANAH MERAH', 6162, 446, 33),
	(76237, 'TOFOI', 6162, 446, 33),
	(76238, 'ARAISUM', 6163, 446, 33),
	(76239, 'BANGUN MULYO/MULIA', 6163, 446, 33),
	(76240, 'MOGOI BARU', 6163, 446, 33),
	(76241, 'TEMBUNI', 6163, 446, 33),
	(76242, 'EKAM', 6164, 446, 33),
	(76243, 'SEBYAR REJOSASI/REJOSARI', 6164, 446, 33),
	(76244, 'TAROY (TORAY/TAROI)', 6164, 446, 33),
	(76245, 'TOMU', 6164, 446, 33),
	(76246, 'KUCIR', 6165, 446, 33),
	(76247, 'SIBENA PERMAI (II)', 6165, 446, 33),
	(76248, 'SIBENA RAYA (I)', 6165, 446, 33),
	(76249, 'TISAIDA', 6165, 446, 33),
	(76250, 'TUHIBA', 6165, 446, 33),
	(76251, 'MAMURANU', 6166, 446, 33),
	(76252, 'WAMESA I - IDOOR', 6166, 446, 33),
	(76253, 'WAMESA II - YAKATI', 6166, 446, 33),
	(76254, 'YANSEI (YANSEY/ENSEI/YENSEI)', 6166, 446, 33),
	(76255, 'MOGOTIRA (MOGMESA)', 6167, 446, 33),
	(76256, 'TUANAIKIN', 6167, 446, 33),
	(76257, 'WERIAGAR', 6167, 446, 33),
	(76258, 'WERIAGAR BARU', 6167, 446, 33),
	(76259, 'WERIAGAR UTARA', 6167, 446, 33),
	(76260, 'AMBUMI', 6168, 447, 33),
	(76261, 'DUSNER', 6168, 447, 33),
	(76262, 'NANIMORI', 6168, 447, 33),
	(76263, 'SIMEI (SIMIEI)', 6168, 447, 33),
	(76264, 'SOBIAR', 6168, 447, 33),
	(76265, 'YERENUSI', 6168, 447, 33),
	(76266, 'INYORA', 6169, 447, 33),
	(76267, 'OYA (OYAA)', 6169, 447, 33),
	(76268, 'SARARTI', 6169, 447, 33),
	(76269, 'UNDURARA', 6169, 447, 33),
	(76270, 'WOSIMI (WOSIMO)', 6169, 447, 33),
	(76271, 'YABORE', 6169, 447, 33),
	(76272, 'KUREI', 6170, 447, 33),
	(76273, 'MAMISI', 6170, 447, 33),
	(76274, 'TAMOGE', 6170, 447, 33),
	(76275, 'WERABUR', 6170, 447, 33),
	(76276, 'WERIANGGI', 6170, 447, 33),
	(76277, 'ISEY (ISEI)', 6171, 447, 33),
	(76278, 'NGGATUM', 6171, 447, 33),
	(76279, 'RASIEY (RASIEI)', 6171, 447, 33),
	(76280, 'SASIREY (SASIREI)', 6171, 447, 33),
	(76281, 'SENDERAWOY / SENDERAWOI (SINDRAWOI)', 6171, 447, 33),
	(76282, 'TANDIA', 6171, 447, 33),
	(76283, 'TOREY', 6171, 447, 33),
	(76284, 'URYEMI (URIYEM/URIEMI)', 6171, 447, 33),
	(76285, 'WEBI', 6171, 447, 33),
	(76286, 'INDAY', 6172, 447, 33),
	(76287, 'MENA', 6172, 447, 33),
	(76288, 'MENARBU', 6172, 447, 33),
	(76289, 'NIAB', 6172, 447, 33),
	(76290, 'SARIAY', 6172, 447, 33),
	(76291, 'SYABES', 6172, 447, 33),
	(76292, 'YENDE', 6172, 447, 33),
	(76293, 'SYEIWAR', 6173, 447, 33),
	(76294, 'WAPRAK', 6173, 447, 33),
	(76295, 'YOMBER', 6173, 447, 33),
	(76296, 'ISENEBUAI (SENEBUAY/SENEBUAI)', 6174, 447, 33),
	(76297, 'ISEREN (ISREN)', 6174, 447, 33),
	(76298, 'WETITINDAU', 6174, 447, 33),
	(76299, 'YARIARI', 6174, 447, 33),
	(76300, 'YEMBEKIRI II', 6174, 447, 33),
	(76301, 'YOMAKAN', 6174, 447, 33),
	(76302, 'YOMBEKIRI I (YEMBE KIRI)', 6174, 447, 33),
	(76303, 'KAPRUS', 6175, 447, 33),
	(76304, 'NUSPAIRO', 6175, 447, 33),
	(76305, 'REYOB (RAYOB)', 6175, 447, 33),
	(76306, 'SIRESI', 6175, 447, 33),
	(76307, 'YARMATUM (YERMATUM)', 6175, 447, 33),
	(76308, 'AISANDAMI', 6176, 447, 33),
	(76309, 'SOBEI (SOBEY)', 6176, 447, 33),
	(76310, 'SOBEY INDAH', 6176, 447, 33),
	(76311, 'YOPANGGAR (YOPENGAR)', 6176, 447, 33),
	(76312, 'AMBUAR', 6177, 447, 33),
	(76313, 'KARUAN', 6177, 447, 33),
	(76314, 'NORDIWAR', 6177, 447, 33),
	(76315, 'SABUBAR', 6177, 447, 33),
	(76316, 'YARPATE', 6177, 447, 33),
	(76317, 'DOTIR', 6178, 447, 33),
	(76318, 'IRIATI', 6178, 447, 33),
	(76319, 'MAIMARE (MAIMARI)', 6178, 447, 33),
	(76320, 'MANIWAK', 6178, 447, 33),
	(76321, 'MORU (GAYABARU)', 6178, 447, 33),
	(76322, 'RADO', 6178, 447, 33),
	(76323, 'RAMIKI', 6178, 447, 33),
	(76324, 'WASIOR I', 6178, 447, 33),
	(76325, 'WASIOR II', 6178, 447, 33),
	(76326, 'WONDAMAWI', 6178, 447, 33),
	(76327, 'ARYOBU', 6179, 447, 33),
	(76328, 'ASAI', 6179, 447, 33),
	(76329, 'KAIRAWI (KARAWI)', 6179, 447, 33),
	(76330, 'KAONDA', 6179, 447, 33),
	(76331, 'MUNGGUI', 6179, 447, 33),
	(76332, 'ROSBORI', 6179, 447, 33),
	(76333, 'SANDEY', 6179, 447, 33),
	(76334, 'SARUMAN', 6179, 447, 33),
	(76335, 'SOMBOKORO', 6179, 447, 33),
	(76336, 'WAISANI', 6179, 447, 33),
	(76337, 'WAMESA TENGAH', 6179, 447, 33),
	(76338, 'WINDESI', 6179, 447, 33),
	(76339, 'WINDESI', 6179, 447, 33),
	(76340, 'YOPMIOS', 6179, 447, 33),
	(76341, 'ISUY (ISSUI)', 6180, 447, 33),
	(76342, 'KABOUW (KABUOW)', 6180, 447, 33),
	(76343, 'KAIBI', 6180, 447, 33),
	(76344, 'MANOPI', 6180, 447, 33),
	(76345, 'WONDIBOI (WONDIBOY)', 6180, 447, 33),
	(76346, 'ASUWETSY', 6181, 448, 34),
	(76347, 'BIS AGATS', 6181, 448, 34),
	(76348, 'BISMAN', 6181, 448, 34),
	(76349, 'BOU', 6181, 448, 34),
	(76350, 'BRITEN (BIRITEN/BERITEN)', 6181, 448, 34),
	(76351, 'KAYE', 6181, 448, 34),
	(76352, 'MBAIT', 6181, 448, 34),
	(76353, 'PER', 6181, 448, 34),
	(76354, 'SAW', 6181, 448, 34),
	(76355, 'SUWRU', 6181, 448, 34),
	(76356, 'UWUS', 6181, 448, 34),
	(76357, 'YAMOTH (YOMOTH)', 6181, 448, 34),
	(76358, 'AKAT (AYAM)', 6182, 448, 34),
	(76359, 'BAYIW PINAM', 6182, 448, 34),
	(76360, 'BECO', 6182, 448, 34),
	(76361, 'BUETKWAR (BETKUAR)', 6182, 448, 34),
	(76362, 'CUMNEW', 6182, 448, 34),
	(76363, 'FAKAN', 6182, 448, 34),
	(76364, 'JEWES', 6182, 448, 34),
	(76365, 'MANEPSIMNI (MANEP)', 6182, 448, 34),
	(76366, 'SIMINI', 6182, 448, 34),
	(76367, 'WAW', 6182, 448, 34),
	(76368, 'YUNI', 6182, 448, 34),
	(76369, 'AMANAMKAI', 6183, 448, 34),
	(76370, 'AMBISU', 6183, 448, 34),
	(76371, 'ATSJ (ATSY)', 6183, 448, 34),
	(76372, 'BAKASEI', 6183, 448, 34),
	(76373, 'BINE', 6183, 448, 34),
	(76374, 'BIPIN (BIPIM)', 6183, 448, 34),
	(76375, 'CEWEWYAMEW', 6183, 448, 34),
	(76376, 'SAGONI / SOGONI', 6183, 448, 34),
	(76377, 'YAISU (YASIU)', 6183, 448, 34),
	(76378, 'COMORO', 6184, 448, 34),
	(76379, 'KAWET', 6184, 448, 34),
	(76380, 'MAUSI', 6184, 448, 34),
	(76381, 'SAGARE', 6184, 448, 34),
	(76382, 'WAGI', 6184, 448, 34),
	(76383, 'YEFIWANGI (YEFUWAGI)', 6184, 448, 34),
	(76384, 'ATAMBUTS', 6185, 448, 34),
	(76385, 'BIWAR LAUT', 6185, 448, 34),
	(76386, 'DESEP', 6185, 448, 34),
	(76387, 'OMANESEP', 6185, 448, 34),
	(76388, 'PIRPIS', 6185, 448, 34),
	(76389, 'WARKAL (WARKAI)', 6185, 448, 34),
	(76390, 'YAUW (YOU)', 6185, 448, 34),
	(76391, 'AMAGAIS', 6186, 448, 34),
	(76392, 'AMARU (AMORO)', 6186, 448, 34),
	(76393, 'AMKAI', 6186, 448, 34),
	(76394, 'AMKUN / AMKUM (ANKUM)', 6186, 448, 34),
	(76395, 'ERO SAMAN', 6186, 448, 34),
	(76396, 'SOHOMANE', 6186, 448, 34),
	(76397, 'SUAGAI / SUWAGAI (SUWA)', 6186, 448, 34),
	(76398, 'YAMKAP (YANKAP)', 6186, 448, 34),
	(76399, 'YERFUM', 6186, 448, 34),
	(76400, 'ACENEP (OCENEP)', 6187, 448, 34),
	(76401, 'AINAMSATO', 6187, 448, 34),
	(76402, 'AIS', 6187, 448, 34),
	(76403, 'AKAN TAPAK', 6187, 448, 34),
	(76404, 'AMAITA', 6187, 448, 34),
	(76405, 'BAGAIR', 6187, 448, 34),
	(76406, 'BAKYOR', 6187, 448, 34),
	(76407, 'BASIM', 6187, 448, 34),
	(76408, 'BAWOS', 6187, 448, 34),
	(76409, 'BIOPIS', 6187, 448, 34),
	(76410, 'BORA', 6187, 448, 34),
	(76411, 'ISAR', 6187, 448, 34),
	(76412, 'KAGAS', 6187, 448, 34),
	(76413, 'KAYARPIS', 6187, 448, 34),
	(76414, 'MAPANE', 6187, 448, 34),
	(76415, 'NANAI (NANAY)', 6187, 448, 34),
	(76416, 'PIRAMAT', 6187, 448, 34),
	(76417, 'PIRIEN', 6187, 448, 34),
	(76418, 'SAYOA', 6187, 448, 34),
	(76419, 'TAURO', 6187, 448, 34),
	(76420, 'WARAS', 6187, 448, 34),
	(76421, 'WIYAR', 6187, 448, 34),
	(76422, 'YAWAS', 6187, 448, 34),
	(76423, 'AKAMAR', 6188, 448, 34),
	(76424, 'BIRAK', 6188, 448, 34),
	(76425, 'DAWER', 6188, 448, 34),
	(76426, 'KATEW', 6188, 448, 34),
	(76427, 'PAU', 6188, 448, 34),
	(76428, 'POWETSY', 6188, 448, 34),
	(76429, 'SISAKEM', 6188, 448, 34),
	(76430, 'YETSY', 6188, 448, 34),
	(76431, 'OMOR', 6189, 448, 34),
	(76432, 'ONAVAI', 6189, 448, 34),
	(76433, 'YAMAS', 6189, 448, 34),
	(76434, 'YAUN', 6189, 448, 34),
	(76435, 'YENI / YEMI', 6189, 448, 34),
	(76436, 'YUFRI (SMITH)', 6189, 448, 34),
	(76437, 'AUBAN', 6190, 448, 34),
	(76438, 'BINAMSAIN', 6190, 448, 34),
	(76439, 'BUTUKATNAU', 6190, 448, 34),
	(76440, 'MABUL', 6190, 448, 34),
	(76441, 'PATIPI (PATIPPI)', 6190, 448, 34),
	(76442, 'PEPERA', 6190, 448, 34),
	(76443, 'PIRABANAK', 6190, 448, 34),
	(76444, 'SIPENAP / SIPANAP (ASAREP)', 6190, 448, 34),
	(76445, 'ULAKIN', 6190, 448, 34),
	(76446, 'WOUTU BRASA', 6190, 448, 34),
	(76447, 'WOUTU KOLOF', 6190, 448, 34),
	(76448, 'AIKUT', 6191, 448, 34),
	(76449, 'HAHARE', 6191, 448, 34),
	(76450, 'HEIYARAM', 6191, 448, 34),
	(76451, 'KAIPOM', 6191, 448, 34),
	(76452, 'KAWEM', 6191, 448, 34),
	(76453, 'SANEP / SANEM', 6191, 448, 34),
	(76454, 'SAPEN (AIROSAPEM)', 6191, 448, 34),
	(76455, 'SASIME', 6191, 448, 34),
	(76456, 'SIMIPIT (SINEPIT)', 6191, 448, 34),
	(76457, 'WAGASU', 6191, 448, 34),
	(76458, 'BAWOR', 6192, 448, 34),
	(76459, 'ESEIB (ESAIB)', 6192, 448, 34),
	(76460, 'HAINAM', 6192, 448, 34),
	(76461, 'KAMUR', 6192, 448, 34),
	(76462, 'SANAPAI (SENAPAI)', 6192, 448, 34),
	(76463, 'SARAMIT / SERAMIT', 6192, 448, 34),
	(76464, 'SARMAFO', 6192, 448, 34),
	(76465, 'YAGAMIT', 6192, 448, 34),
	(76466, 'YAHUI (YAHOI)', 6192, 448, 34),
	(76467, 'AOEP', 6193, 448, 34),
	(76468, 'AOU', 6193, 448, 34),
	(76469, 'AS', 6193, 448, 34),
	(76470, 'ATAT', 6193, 448, 34),
	(76471, 'EROKO', 6193, 448, 34),
	(76472, 'ESMAPAN', 6193, 448, 34),
	(76473, 'FUMERIPIT', 6193, 448, 34),
	(76474, 'KAPI', 6193, 448, 34),
	(76475, 'NAKAI', 6193, 448, 34),
	(76476, 'PULAU TIGA', 6193, 448, 34),
	(76477, 'SABANG MAWANG', 6193, 448, 34),
	(76478, 'SABANG MAWANG BARAT', 6193, 448, 34),
	(76479, 'SEDEDAP', 6193, 448, 34),
	(76480, 'SELADING', 6193, 448, 34),
	(76481, 'SERANTAS', 6193, 448, 34),
	(76482, 'SETUMUK', 6193, 448, 34),
	(76483, 'TANJUNG BATANG', 6193, 448, 34),
	(76484, 'TANJUNG KUMBIK UTARA', 6193, 448, 34),
	(76485, 'TELUK LABUH', 6193, 448, 34),
	(76486, 'WEO', 6193, 448, 34),
	(76487, 'YAKAPIS', 6193, 448, 34),
	(76488, 'AWORKEY / AWORKET', 6194, 448, 34),
	(76489, 'BAYUN', 6194, 448, 34),
	(76490, 'EMENE (EUMENE)', 6194, 448, 34),
	(76491, 'JITORSOK', 6194, 448, 34),
	(76492, 'KAGIRIM / KAIRIN', 6194, 448, 34),
	(76493, 'PRIMAPUN (FIRIMAPUN/FIRIMAPON)', 6194, 448, 34),
	(76494, 'SAMAN', 6194, 448, 34),
	(76495, 'SAMENDORO / SEMENDORO', 6194, 448, 34),
	(76496, 'SANTABOR (SANTAMBOR)', 6194, 448, 34),
	(76497, 'SIMSAGAR', 6194, 448, 34),
	(76498, 'TEREO / TAREO (TAERO)', 6194, 448, 34),
	(76499, 'YAPTAMBOR', 6194, 448, 34),
	(76500, 'AGANI', 6195, 448, 34),
	(76501, 'BU', 6195, 448, 34),
	(76502, 'ER', 6195, 448, 34),
	(76503, 'ERMA', 6195, 448, 34),
	(76504, 'MUMUGU', 6195, 448, 34),
	(76505, 'MUMUGU DUA', 6195, 448, 34),
	(76506, 'PUPIS', 6195, 448, 34),
	(76507, 'SAUTI', 6195, 448, 34),
	(76508, 'SAWA', 6195, 448, 34),
	(76509, 'SONA', 6195, 448, 34),
	(76510, 'AWOK', 6196, 448, 34),
	(76511, 'BIWAR DARAT', 6196, 448, 34),
	(76512, 'DAMEN', 6196, 448, 34),
	(76513, 'FOS', 6196, 448, 34),
	(76514, 'KAIMO', 6196, 448, 34),
	(76515, 'SAKOR', 6196, 448, 34),
	(76516, 'WAGANU', 6196, 448, 34),
	(76517, 'YAOSAKOR', 6196, 448, 34),
	(76518, 'AMAKOT', 6197, 448, 34),
	(76519, 'AYAK', 6197, 448, 34),
	(76520, 'BANUM', 6197, 448, 34),
	(76521, 'BINAM', 6197, 448, 34),
	(76522, 'BOR', 6197, 448, 34),
	(76523, 'BUBIS', 6197, 448, 34),
	(76524, 'BUMU', 6197, 448, 34),
	(76525, 'BURBIS', 6197, 448, 34),
	(76526, 'DAIKOT', 6197, 448, 34),
	(76527, 'DEKAMOR / DEKAMER', 6197, 448, 34),
	(76528, 'EMNAM', 6197, 448, 34),
	(76529, 'JINAK', 6197, 448, 34),
	(76530, 'KAPAYAP DUA', 6197, 448, 34),
	(76531, 'KAPAYAP SATU', 6197, 448, 34),
	(76532, 'KAPAYAP TIGA', 6197, 448, 34),
	(76533, 'KARBIS', 6197, 448, 34),
	(76534, 'NAGATUN', 6197, 448, 34),
	(76535, 'SOMNAK', 6197, 448, 34),
	(76536, 'SORAY / SORAI', 6197, 448, 34),
	(76537, 'UJUNG BATU', 6197, 448, 34),
	(76538, 'VAKAM', 6197, 448, 34),
	(76539, 'VAKAM DUA', 6197, 448, 34),
	(76540, 'WABAK', 6197, 448, 34),
	(76541, 'WAGABUS', 6197, 448, 34),
	(76542, 'WAGANU DUA', 6197, 448, 34),
	(76543, 'WAIJENS', 6197, 448, 34),
	(76544, 'WOWI', 6197, 448, 34),
	(76545, 'AJIN', 6198, 448, 34),
	(76546, 'ASGUN', 6198, 448, 34),
	(76547, 'BERIMONO', 6198, 448, 34),
	(76548, 'BESIKA', 6198, 448, 34),
	(76549, 'DUMATEN', 6198, 448, 34),
	(76550, 'HOM-HOM', 6198, 448, 34),
	(76551, 'HULAM', 6198, 448, 34),
	(76552, 'JIFAK', 6198, 448, 34),
	(76553, 'KATALINA', 6198, 448, 34),
	(76554, 'KEBIKDUK', 6198, 448, 34),
	(76555, 'KOBA', 6198, 448, 34),
	(76556, 'KOROBUK', 6198, 448, 34),
	(76557, 'LALUK', 6198, 448, 34),
	(76558, 'LINDUK', 6198, 448, 34),
	(76559, 'OBIO', 6198, 448, 34),
	(76560, 'SAGAPU', 6198, 448, 34),
	(76561, 'SALBIK', 6198, 448, 34),
	(76562, 'SE (SEE)', 6198, 448, 34),
	(76563, 'SURU-SURU', 6198, 448, 34),
	(76564, 'TII', 6198, 448, 34),
	(76565, 'TOMOR', 6198, 448, 34),
	(76566, 'WALASE', 6198, 448, 34),
	(76567, 'YENSUKU (JESENKO)', 6198, 448, 34),
	(76568, 'ABAMU', 6199, 448, 34),
	(76569, 'AMOR', 6199, 448, 34),
	(76570, 'AYIR', 6199, 448, 34),
	(76571, 'BERIP / BIRIP', 6199, 448, 34),
	(76572, 'JIPAWER (YIPAWER)', 6199, 448, 34),
	(76573, 'KOMOR', 6199, 448, 34),
	(76574, 'MUNU', 6199, 448, 34),
	(76575, 'PAAR', 6199, 448, 34),
	(76576, 'WERER (WARER)', 6199, 448, 34),
	(76577, 'ANOBO', 6200, 449, 34),
	(76578, 'KARABAI', 6200, 449, 34),
	(76579, 'MBROMSI', 6200, 449, 34),
	(76580, 'MEOS MANGUANDI (MANGGUANDI)', 6200, 449, 34),
	(76581, 'NYANSOREN', 6200, 449, 34),
	(76582, 'PADAIDO', 6200, 449, 34),
	(76583, 'PASI', 6200, 449, 34),
	(76584, 'SAMBER PASI', 6200, 449, 34),
	(76585, 'SARIBRA', 6200, 449, 34),
	(76586, 'SASARI', 6200, 449, 34),
	(76587, 'SUPRAIMA', 6200, 449, 34),
	(76588, 'YEN MANAINA', 6200, 449, 34),
	(76589, 'YERI', 6200, 449, 34),
	(76590, 'ARMNU', 6201, 449, 34),
	(76591, 'DASDO', 6201, 449, 34),
	(76592, 'FAKNIKDI', 6201, 449, 34),
	(76593, 'MAMORBO', 6201, 449, 34),
	(76594, 'RODIFU', 6201, 449, 34),
	(76595, 'RUMBIN', 6201, 449, 34),
	(76596, 'SUP MBRUR', 6201, 449, 34),
	(76597, 'WARBINSI', 6201, 449, 34),
	(76598, 'WODU', 6201, 449, 34),
	(76599, 'WODU (MAKUKER)', 6201, 449, 34),
	(76600, 'WONABRAIDI', 6201, 449, 34),
	(76601, 'WOUNA', 6201, 449, 34),
	(76602, 'AMPONBUKOR', 6202, 449, 34),
	(76603, 'ANDEI', 6202, 449, 34),
	(76604, 'ASARKIR', 6202, 449, 34),
	(76605, 'ASARYENDI', 6202, 449, 34),
	(76606, 'BANASRARES', 6202, 449, 34),
	(76607, 'DEDIFU', 6202, 449, 34),
	(76608, 'DOUSI', 6202, 449, 34),
	(76609, 'INSIRI', 6202, 449, 34),
	(76610, 'KABABUR', 6202, 449, 34),
	(76611, 'KANAAN', 6202, 449, 34),
	(76612, 'KARNINDI', 6202, 449, 34),
	(76613, 'MAMORIBO', 6202, 449, 34),
	(76614, 'MARISEN', 6202, 449, 34),
	(76615, 'OPURI', 6202, 449, 34),
	(76616, 'RARSBARI', 6202, 449, 34),
	(76617, 'SOPENDO', 6202, 449, 34),
	(76618, 'SOPENDO SUP KARKIR', 6202, 449, 34),
	(76619, 'SOSMAY / SOSMAI', 6202, 449, 34),
	(76620, 'SUMBINYA (SUNBINYA)', 6202, 449, 34),
	(76621, 'WARBERIK (WABERIK)', 6202, 449, 34),
	(76622, 'WASYAI', 6202, 449, 34),
	(76623, 'YOMDORI', 6202, 449, 34),
	(76624, 'ANGGRAIDI', 6203, 449, 34),
	(76625, 'BABRINBO', 6203, 449, 34),
	(76626, 'BUROKUB', 6203, 449, 34),
	(76627, 'FANDOI', 6203, 449, 34),
	(76628, 'INGGIRI', 6203, 449, 34),
	(76629, 'INGGUPI', 6203, 449, 34),
	(76630, 'INSROM', 6203, 449, 34),
	(76631, 'KABABUR', 6203, 449, 34),
	(76632, 'KARYENDI', 6203, 449, 34),
	(76633, 'MANDALA', 6203, 449, 34),
	(76634, 'MANGGANDISAPI', 6203, 449, 34),
	(76635, 'MANSWAN', 6203, 449, 34),
	(76636, 'MNUBABO (AMBROBEN)', 6203, 449, 34),
	(76637, 'MOKMER', 6203, 449, 34),
	(76638, 'PARAY (PARAI)', 6203, 449, 34),
	(76639, 'SAMAU', 6203, 449, 34),
	(76640, 'SANUMI', 6203, 449, 34),
	(76641, 'SARAMOM', 6203, 449, 34),
	(76642, 'SORIDO', 6203, 449, 34),
	(76643, 'SWAPODIBO', 6203, 449, 34),
	(76644, 'WAUPNOR', 6203, 449, 34),
	(76645, 'YENURES', 6203, 449, 34),
	(76646, 'ADORBARI', 6204, 449, 34),
	(76647, 'AFEFBO', 6204, 449, 34),
	(76648, 'BINDUSI', 6204, 449, 34),
	(76649, 'BOSNIK SUP', 6204, 449, 34),
	(76650, 'INMDI', 6204, 449, 34),
	(76651, 'INOFI', 6204, 449, 34),
	(76652, 'INSUMAMIRES (INSUMARIRES)', 6204, 449, 34),
	(76653, 'KAJASBO', 6204, 449, 34),
	(76654, 'KAJASI', 6204, 449, 34),
	(76655, 'MANDON', 6204, 449, 34),
	(76656, 'ORWER', 6204, 449, 34),
	(76657, 'OWI', 6204, 449, 34),
	(76658, 'RIM', 6204, 449, 34),
	(76659, 'RIMBA JAYA', 6204, 449, 34),
	(76660, 'RUAR', 6204, 449, 34),
	(76661, 'SAREIDI (SARAEIDI)', 6204, 449, 34),
	(76662, 'SEEPSE (SEPSE)', 6204, 449, 34),
	(76663, 'SOON', 6204, 449, 34),
	(76664, 'SORYAR', 6204, 449, 34),
	(76665, 'SUNDE', 6204, 449, 34),
	(76666, 'WADERBO', 6204, 449, 34),
	(76667, 'WASORI', 6204, 449, 34),
	(76668, 'WONIKI', 6204, 449, 34),
	(76669, 'YENBEROK', 6204, 449, 34),
	(76670, 'YENDAKEM', 6204, 449, 34),
	(76671, 'YENUSI', 6204, 449, 34),
	(76672, 'ANDEI', 6205, 449, 34),
	(76673, 'DERNAFI', 6205, 449, 34),
	(76674, 'DOFYO WAFOR', 6205, 449, 34),
	(76675, 'KOBEOSER', 6205, 449, 34),
	(76676, 'KOREM', 6205, 449, 34),
	(76677, 'MAMBESAK', 6205, 449, 34),
	(76678, 'MNUSWOR', 6205, 449, 34),
	(76679, 'NERMNU', 6205, 449, 34),
	(76680, 'ROSAYENDO', 6205, 449, 34),
	(76681, 'SARWOM', 6205, 449, 34),
	(76682, 'SAUKOBYE', 6205, 449, 34),
	(76683, 'WARBON', 6205, 449, 34),
	(76684, 'WARI', 6205, 449, 34),
	(76685, 'WAROMI', 6205, 449, 34),
	(76686, 'WARSANSAN', 6205, 449, 34),
	(76687, 'YOBDI', 6205, 449, 34),
	(76688, 'DOUBO', 6206, 449, 34),
	(76689, 'SANSUNDI', 6206, 449, 34),
	(76690, 'SYURDORI', 6206, 449, 34),
	(76691, 'WANDOS', 6206, 449, 34),
	(76692, 'WOPES', 6206, 449, 34),
	(76693, 'AMBERPAREM (AMPEREM)', 6207, 449, 34),
	(76694, 'ARIMI JAYA', 6207, 449, 34),
	(76695, 'BRUYADORI', 6207, 449, 34),
	(76696, 'DAFI', 6207, 449, 34),
	(76697, 'DUAI', 6207, 449, 34),
	(76698, 'KAMUMI', 6207, 449, 34),
	(76699, 'MANDORI', 6207, 449, 34),
	(76700, 'MANDORI SUP', 6207, 449, 34),
	(76701, 'SANDOU (SANDAU)', 6207, 449, 34),
	(76702, 'WARBUKOR', 6207, 449, 34),
	(76703, 'BARUKI', 6208, 449, 34),
	(76704, 'KAMERI', 6208, 449, 34),
	(76705, 'KANSAI', 6208, 449, 34),
	(76706, 'MAMBODO SAWAI', 6208, 449, 34),
	(76707, 'NAMBER', 6208, 449, 34),
	(76708, 'POMDORI SUP', 6208, 449, 34),
	(76709, 'PONDORI (POMDORI)', 6208, 449, 34),
	(76710, 'RIMBARAYA', 6208, 449, 34),
	(76711, 'SEMAYEN', 6208, 449, 34),
	(76712, 'SERBIN', 6208, 449, 34),
	(76713, 'SUBANGUNGSI (SUBMANGGUNSI)', 6208, 449, 34),
	(76714, 'WARIDO', 6208, 449, 34),
	(76715, 'ASAIBORI', 6209, 449, 34),
	(76716, 'ASARYENDI', 6209, 449, 34),
	(76717, 'BARKORI', 6209, 449, 34),
	(76718, 'INDAIBORI', 6209, 449, 34),
	(76719, 'KORNASOREN', 6209, 449, 34),
	(76720, 'PYEFURI', 6209, 449, 34),
	(76721, 'RARSIBO', 6209, 449, 34),
	(76722, 'YENBURWO', 6209, 449, 34),
	(76723, 'YENMARU', 6209, 449, 34),
	(76724, 'ANGGADUBER', 6210, 449, 34),
	(76725, 'ANGGOPI', 6210, 449, 34),
	(76726, 'ANIMI', 6210, 449, 34),
	(76727, 'BAKRIBO', 6210, 449, 34),
	(76728, 'KAKUR', 6210, 449, 34),
	(76729, 'MAKMAKERBO', 6210, 449, 34),
	(76730, 'MARAO', 6210, 449, 34),
	(76731, 'OPIAREF', 6210, 449, 34),
	(76732, 'SAURI', 6210, 449, 34),
	(76733, 'SAWA', 6210, 449, 34),
	(76734, 'SAWADORI', 6210, 449, 34),
	(76735, 'TANJUNG BARARI', 6210, 449, 34),
	(76736, 'WADIBU', 6210, 449, 34),
	(76737, 'YENSAMA', 6210, 449, 34),
	(76738, 'MANWOR INDAH', 6211, 449, 34),
	(76739, 'MASYARA', 6211, 449, 34),
	(76740, 'PAKREKI', 6211, 449, 34),
	(76741, 'RAWAR', 6211, 449, 34),
	(76742, 'SARIBI', 6211, 449, 34),
	(76743, 'SUPMANDER', 6211, 449, 34),
	(76744, 'WANSRA', 6211, 449, 34),
	(76745, 'YENBEBA', 6211, 449, 34),
	(76746, 'YENBEBON (YENBEPON)', 6211, 449, 34),
	(76747, 'AUKI', 6212, 449, 34),
	(76748, 'INYEBOMI (INBEYOMI)', 6212, 449, 34),
	(76749, 'KANAI', 6212, 449, 34),
	(76750, 'NUSI', 6212, 449, 34),
	(76751, 'NUSI BABARUK', 6212, 449, 34),
	(76752, 'PAI', 6212, 449, 34),
	(76753, 'PAIDORI', 6212, 449, 34),
	(76754, 'SANDIDORI (SANDEDORI)', 6212, 449, 34),
	(76755, 'SOKANI', 6212, 449, 34),
	(76756, 'SORINA', 6212, 449, 34),
	(76757, 'WUNDI', 6212, 449, 34),
	(76758, 'ANDEI', 6213, 449, 34),
	(76759, 'ANDEI SUP', 6213, 449, 34),
	(76760, 'BAWEI', 6213, 449, 34),
	(76761, 'KORYAKAM', 6213, 449, 34),
	(76762, 'MANGGARI', 6213, 449, 34),
	(76763, 'SAUNBRI (SAURIBRU)', 6213, 449, 34),
	(76764, 'SAURI', 6213, 449, 34),
	(76765, 'SERDORI', 6213, 449, 34),
	(76766, 'SYORIBO', 6213, 449, 34),
	(76767, 'ADADINASNOSEN (ADAINASNOSEN)', 6214, 449, 34),
	(76768, 'ANJAREW', 6214, 449, 34),
	(76769, 'ANJEREUW', 6214, 449, 34),
	(76770, 'BRAMBAKEN', 6214, 449, 34),
	(76771, 'DARMOPIS', 6214, 449, 34),
	(76772, 'KAMORFUAR', 6214, 449, 34),
	(76773, 'KARANG MULIA', 6214, 449, 34),
	(76774, 'KINMOM', 6214, 449, 34),
	(76775, 'MANDOUW', 6214, 449, 34),
	(76776, 'MANSINYAS', 6214, 449, 34),
	(76777, 'MARYENDI', 6214, 449, 34),
	(76778, 'SAMBAWOFUAR', 6214, 449, 34),
	(76779, 'SAMOFA', 6214, 449, 34),
	(76780, 'SNERBO', 6214, 449, 34),
	(76781, 'SUMBERKER', 6214, 449, 34),
	(76782, 'WISATA BINSARI', 6214, 449, 34),
	(76783, 'YAFDAS', 6214, 449, 34),
	(76784, 'YAFDAS', 6214, 449, 34),
	(76785, 'ANDOINA', 6215, 449, 34),
	(76786, 'BUSDORI', 6215, 449, 34),
	(76787, 'FARUSI', 6215, 449, 34),
	(76788, 'INSUSBARI', 6215, 449, 34),
	(76789, 'MANDENDERI', 6215, 449, 34),
	(76790, 'MARDORI', 6215, 449, 34),
	(76791, 'NAPDORI', 6215, 449, 34),
	(76792, 'ORKDORI', 6215, 449, 34),
	(76793, 'RAMDORI', 6215, 449, 34),
	(76794, 'SARWA', 6215, 449, 34),
	(76795, 'SASWARBO', 6215, 449, 34),
	(76796, 'SWAINOBER', 6215, 449, 34),
	(76797, 'SWAIPAK', 6215, 449, 34),
	(76798, 'WOMBRISAUW', 6215, 449, 34),
	(76799, 'YENBEPIOPER (YEMBEPIOPER)', 6215, 449, 34),
	(76800, 'AMAN', 6216, 449, 34),
	(76801, 'AMMOY', 6216, 449, 34),
	(76802, 'BIAWER', 6216, 449, 34),
	(76803, 'DIANO', 6216, 449, 34),
	(76804, 'IMBARI', 6216, 449, 34),
	(76805, 'INSWANBESI', 6216, 449, 34),
	(76806, 'INSWANBESI SUP', 6216, 449, 34),
	(76807, 'INYOBI', 6216, 449, 34),
	(76808, 'KARUIBERIK', 6216, 449, 34),
	(76809, 'KOMBOY', 6216, 449, 34),
	(76810, 'KOYOMI', 6216, 449, 34),
	(76811, 'MAMFIAS', 6216, 449, 34),
	(76812, 'MANBEORI', 6216, 449, 34),
	(76813, 'MANIRI', 6216, 449, 34),
	(76814, 'MARUR', 6216, 449, 34),
	(76815, 'SAWAY', 6216, 449, 34),
	(76816, 'WARAWAF', 6216, 449, 34),
	(76817, 'WASANI', 6216, 449, 34),
	(76818, 'WIR INSOS', 6216, 449, 34),
	(76819, 'YERUBOY', 6216, 449, 34),
	(76820, 'ASUR', 6217, 449, 34),
	(76821, 'BOSNABRAIDI', 6217, 449, 34),
	(76822, 'INDAWI', 6217, 449, 34),
	(76823, 'KARMON', 6217, 449, 34),
	(76824, 'MADIRAI', 6217, 449, 34),
	(76825, 'SOOR (SOR)', 6217, 449, 34),
	(76826, 'WASORI', 6217, 449, 34),
	(76827, 'YAWOSI (FANINDI)', 6217, 449, 34),
	(76828, 'ADOKI', 6218, 449, 34),
	(76829, 'AMYABENRAM', 6218, 449, 34),
	(76830, 'BINYERI', 6218, 449, 34),
	(76831, 'BIRUB', 6218, 449, 34),
	(76832, 'INPENDI', 6218, 449, 34),
	(76833, 'KABIDON', 6218, 449, 34),
	(76834, 'MOIBAKEN', 6218, 449, 34),
	(76835, 'PADWA', 6218, 449, 34),
	(76836, 'PADWA PANTAI', 6218, 449, 34),
	(76837, 'RARMPIMBO (RAMPIMBO)', 6218, 449, 34),
	(76838, 'SAMBER', 6218, 449, 34),
	(76839, 'SAMBER SUP', 6218, 449, 34),
	(76840, 'SUNERI', 6218, 449, 34),
	(76841, 'SUNYAR', 6218, 449, 34),
	(76842, 'SYABES', 6218, 449, 34),
	(76843, 'URFU', 6218, 449, 34),
	(76844, 'WAROI (WAROY)', 6218, 449, 34),
	(76845, 'WIRMAKER', 6218, 449, 34),
	(76846, 'YENDIDORI', 6218, 449, 34),
	(76847, 'ANYUMKA', 6219, 450, 34),
	(76848, 'ARIMBIT', 6219, 450, 34),
	(76849, 'AWAKEN (NYUM)', 6219, 450, 34),
	(76850, 'KLOFKAMP (KOLOPKAM)', 6219, 450, 34),
	(76851, 'KUKEN', 6219, 450, 34),
	(76852, 'ARIMBET', 6220, 450, 34),
	(76853, 'AROA', 6220, 450, 34),
	(76854, 'BUKIT', 6220, 450, 34),
	(76855, 'GINGGIMOP (GINGGIMOB)', 6220, 450, 34),
	(76856, 'MAJU', 6220, 450, 34),
	(76857, 'PATRIOT', 6220, 450, 34),
	(76858, 'UJUNG', 6220, 450, 34),
	(76859, 'AIFO', 6221, 450, 34),
	(76860, 'BOMAKIA DUA (II)', 6221, 450, 34),
	(76861, 'BOMAKIA SATU (I)', 6221, 450, 34),
	(76862, 'SOMI', 6221, 450, 34),
	(76863, 'UNI', 6221, 450, 34),
	(76864, 'FIRIWAGE', 6222, 450, 34),
	(76865, 'KABUWAGE', 6222, 450, 34),
	(76866, 'KARUWAGE', 6222, 450, 34),
	(76867, 'WALIBURU', 6222, 450, 34),
	(76868, 'BANGUN', 6223, 450, 34),
	(76869, 'DOMO', 6223, 450, 34),
	(76870, 'HAMKHU', 6223, 450, 34),
	(76871, 'HELLO', 6223, 450, 34),
	(76872, 'MAKMUR', 6223, 450, 34),
	(76873, 'NAVINI', 6223, 450, 34),
	(76874, 'SADAR', 6223, 450, 34),
	(76875, 'SOHOKANGGO', 6223, 450, 34),
	(76876, 'AUTRIOP', 6224, 450, 34),
	(76877, 'LANGGOAN', 6224, 450, 34),
	(76878, 'OGENETAN (OGENATHAN)', 6224, 450, 34),
	(76879, 'TETOP', 6224, 450, 34),
	(76880, 'WARIKTOP (WARIKTOOP)', 6224, 450, 34),
	(76881, 'ANGGAI', 6225, 450, 34),
	(76882, 'ASIKI', 6225, 450, 34),
	(76883, 'BUTIPTIRI', 6225, 450, 34),
	(76884, 'GETENTIRI', 6225, 450, 34),
	(76885, 'KAPOGU', 6225, 450, 34),
	(76886, 'MIRI', 6225, 450, 34),
	(76887, 'BIWAGE', 6226, 450, 34),
	(76888, 'BIWAGE DUA', 6226, 450, 34),
	(76889, 'KAWAGIT', 6226, 450, 34),
	(76890, 'KOMBAI', 6226, 450, 34),
	(76891, 'NIOP', 6226, 450, 34),
	(76892, 'WANGGOM', 6226, 450, 34),
	(76893, 'METTO', 6227, 450, 34),
	(76894, 'OBINANGGE', 6227, 450, 34),
	(76895, 'UJUNGKIA', 6227, 450, 34),
	(76896, 'WATEMU', 6227, 450, 34),
	(76897, 'DEMA', 6228, 450, 34),
	(76898, 'SINIMBURU', 6228, 450, 34),
	(76899, 'UGO', 6228, 450, 34),
	(76900, 'WANGGEMALO', 6228, 450, 34),
	(76901, 'YAFUFLA', 6228, 450, 34),
	(76902, 'AMUAN', 6229, 450, 34),
	(76903, 'AWAYANKA', 6229, 450, 34),
	(76904, 'KAWANGTET', 6229, 450, 34),
	(76905, 'MOKBIRAN', 6229, 450, 34),
	(76906, 'JAIR', 6230, 450, 34),
	(76907, 'KOUH', 6230, 450, 34),
	(76908, 'MANDOBO', 6230, 450, 34),
	(76909, 'AMPERA', 6231, 450, 34),
	(76910, 'MARIAM', 6231, 450, 34),
	(76911, 'MAWAN', 6231, 450, 34),
	(76912, 'PERSATUAN', 6231, 450, 34),
	(76913, 'SOHOKANGGO (SOKANGGO)', 6231, 450, 34),
	(76914, 'BAYANGGOP', 6232, 450, 34),
	(76915, 'BURUNGGOP', 6232, 450, 34),
	(76916, 'GAGUOP', 6232, 450, 34),
	(76917, 'KEWAM', 6232, 450, 34),
	(76918, 'MANGGA TIGA', 6232, 450, 34),
	(76919, 'MANGGELUM', 6232, 450, 34),
	(76920, 'ANDOPBIT', 6233, 450, 34),
	(76921, 'ANGGUMBIT', 6233, 450, 34),
	(76922, 'AWAYANKA', 6233, 450, 34),
	(76923, 'EPSEMBIT', 6233, 450, 34),
	(76924, 'IMKO', 6233, 450, 34),
	(76925, 'KAKUNA', 6233, 450, 34),
	(76926, 'KAMKA', 6233, 450, 34),
	(76927, 'MINDIPTANA', 6233, 450, 34),
	(76928, 'NIYIMBANG', 6233, 450, 34),
	(76929, 'OSSO', 6233, 450, 34),
	(76930, 'TINGGAM', 6233, 450, 34),
	(76931, 'UMAP', 6233, 450, 34),
	(76932, 'WANGGAT KIBI', 6233, 450, 34),
	(76933, 'KAWAKTEMBUT', 6234, 450, 34),
	(76934, 'NINATI', 6234, 450, 34),
	(76935, 'TEMBUTKA', 6234, 450, 34),
	(76936, 'TIMKA', 6234, 450, 34),
	(76937, 'UPYETETKO', 6234, 450, 34),
	(76938, 'AMBORAN', 6235, 450, 34),
	(76939, 'ANGGAMBURAN', 6235, 450, 34),
	(76940, 'KANGGUP', 6235, 450, 34),
	(76941, 'SESNUK', 6235, 450, 34),
	(76942, 'YOMKONDO', 6235, 450, 34),
	(76943, 'AIWAT', 6236, 450, 34),
	(76944, 'KAISA (KAISAH)', 6236, 450, 34),
	(76945, 'SUBUR', 6236, 450, 34),
	(76946, 'WAGHAI (BHAKTI)', 6236, 450, 34),
	(76947, 'IKCAN', 6237, 450, 34),
	(76948, 'INGGEMBIT (ENGGEMBIT)', 6237, 450, 34),
	(76949, 'JETETKUN / YETETKUN', 6237, 450, 34),
	(76950, 'KANGGEWOT', 6237, 450, 34),
	(76951, 'UPKIM', 6237, 450, 34),
	(76952, 'WAMETKAPA', 6237, 450, 34),
	(76953, 'WAROPKO', 6237, 450, 34),
	(76954, 'WINIKTIT', 6237, 450, 34),
	(76955, 'WOMBON', 6237, 450, 34),
	(76956, 'FEFERO (FEFEKO)', 6238, 450, 34),
	(76957, 'MANGGEMAKHE (MANGGEMAHE)', 6238, 450, 34),
	(76958, 'YANIRUMA', 6238, 450, 34),
	(76959, 'KOPAI DUA (II)', 6239, 451, 34),
	(76960, 'KOPAI SATU (I)', 6239, 451, 34),
	(76961, 'WOGE', 6239, 451, 34),
	(76962, 'KAPIRAYA (KOMAUTO)', 6240, 451, 34),
	(76963, 'AMAGO', 6241, 451, 34),
	(76964, 'ATOUDA', 6241, 451, 34),
	(76965, 'BOBAI', 6241, 451, 34),
	(76966, 'BOMOU DUA (BOMAU II)', 6241, 451, 34),
	(76967, 'BOMOU III', 6241, 451, 34),
	(76968, 'BOMOU SATU (I)', 6241, 451, 34),
	(76969, 'BUWOUDIMI', 6241, 451, 34),
	(76970, 'AYATE (AYATEI)', 6242, 451, 34),
	(76971, 'DEMAGO (DEEMAGO)', 6242, 451, 34),
	(76972, 'DIGIBAGATA', 6242, 451, 34),
	(76973, 'DIYAI', 6242, 451, 34),
	(76974, 'GAKOKEBO', 6242, 451, 34),
	(76975, 'ONAGO', 6242, 451, 34),
	(76976, 'PIYAKADIMI (PIYEKEDIMI)', 6242, 451, 34),
	(76977, 'TENEDAGI', 6242, 451, 34),
	(76978, 'WAGOMANI', 6242, 451, 34),
	(76979, 'WIDIMEI (WIDI MEY)', 6242, 451, 34),
	(76980, 'WIDUWAKIYA (WIDUWAKIA)', 6242, 451, 34),
	(76981, 'YENUDOBA/JINIDOBA', 6242, 451, 34),
	(76982, 'BAGOU (BAGAU)', 6243, 451, 34),
	(76983, 'BAGUMOMA', 6243, 451, 34),
	(76984, 'BOWOUBADO', 6243, 451, 34),
	(76985, 'DAKEBO', 6243, 451, 34),
	(76986, 'DAMABAGATA', 6243, 451, 34),
	(76987, 'EDAGOTADI', 6243, 451, 34),
	(76988, 'EDAIYODAGI', 6243, 451, 34),
	(76989, 'BOBUBUTU', 6244, 452, 34),
	(76990, 'DENEMANI', 6244, 452, 34),
	(76991, 'DOGIMANI (ABGOIGI)', 6244, 452, 34),
	(76992, 'EGEBUTU', 6244, 452, 34),
	(76993, 'IDADAGI', 6244, 452, 34),
	(76994, 'KEGEMANI (KIGAMANI)', 6244, 452, 34),
	(76995, 'MAKIDIMI (MAKIDINI)', 6244, 452, 34),
	(76996, 'MOTITO', 6244, 452, 34),
	(76997, 'PONA (SUKAMAJU)', 6244, 452, 34),
	(76998, 'BUKAPA (BUGAPA)', 6245, 452, 34),
	(76999, 'DIKIYOUWA/DIKIYOUMA (TOKAPO)', 6245, 452, 34)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 76")
}
