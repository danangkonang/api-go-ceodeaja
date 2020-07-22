package kel70

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel77() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(77000, 'EKEMANIDA (EKEMENIDA)', 6245, 452, 34),
	(77001, 'IDAKOTU', 6245, 452, 34),
	(77002, 'IKEBO (MOANEMANI)', 6245, 452, 34),
	(77003, 'KIMIPUGI (KIMUPUGI)', 6245, 452, 34),
	(77004, 'MAUWA', 6245, 452, 34),
	(77005, 'PUTAPA (UPUTAPA)', 6245, 452, 34),
	(77006, 'BOGIYATEUGI', 6246, 452, 34),
	(77007, 'BOTUMUMA (BOTUMOMA)', 6246, 452, 34),
	(77008, 'DIGIKEBO', 6246, 452, 34),
	(77009, 'MATADI', 6246, 452, 34),
	(77010, 'OBAIBEGA', 6246, 452, 34),
	(77011, 'POUWODA (POUWUODA)', 6246, 452, 34),
	(77012, 'PUWETA/PUETA I', 6246, 452, 34),
	(77013, 'PUWETA/PUETA II', 6246, 452, 34),
	(77014, 'TUWAIDA', 6246, 452, 34),
	(77015, 'UGIKAGOUDA', 6246, 452, 34),
	(77016, 'UGIKEBO', 6246, 452, 34),
	(77017, 'YEPO (MAKIDIMI)', 6246, 452, 34),
	(77018, 'BODUDA', 6247, 452, 34),
	(77019, 'BOKAIBUTU (BUKAIBUTO)', 6247, 452, 34),
	(77020, 'BUNAUWO', 6247, 452, 34),
	(77021, 'DEIYAPA (DEIYEPA)', 6247, 452, 34),
	(77022, 'NUWA (NUA)', 6247, 452, 34),
	(77023, 'UGAPUGA', 6247, 452, 34),
	(77024, 'YOTAPUGA', 6247, 452, 34),
	(77025, 'DUNTEK (DUMTEK)', 6248, 452, 34),
	(77026, 'EKIMANI (EKIWANI)', 6248, 452, 34),
	(77027, 'IDAKEBO', 6248, 452, 34),
	(77028, 'IKRAR', 6248, 452, 34),
	(77029, 'KOYAKAGO (KUYAKAGO)', 6248, 452, 34),
	(77030, 'MOGOU', 6248, 452, 34),
	(77031, 'OBAYO', 6248, 452, 34),
	(77032, 'PUGATADI DUA (II)', 6248, 452, 34),
	(77033, 'PUGATADI SATU (I)', 6248, 452, 34),
	(77034, 'YAWETADI', 6248, 452, 34),
	(77035, 'ABAIMAIDA', 6249, 452, 34),
	(77036, 'BOMOMANI', 6249, 452, 34),
	(77037, 'DAWAIKUNU', 6249, 452, 34),
	(77038, 'DIYOUDIMI', 6249, 452, 34),
	(77039, 'GAPOYA', 6249, 452, 34),
	(77040, 'MAGODE (MAGONE)', 6249, 452, 34),
	(77041, 'OBAIKAGOPA (ABAUGI)', 6249, 452, 34),
	(77042, 'ABOUYAGA', 6250, 452, 34),
	(77043, 'MAIKOTU', 6250, 452, 34),
	(77044, 'TOUBAIKEBO (TOBAIKEBO)', 6250, 452, 34),
	(77045, 'YEGOUKOTU', 6250, 452, 34),
	(77046, 'ADAUWO', 6251, 452, 34),
	(77047, 'ATOU (TUAMANI)', 6251, 452, 34),
	(77048, 'DIYEUGI', 6251, 452, 34),
	(77049, 'GABAIKUNU', 6251, 452, 34),
	(77050, 'MEGAIKEBO', 6251, 452, 34),
	(77051, 'MODIO', 6251, 452, 34),
	(77052, 'PIYAKUNU', 6251, 452, 34),
	(77053, 'PUTAPA', 6251, 452, 34),
	(77054, 'TIMEPA', 6251, 452, 34),
	(77055, 'UPIBEGA (UKUDAWATA)', 6251, 452, 34),
	(77056, 'APOGOMAKIDA (APOROMAKIDA)', 6252, 452, 34),
	(77057, 'DENIYODE (DENEIODE)', 6252, 452, 34),
	(77058, 'EGIPA', 6252, 452, 34),
	(77059, 'IDEDUWA', 6252, 452, 34),
	(77060, 'KEGATA', 6252, 452, 34),
	(77061, 'TIPAUGI', 6252, 452, 34),
	(77062, 'UKAGU', 6252, 452, 34),
	(77063, 'YEGIYEPA (YEGEIYEPA)', 6252, 452, 34),
	(77064, 'IYARO', 6253, 452, 34),
	(77065, 'SUKIKAI', 6253, 452, 34),
	(77066, 'UNITO', 6253, 452, 34),
	(77067, 'WIGOUMAKIDA', 6253, 452, 34),
	(77068, 'AGISIGA', 6254, 453, 34),
	(77069, 'BALAMAI (BALAIMAI)', 6254, 453, 34),
	(77070, 'BAMOGO (MBAMOGO)', 6254, 453, 34),
	(77071, 'DANGKOA', 6254, 453, 34),
	(77072, 'DAPIAGA', 6254, 453, 34),
	(77073, 'FUGISIGA (PUGISIGA)', 6254, 453, 34),
	(77074, 'JAWASIGA', 6254, 453, 34),
	(77075, 'KEMBAGEL', 6254, 453, 34),
	(77076, 'KOMBOGASIGA', 6254, 453, 34),
	(77077, 'SOLIP', 6254, 453, 34),
	(77078, 'TAUSIGA (TOUSIGA)', 6254, 453, 34),
	(77079, 'TOMASIGA', 6254, 453, 34),
	(77080, 'UNABINDAGA (UNABUNDAGA/UNABUNDOGA)', 6254, 453, 34),
	(77081, 'ANEYA', 6255, 453, 34),
	(77082, 'BIANDOGA', 6255, 453, 34),
	(77083, 'BIATAPA', 6255, 453, 34),
	(77084, 'BUGALAGA', 6255, 453, 34),
	(77085, 'DANGGATADI', 6255, 453, 34),
	(77086, 'DEBASIGA', 6255, 453, 34),
	(77087, 'DEBASIGA DUA / II', 6255, 453, 34),
	(77088, 'HIPADIPA (HITADIPA)', 6255, 453, 34),
	(77089, 'ISAN DOGA', 6255, 453, 34),
	(77090, 'KALAWA', 6255, 453, 34),
	(77091, 'KIGITADI', 6255, 453, 34),
	(77092, 'MBIALAPA', 6255, 453, 34),
	(77093, 'YAGAITO (JAGAITO/TAGITO)', 6255, 453, 34),
	(77094, 'YANEI (JANEI/YAMEI)', 6255, 453, 34),
	(77095, 'DANGGOMBA', 6256, 453, 34),
	(77096, 'EKNEMBA/ELENEMBA', 6256, 453, 34),
	(77097, 'HITADIPA', 6256, 453, 34),
	(77098, 'KULAPA', 6256, 453, 34),
	(77099, 'SANABA (SANAMBA)', 6256, 453, 34),
	(77100, 'SOAGAMA', 6256, 453, 34),
	(77101, 'TITIGI', 6256, 453, 34),
	(77102, 'WABUI', 6256, 453, 34),
	(77103, 'AGAPA', 6257, 453, 34),
	(77104, 'BILAI DUA (II)', 6257, 453, 34),
	(77105, 'BILAI SATU (I)', 6257, 453, 34),
	(77106, 'BONOGO', 6257, 453, 34),
	(77107, 'BUBISIGA (BUGUSIGA)', 6257, 453, 34),
	(77108, 'DEGESIGA (DEGESISE)', 6257, 453, 34),
	(77109, 'ENGGANEGA', 6257, 453, 34),
	(77110, 'HIYABU', 6257, 453, 34),
	(77111, 'HUGITAPA / HUBITAPA', 6257, 453, 34),
	(77112, 'KENETAPA (KENDETAPA/KENEDEPA)', 6257, 453, 34),
	(77113, 'KOBAE (KOBAI)', 6257, 453, 34),
	(77114, 'MAPA', 6257, 453, 34),
	(77115, 'MAYA', 6257, 453, 34),
	(77116, 'OGEAPA / OGEAYAPA', 6257, 453, 34),
	(77117, 'POGAPA', 6257, 453, 34),
	(77118, 'SELEMANA (SELEMAMA)', 6257, 453, 34),
	(77119, 'WAIAGEPA / WAIGEPA', 6257, 453, 34),
	(77120, 'ZOMBAN DOGA (ZOMBANDOGA)', 6257, 453, 34),
	(77121, 'BILOGAI', 6258, 453, 34),
	(77122, 'DEGEYABU', 6258, 453, 34),
	(77123, 'EKNEMBA / ENEMBA', 6258, 453, 34),
	(77124, 'EMONDI (KEMONDI)', 6258, 453, 34),
	(77125, 'KUMLAGUPA', 6258, 453, 34),
	(77126, 'MAMBA', 6258, 453, 34),
	(77127, 'MBILUSIGA', 6258, 453, 34),
	(77128, 'MINDAU', 6258, 453, 34),
	(77129, 'NDUGUSIGA', 6258, 453, 34),
	(77130, 'PUYAGIYA', 6258, 453, 34),
	(77131, 'TITIGI', 6258, 453, 34),
	(77132, 'UGIMBA', 6258, 453, 34),
	(77133, 'WANDOGA', 6258, 453, 34),
	(77134, 'YALAI (JABAI/JALAI)', 6258, 453, 34),
	(77135, 'YOKATAPA (TOKATAPA)', 6258, 453, 34),
	(77136, 'YOPARU', 6258, 453, 34),
	(77137, 'DEBASIGA DUA PEMDA', 6259, 453, 34),
	(77138, 'DEBASIGA PEMDA', 6259, 453, 34),
	(77139, 'DUBASIGA PEMDA', 6259, 453, 34),
	(77140, 'HULAPUGA', 6259, 453, 34),
	(77141, 'ISANDOGA PEMDA', 6259, 453, 34),
	(77142, 'JAE PEMDA', 6259, 453, 34),
	(77143, 'MBUGULO PEMDA', 6259, 453, 34),
	(77144, 'MOGALO PEMDA', 6259, 453, 34),
	(77145, 'SABISA', 6259, 453, 34),
	(77146, 'ABEPANTAI', 6260, 454, 34),
	(77147, 'ASANO', 6260, 454, 34),
	(77148, 'AWIYO', 6260, 454, 34),
	(77149, 'ENGROS (ENGGROS)', 6260, 454, 34),
	(77150, 'KOTA BARU', 6260, 454, 34),
	(77151, 'KOYA KOSO', 6260, 454, 34),
	(77152, 'NAFRI', 6260, 454, 34),
	(77153, 'WAENA', 6260, 454, 34),
	(77154, 'YOBE', 6260, 454, 34),
	(77155, 'AURINA', 6261, 454, 34),
	(77156, 'HULU ATAS', 6261, 454, 34),
	(77157, 'KAMIKARU', 6261, 454, 34),
	(77158, 'MUARA NAWA (MUARAH MAWA)', 6261, 454, 34),
	(77159, 'NAIRA', 6261, 454, 34),
	(77160, 'PAGAI', 6261, 454, 34),
	(77161, 'AMBORA', 6262, 454, 34),
	(77162, 'DEMTA KOTA', 6262, 454, 34),
	(77163, 'KAMDERA', 6262, 454, 34),
	(77164, 'MUAIF', 6262, 454, 34),
	(77165, 'MURIS KECIL', 6262, 454, 34),
	(77166, 'YAKORE', 6262, 454, 34),
	(77167, 'YAUGAPSA (YOUGAPSA)', 6262, 454, 34),
	(77168, 'DOROMENA', 6263, 454, 34),
	(77169, 'ENTIYEBO', 6263, 454, 34),
	(77170, 'KENDATE', 6263, 454, 34),
	(77171, 'TABLASUPA', 6263, 454, 34),
	(77172, 'WAIYA', 6263, 454, 34),
	(77173, 'WAMBENA', 6263, 454, 34),
	(77174, 'YEFASE (YEPASE)', 6263, 454, 34),
	(77175, 'YEWENA', 6263, 454, 34),
	(77176, 'ABAR (ATABAR)', 6264, 454, 34),
	(77177, 'BABRONGKO (IFAR BOBRONGKO)', 6264, 454, 34),
	(77178, 'EBUNGFA', 6264, 454, 34),
	(77179, 'KHAMEYAKA (KAMEYOKA/KHAMAEKA)', 6264, 454, 34),
	(77180, 'SIMPORO (BABO/YOSIBA/HOMF)', 6264, 454, 34),
	(77181, 'BANGAI', 6265, 454, 34),
	(77182, 'IWON', 6265, 454, 34),
	(77183, 'KLAISU', 6265, 454, 34),
	(77184, 'OMON', 6265, 454, 34),
	(77185, 'HEDAM', 6266, 454, 34),
	(77186, 'WAENA', 6266, 454, 34),
	(77187, 'YABANSAI', 6266, 454, 34),
	(77188, 'YOKA', 6266, 454, 34),
	(77189, 'ARDIPURA', 6267, 454, 34),
	(77190, 'ARGAPURA', 6267, 454, 34),
	(77191, 'ENTROP', 6267, 454, 34),
	(77192, 'HAMADI', 6267, 454, 34),
	(77193, 'NUMBAI (NUMBAY)', 6267, 454, 34),
	(77194, 'TAHIMA SORAMA', 6267, 454, 34),
	(77195, 'TOBATI', 6267, 454, 34),
	(77196, 'VIM', 6267, 454, 34),
	(77197, 'WAHNO', 6267, 454, 34),
	(77198, 'WAY MHOROCK', 6267, 454, 34),
	(77199, 'ANGKASAPURA', 6268, 454, 34),
	(77200, 'BHAYANGKARA (BAYANGKARA)', 6268, 454, 34),
	(77201, 'GURABESI', 6268, 454, 34),
	(77202, 'IMBI', 6268, 454, 34),
	(77203, 'MANDALA', 6268, 454, 34),
	(77204, 'TANJUNG RIA', 6268, 454, 34),
	(77205, 'TRIKORA', 6268, 454, 34),
	(77206, 'LAPUA', 6269, 454, 34),
	(77207, 'SEBUM', 6269, 454, 34),
	(77208, 'SOSKOTEK', 6269, 454, 34),
	(77209, 'UMBRON', 6269, 454, 34),
	(77210, 'YADOUW (YADAUW)', 6269, 454, 34),
	(77211, 'AIB', 6270, 454, 34),
	(77212, 'BENGGUIN PROGO', 6270, 454, 34),
	(77213, 'KWANSU', 6270, 454, 34),
	(77214, 'MAMDA', 6270, 454, 34),
	(77215, 'MANDAYAWAN', 6270, 454, 34),
	(77216, 'NAMBON', 6270, 454, 34),
	(77217, 'NAMEI', 6270, 454, 34),
	(77218, 'SABEAB KECIL', 6270, 454, 34),
	(77219, 'SAMA', 6270, 454, 34),
	(77220, 'SEKORI', 6270, 454, 34),
	(77221, 'SKOAIM', 6270, 454, 34),
	(77222, 'SOAIB', 6270, 454, 34),
	(77223, 'BRASO', 6271, 454, 34),
	(77224, 'BRING', 6271, 454, 34),
	(77225, 'DAMOI KATI', 6271, 454, 34),
	(77226, 'DEMETIN', 6271, 454, 34),
	(77227, 'HATIB', 6271, 454, 34),
	(77228, 'IBUB', 6271, 454, 34),
	(77229, 'JAGRANG', 6271, 454, 34),
	(77230, 'NEMBU GRESI', 6271, 454, 34),
	(77231, 'OMON', 6271, 454, 34),
	(77232, 'PUPEHABU', 6271, 454, 34),
	(77233, 'SWENTAB', 6271, 454, 34),
	(77234, 'YANBRA (YANIM)', 6271, 454, 34),
	(77235, 'HOLTEKAMP', 6272, 454, 34),
	(77236, 'KOYA BARAT', 6272, 454, 34),
	(77237, 'KOYA TENGAH', 6272, 454, 34),
	(77238, 'KOYA TIMUR', 6272, 454, 34),
	(77239, 'MOSSO', 6272, 454, 34),
	(77240, 'SKOW MABO (SKOUW MABO )', 6272, 454, 34),
	(77241, 'SKOW SAE (SKOUW SAE)', 6272, 454, 34),
	(77242, 'SKOW YAMBE (SKOUW YAMBE)', 6272, 454, 34),
	(77243, 'BESUM', 6273, 454, 34),
	(77244, 'HANGGAIY HAMONG', 6273, 454, 34),
	(77245, 'IMESTUM (IMUSTUM)', 6273, 454, 34),
	(77246, 'KARYA BUMI', 6273, 454, 34),
	(77247, 'SANGGAI', 6273, 454, 34),
	(77248, 'SARMAI ATAS', 6273, 454, 34),
	(77249, 'SARMAI BAWAH', 6273, 454, 34),
	(77250, 'SUMBE', 6273, 454, 34),
	(77251, 'YAKASIB (YAKASIP/YOKASIB)', 6273, 454, 34),
	(77252, 'BENYOM JAYA I', 6274, 454, 34),
	(77253, 'BENYOM JAYA II', 6274, 454, 34),
	(77254, 'BERAB', 6274, 454, 34),
	(77255, 'BUNYOM', 6274, 454, 34),
	(77256, 'HAMONGKRANG', 6274, 454, 34),
	(77257, 'NEMBUKRANG (NIMBOKRANG)', 6274, 454, 34),
	(77258, 'NEMBUKRANG SARI (NIMBOKRANG SARI)', 6274, 454, 34),
	(77259, 'RHEPANG MUAIF/MUAIB', 6274, 454, 34),
	(77260, 'WAHAB', 6274, 454, 34),
	(77261, 'BENYOM', 6275, 454, 34),
	(77262, 'GEMEBS', 6275, 454, 34),
	(77263, 'IMSAR', 6275, 454, 34),
	(77264, 'KAITEMUNG (KAUTEMUNG)', 6275, 454, 34),
	(77265, 'KUIPON (KUIPONS)', 6275, 454, 34),
	(77266, 'KUWASE', 6275, 454, 34),
	(77267, 'MEYU', 6275, 454, 34),
	(77268, 'OYENGSI', 6275, 454, 34),
	(77269, 'POBAIM', 6275, 454, 34),
	(77270, 'SINGGRI', 6275, 454, 34),
	(77271, 'SINGGRIWAY (SINGGRIWAI)', 6275, 454, 34),
	(77272, 'TABRI', 6275, 454, 34),
	(77273, 'YENGGU (YENGGU LAMA)', 6275, 454, 34),
	(77274, 'YENGGU BARU', 6275, 454, 34),
	(77275, 'NEHIBE (NACHA TAWA)', 6276, 454, 34),
	(77276, 'ORMUWARI (NEWA)', 6276, 454, 34),
	(77277, 'YONGSU BESAR (DOSOYO)', 6276, 454, 34),
	(77278, 'YONGSU SAPARI', 6276, 454, 34),
	(77279, 'DOBONSOLO', 6277, 454, 34),
	(77280, 'HINEKOMBE', 6277, 454, 34),
	(77281, 'HOBONG', 6277, 454, 34),
	(77282, 'IFALE (ILFELE)', 6277, 454, 34),
	(77283, 'IFAR BESAR', 6277, 454, 34),
	(77284, 'KEHERAN (KEHIRAN)', 6277, 454, 34),
	(77285, 'SENTANI KOTA', 6277, 454, 34),
	(77286, 'SEREH', 6277, 454, 34),
	(77287, 'YOBEH', 6277, 454, 34),
	(77288, 'YOBOY', 6277, 454, 34),
	(77289, 'DOSAY (DOSAI)', 6278, 454, 34),
	(77290, 'MARIBU', 6278, 454, 34),
	(77291, 'SABRON SARI', 6278, 454, 34),
	(77292, 'SABRON YARU (SABRO YARU)', 6278, 454, 34),
	(77293, 'WAIBRON', 6278, 454, 34),
	(77294, 'ASEI BESAR', 6279, 454, 34),
	(77295, 'ASEI KECIL', 6279, 454, 34),
	(77296, 'AYAPO (ITAKIWA)', 6279, 454, 34),
	(77297, 'NENDALI', 6279, 454, 34),
	(77298, 'NOLOKLA', 6279, 454, 34),
	(77299, 'PUAY (PUAI)', 6279, 454, 34),
	(77300, 'YOKIWA', 6279, 454, 34),
	(77301, 'BENEIK', 6280, 454, 34),
	(77302, 'GARUSA', 6280, 454, 34),
	(77303, 'GURIYAD (GURYARD)', 6280, 454, 34),
	(77304, 'NANDALZI (NANDAIZI)', 6280, 454, 34),
	(77305, 'SAWESUMA (SAWA SUMA)', 6280, 454, 34),
	(77306, 'SENTOSA (SANTOSA)', 6280, 454, 34),
	(77307, 'BAMBAR', 6281, 454, 34),
	(77308, 'DONDAI (DONDAY)', 6281, 454, 34),
	(77309, 'DOYO BARU', 6281, 454, 34),
	(77310, 'DOYO LAMA', 6281, 454, 34),
	(77311, 'KWADEWARE (KANDA)', 6281, 454, 34),
	(77312, 'SOSIRI', 6281, 454, 34),
	(77313, 'YAKONDE', 6281, 454, 34),
	(77314, 'BUMI SAHAJA', 6282, 454, 34),
	(77315, 'BUNDRU', 6282, 454, 34),
	(77316, 'KWARJA (KWAJA)', 6282, 454, 34),
	(77317, 'NAWA MUKTI', 6282, 454, 34),
	(77318, 'NAWA MULYA', 6282, 454, 34),
	(77319, 'ONGAN JAYA', 6282, 454, 34),
	(77320, 'PURNAWAJATI (PURNAMA JATI)', 6282, 454, 34),
	(77321, 'TABBEYAN (TABEYAN)', 6282, 454, 34),
	(77322, 'TAQWA/TAKWA BANGUN', 6282, 454, 34),
	(77323, 'BUSERYO', 6283, 454, 34),
	(77324, 'ENDOKISI', 6283, 454, 34),
	(77325, 'MARUWAY', 6283, 454, 34),
	(77326, 'MEUKISI', 6283, 454, 34),
	(77327, 'SNAMAY', 6283, 454, 34),
	(77328, 'ARABODA', 6284, 455, 34),
	(77329, 'KIMBIM', 6284, 455, 34),
	(77330, 'KOMBAGWE', 6284, 455, 34),
	(77331, 'LOGOTPAGA', 6284, 455, 34),
	(77332, 'LOKI', 6284, 455, 34),
	(77333, 'MILIGATNEM', 6284, 455, 34),
	(77334, 'TIKAWO', 6284, 455, 34),
	(77335, 'WALAK', 6284, 455, 34),
	(77336, 'WAWANCA', 6284, 455, 34),
	(77337, 'ASOLOKOBAL', 6285, 455, 34),
	(77338, 'ASOTAPO', 6285, 455, 34),
	(77339, 'HELALUWA', 6285, 455, 34),
	(77340, 'HESATOM', 6285, 455, 34),
	(77341, 'MULINEKAMA', 6285, 455, 34),
	(77342, 'NINABUA', 6285, 455, 34),
	(77343, 'SINATA', 6285, 455, 34),
	(77344, 'WIAIMA', 6285, 455, 34),
	(77345, 'YAPEMA', 6285, 455, 34),
	(77346, 'AIR GARAM', 6286, 455, 34),
	(77347, 'ASOTIPO', 6286, 455, 34),
	(77348, 'HEBERIMA', 6286, 455, 34),
	(77349, 'HITIGIMA', 6286, 455, 34),
	(77350, 'HUKULIMO', 6286, 455, 34),
	(77351, 'IWIGIMA', 6286, 455, 34),
	(77352, 'KUANTAPO', 6286, 455, 34),
	(77353, 'POBIATMA', 6286, 455, 34),
	(77354, 'PUTAGEIMA (PUTAGAIMA)', 6286, 455, 34),
	(77355, 'SOGOKMO (SOGOMO)', 6286, 455, 34),
	(77356, 'BANDUA', 6287, 455, 34),
	(77357, 'BIMU', 6287, 455, 34),
	(77358, 'BOLAKME', 6287, 455, 34),
	(77359, 'KUGITERO', 6287, 455, 34),
	(77360, 'LANI TIMUR', 6287, 455, 34),
	(77361, 'MUNAK', 6287, 455, 34),
	(77362, 'NUNGGARUGUM', 6287, 455, 34),
	(77363, 'OWAGAMBAK', 6287, 455, 34),
	(77364, 'POITMOS', 6287, 455, 34),
	(77365, 'TEKANI', 6287, 455, 34),
	(77366, 'TONONGGAME (TENONGGAME)', 6287, 455, 34),
	(77367, 'WENAMELA', 6287, 455, 34),
	(77368, 'AYOMA', 6288, 455, 34),
	(77369, 'DLINGGAMA', 6288, 455, 34),
	(77370, 'DLONGGOKI', 6288, 455, 34),
	(77371, 'IRILINGA / IRILINGGA (IRILINGAN)', 6288, 455, 34),
	(77372, 'ONGGABA (ONGGABAGA/ONGGOBAGA)', 6288, 455, 34),
	(77373, 'TIRUNGGU', 6288, 455, 34),
	(77374, 'WALAKMA', 6288, 455, 34),
	(77375, 'AIR GARAM', 6289, 455, 34),
	(77376, 'BUGI', 6289, 455, 34),
	(77377, 'DEWENE', 6289, 455, 34),
	(77378, 'KODLANGGA', 6289, 455, 34),
	(77379, 'MANDA', 6289, 455, 34),
	(77380, 'TAGULIK', 6289, 455, 34),
	(77381, 'TOTNI', 6289, 455, 34),
	(77382, 'WALAK SELATAN', 6289, 455, 34),
	(77383, 'DOKOPKU', 6290, 455, 34),
	(77384, 'HITUMA (HETUMA)', 6290, 455, 34),
	(77385, 'HOM-HOM', 6290, 455, 34),
	(77386, 'HUBUKIAK / HUBIKIAK', 6290, 455, 34),
	(77387, 'HUSOAK', 6290, 455, 34),
	(77388, 'LIKINO', 6290, 455, 34),
	(77389, 'MUSIAMIA', 6290, 455, 34),
	(77390, 'MUSIAMIA DUA', 6290, 455, 34),
	(77391, 'HUBIKOSI', 6291, 455, 34),
	(77392, 'IKILUMO', 6291, 455, 34),
	(77393, 'ISAKUSA', 6291, 455, 34),
	(77394, 'JIBILABAGA', 6291, 455, 34),
	(77395, 'KIKHUMO', 6291, 455, 34),
	(77396, 'KOSIHILAPOK', 6291, 455, 34),
	(77397, 'KOSIMEAGE (KOSIMEAGA)', 6291, 455, 34),
	(77398, 'MEAGAMA', 6291, 455, 34),
	(77399, 'PELIMA', 6291, 455, 34),
	(77400, 'PIPUKMO', 6291, 455, 34),
	(77401, 'SUNILI', 6291, 455, 34),
	(77402, 'AYOBAIBUR (AYUBAIBUR)', 6292, 455, 34),
	(77403, 'HABEMA', 6292, 455, 34),
	(77404, 'HOLALIBA', 6292, 455, 34),
	(77405, 'IBELE', 6292, 455, 34),
	(77406, 'TIPALOK', 6292, 455, 34),
	(77407, 'YAGAROBAK (YAGAKOBAK)', 6292, 455, 34),
	(77408, 'YELEBAREK', 6292, 455, 34),
	(77409, 'YOKALPALEK', 6292, 455, 34),
	(77410, 'ZAPMA', 6292, 455, 34),
	(77411, 'ZINAI', 6292, 455, 34),
	(77412, 'HELEPALEGEM', 6293, 455, 34),
	(77413, 'KEMISAKE', 6293, 455, 34),
	(77414, 'LUKAKEN', 6293, 455, 34),
	(77415, 'MIAMI', 6293, 455, 34),
	(77416, 'SILIWA', 6293, 455, 34),
	(77417, 'SUMUNIKAMA', 6293, 455, 34),
	(77418, 'TOMISA', 6293, 455, 34),
	(77419, 'WAOBA (WUROBA/WAROBA)', 6293, 455, 34),
	(77420, 'YOGONIMA', 6293, 455, 34),
	(77421, 'KORAGI', 6294, 455, 34),
	(77422, 'KUMUDILUK', 6294, 455, 34),
	(77423, 'TAGIBAGA', 6294, 455, 34),
	(77424, 'TELEGAI', 6294, 455, 34),
	(77425, 'TENONDEK', 6294, 455, 34),
	(77426, 'ABUSA', 6295, 455, 34),
	(77427, 'ERAGAMA (EREMA)', 6295, 455, 34),
	(77428, 'HOPAMA', 6295, 455, 34),
	(77429, 'IYANTIK', 6295, 455, 34),
	(77430, 'KUMIMA (KIMINA)', 6295, 455, 34),
	(77431, 'MEBAGAIMA', 6295, 455, 34),
	(77432, 'MULIMA', 6295, 455, 34),
	(77433, 'OBYA', 6295, 455, 34),
	(77434, 'UMPAGALO', 6295, 455, 34),
	(77435, 'UTKOLO SATU', 6295, 455, 34),
	(77436, 'WAGA WAGA', 6295, 455, 34),
	(77437, 'YIWIKA (JIWIKA)', 6295, 455, 34),
	(77438, 'KILUBAGA', 6296, 455, 34),
	(77439, 'MULIAMA', 6296, 455, 34),
	(77440, 'MUSALFAK', 6296, 455, 34),
	(77441, 'PUNAKUL', 6296, 455, 34),
	(77442, 'WENABUBAGA', 6296, 455, 34),
	(77443, 'ESIAK', 6297, 455, 34),
	(77444, 'HERAEWA', 6297, 455, 34),
	(77445, 'HUSEWA', 6297, 455, 34),
	(77446, 'KEPI', 6297, 455, 34),
	(77447, 'MAIMA', 6297, 455, 34),
	(77448, 'MENAGAIMA', 6297, 455, 34),
	(77449, 'MINIMO', 6297, 455, 34),
	(77450, 'KWIGILUK', 6298, 455, 34),
	(77451, 'MEBUNUKME', 6298, 455, 34),
	(77452, 'MOLAGALOME', 6298, 455, 34),
	(77453, 'OKWA', 6298, 455, 34),
	(77454, 'TANAHMERAH', 6298, 455, 34),
	(77455, 'TOWAGAME', 6298, 455, 34),
	(77456, 'ASOLOGAIMA', 6299, 455, 34),
	(77457, 'DELEKAMA', 6299, 455, 34),
	(77458, 'HELEFA / HALEFA', 6299, 455, 34),
	(77459, 'HOLKIMA', 6299, 455, 34),
	(77460, 'HUKURAGI', 6299, 455, 34),
	(77461, 'KEWIN', 6299, 455, 34),
	(77462, 'KONAN', 6299, 455, 34),
	(77463, 'MILIAMA (MULIAMA)', 6299, 455, 34),
	(77464, 'MOLEBAGA', 6299, 455, 34),
	(77465, 'PILIBAGA', 6299, 455, 34),
	(77466, 'SEKOM', 6299, 455, 34),
	(77467, 'SILAMIK', 6299, 455, 34),
	(77468, 'ABULUKMO', 6300, 455, 34),
	(77469, 'ANEGARA (ANEGERA)', 6300, 455, 34),
	(77470, 'ELABUKAMA', 6300, 455, 34),
	(77471, 'HAMUHI', 6300, 455, 34),
	(77472, 'KOSIHAVE', 6300, 455, 34),
	(77473, 'MULUPALEK', 6300, 455, 34),
	(77474, 'PUMASILI', 6300, 455, 34),
	(77475, 'SIAPMA', 6300, 455, 34),
	(77476, 'TEMIA', 6300, 455, 34),
	(77477, 'YUMUGIMA', 6300, 455, 34),
	(77478, 'HEALEKMA', 6301, 455, 34),
	(77479, 'HOLIMA', 6301, 455, 34),
	(77480, 'LANI MATUAN', 6301, 455, 34),
	(77481, 'NAPUA', 6301, 455, 34),
	(77482, 'OKILIK', 6301, 455, 34),
	(77483, 'SAPALEK', 6301, 455, 34),
	(77484, 'WILEKAMA', 6301, 455, 34),
	(77485, 'YELEKAMA', 6301, 455, 34),
	(77486, 'YOMAIMA', 6301, 455, 34),
	(77487, 'ALEAK', 6302, 455, 34),
	(77488, 'DUABALEK', 6302, 455, 34),
	(77489, 'FILIA', 6302, 455, 34),
	(77490, 'HEATNEM', 6302, 455, 34),
	(77491, 'HITELOWA', 6302, 455, 34),
	(77492, 'INANEKELOK', 6302, 455, 34),
	(77493, 'ISUGUNIK', 6302, 455, 34),
	(77494, 'LANDIA', 6302, 455, 34),
	(77495, 'MULUKMO', 6302, 455, 34),
	(77496, 'WAUKAHILAPOK', 6302, 455, 34),
	(77497, 'WILILIMO', 6302, 455, 34),
	(77498, 'WITALAK', 6302, 455, 34),
	(77499, 'YABEM', 6302, 455, 34),
	(77500, 'ABONERI', 6303, 455, 34),
	(77501, 'ALGONIK', 6303, 455, 34),
	(77502, 'BALIMA', 6303, 455, 34),
	(77503, 'BEAM', 6303, 455, 34),
	(77504, 'GOBALIMO', 6303, 455, 34),
	(77505, 'PERABAGA', 6303, 455, 34),
	(77506, 'PIRAMID', 6303, 455, 34),
	(77507, 'YALINGGUME', 6303, 455, 34),
	(77508, 'YONGGIME', 6303, 455, 34),
	(77509, 'YUMBUN', 6303, 455, 34),
	(77510, 'AIKIMA', 6304, 455, 34),
	(77511, 'AKIAPUT', 6304, 455, 34),
	(77512, 'PABUMA', 6304, 455, 34),
	(77513, 'PIKHE', 6304, 455, 34),
	(77514, 'PISUGI', 6304, 455, 34),
	(77515, 'SUROBA', 6304, 455, 34),
	(77516, 'WARA', 6304, 455, 34),
	(77517, 'PUPUGOBA (POPUGOBA)', 6305, 455, 34),
	(77518, 'WAIMA', 6305, 455, 34),
	(77519, 'YELELO', 6305, 455, 34),
	(77520, 'YOMOTE', 6305, 455, 34),
	(77521, 'ISAWA HIMAN', 6306, 455, 34),
	(77522, 'LUNAIMA', 6306, 455, 34),
	(77523, 'MANIKA', 6306, 455, 34),
	(77524, 'NOAGALO', 6306, 455, 34),
	(77525, 'SEKAN', 6306, 455, 34),
	(77526, 'SEKAN DALAM', 6306, 455, 34),
	(77527, 'SIEPKOSI', 6306, 455, 34),
	(77528, 'WEKIA', 6306, 455, 34),
	(77529, 'YUMOGIMA (ISIBITLAH)', 6306, 455, 34),
	(77530, 'APNAE', 6307, 455, 34),
	(77531, 'ELABOGA (ELABOGE)', 6307, 455, 34),
	(77532, 'GIGILOBO', 6307, 455, 34),
	(77533, 'HOLASILI', 6307, 455, 34),
	(77534, 'OLAGI', 6307, 455, 34),
	(77535, 'WOGI', 6307, 455, 34),
	(77536, 'WONENGGULIK', 6307, 455, 34),
	(77537, 'YEREGA', 6307, 455, 34),
	(77538, 'BOKIEM', 6308, 455, 34),
	(77539, 'BUDLIEM', 6308, 455, 34),
	(77540, 'ENTAGEKOKMA', 6308, 455, 34),
	(77541, 'HULUAIMA', 6308, 455, 34),
	(77542, 'IYORA', 6308, 455, 34),
	(77543, 'SENOGOLIK', 6308, 455, 34),
	(77544, 'TAILAREK (TAILA)', 6308, 455, 34),
	(77545, 'YOMAN WEYA', 6308, 455, 34),
	(77546, 'ANGGULPA', 6309, 455, 34),
	(77547, 'DINGGILIMO', 6309, 455, 34),
	(77548, 'KORA JAYA', 6309, 455, 34),
	(77549, 'KORAMBIRIK', 6309, 455, 34),
	(77550, 'NANGGO', 6309, 455, 34),
	(77551, 'TRIKORA', 6309, 455, 34),
	(77552, 'ABUTPUK', 6310, 455, 34),
	(77553, 'ALONA', 6310, 455, 34),
	(77554, 'FIKHA', 6310, 455, 34),
	(77555, 'GUA WISATA', 6310, 455, 34),
	(77556, 'ISAIMAN', 6310, 455, 34),
	(77557, 'MEAGAMIA', 6310, 455, 34),
	(77558, 'SIBA', 6310, 455, 34),
	(77559, 'UNDULUMO', 6310, 455, 34),
	(77560, 'USILIMO', 6310, 455, 34),
	(77561, 'WOSIALA', 6310, 455, 34),
	(77562, 'AGULIMO', 6311, 455, 34),
	(77563, 'LUKU-LUKU', 6311, 455, 34),
	(77564, 'MUSIEM', 6311, 455, 34),
	(77565, 'WADANGKU', 6311, 455, 34),
	(77566, 'YOMOSIMO', 6311, 455, 34),
	(77567, 'ELAREK', 6312, 455, 34),
	(77568, 'HOLIMA', 6312, 455, 34),
	(77569, 'WALAIK', 6312, 455, 34),
	(77570, 'WELEKAMA', 6312, 455, 34),
	(77571, 'YELAI', 6312, 455, 34),
	(77572, 'ITLAY HALITOPO', 6313, 455, 34),
	(77573, 'KUBULAKMA', 6313, 455, 34),
	(77574, 'KULAKEN', 6313, 455, 34),
	(77575, 'PUGIMA', 6313, 455, 34),
	(77576, 'WALELAGAMA', 6313, 455, 34),
	(77577, 'WAMUSAGE', 6313, 455, 34),
	(77578, 'DOGONAME', 6314, 455, 34),
	(77579, 'DUMAPAGA', 6314, 455, 34),
	(77580, 'WAME', 6314, 455, 34),
	(77581, 'WANGGONOMA', 6314, 455, 34),
	(77582, 'YANENGGAME', 6314, 455, 34),
	(77583, 'AUTAKMA', 6315, 455, 34),
	(77584, 'BATU MERAH', 6315, 455, 34),
	(77585, 'HONAIMA', 6315, 455, 34),
	(77586, 'HONELAMA', 6315, 455, 34),
	(77587, 'HONELAMA DUA', 6315, 455, 34),
	(77588, 'HULEKAMA / HUREKAMA', 6315, 455, 34),
	(77589, 'LANITIPO', 6315, 455, 34),
	(77590, 'SINAKMA', 6315, 455, 34),
	(77591, 'SINAPUK', 6315, 455, 34),
	(77592, 'WAMAROMA', 6315, 455, 34),
	(77593, 'WAMENA KOTA', 6315, 455, 34),
	(77594, 'APENAS', 6316, 455, 34),
	(77595, 'ASOJELIPELE', 6316, 455, 34),
	(77596, 'LANTIPO', 6316, 455, 34),
	(77597, 'PAWEKAMA', 6316, 455, 34),
	(77598, 'TULIMA', 6316, 455, 34),
	(77599, 'WELESI', 6316, 455, 34),
	(77600, 'YAGARA', 6316, 455, 34),
	(77601, 'AGAMOA', 6317, 455, 34),
	(77602, 'ILOKAMA', 6317, 455, 34),
	(77603, 'KAMA', 6317, 455, 34),
	(77604, 'MAWAMPI', 6317, 455, 34),
	(77605, 'PAREMA', 6317, 455, 34),
	(77606, 'SILUMAREK', 6317, 455, 34),
	(77607, 'WESAGANYA', 6317, 455, 34),
	(77608, 'YALOAPUT', 6317, 455, 34),
	(77609, 'ALOLIK', 6318, 455, 34),
	(77610, 'ALULA', 6318, 455, 34),
	(77611, 'KOMA', 6318, 455, 34),
	(77612, 'TULEM', 6318, 455, 34),
	(77613, 'WILIGIMA', 6318, 455, 34),
	(77614, 'ALUGI', 6319, 455, 34),
	(77615, 'KUKURIMA', 6319, 455, 34),
	(77616, 'PIRAMBOR (PIRAMBOT)', 6319, 455, 34),
	(77617, 'TEGABAGA', 6319, 455, 34),
	(77618, 'WODLOMA', 6319, 455, 34),
	(77619, 'WOLLO TIMUR', 6319, 455, 34),
	(77620, 'WOLO (WOLLO)', 6319, 455, 34),
	(77621, 'WUNAN', 6319, 455, 34),
	(77622, 'KETIMAVIT', 6320, 455, 34),
	(77623, 'LOGONOBA', 6320, 455, 34),
	(77624, 'PIPITMO', 6320, 455, 34),
	(77625, 'SINAREKOWA', 6320, 455, 34),
	(77626, 'WESAKIN', 6320, 455, 34),
	(77627, 'WESAKMA', 6320, 455, 34),
	(77628, 'WOUMA', 6320, 455, 34),
	(77629, 'AIPAKMA', 6321, 455, 34),
	(77630, 'AKOREK', 6321, 455, 34),
	(77631, 'BITTI', 6321, 455, 34),
	(77632, 'MANILI', 6321, 455, 34),
	(77633, 'PILIMO', 6321, 455, 34),
	(77634, 'TAGANIK', 6321, 455, 34),
	(77635, 'TUMUN', 6321, 455, 34),
	(77636, 'WAMANUK DUA', 6321, 455, 34),
	(77637, 'WANANUK', 6321, 455, 34),
	(77638, 'WUGURIMA', 6321, 455, 34),
	(77639, 'YALENGGA (YELENGGA)', 6321, 455, 34),
	(77640, 'ARSO KOTA', 6322, 456, 34),
	(77641, 'ARSO SWAKARSA (ASIAMAN)', 6322, 456, 34),
	(77642, 'BAGIA (PIR 3)', 6322, 456, 34),
	(77643, 'DUKWIA (ARSO 8)', 6322, 456, 34),
	(77644, 'IFIA-FIA (ARSO 11)', 6322, 456, 34),
	(77645, 'KWIMI (KWIME)', 6322, 456, 34),
	(77646, 'NAWA (SAWANAWA)', 6322, 456, 34),
	(77647, 'SANGGARIA (ARSO 1)', 6322, 456, 34),
	(77648, 'SAWIATAMI (SAWYATAMI)', 6322, 456, 34),
	(77649, 'UBIYAU', 6322, 456, 34),
	(77650, 'WARBO (ARSO 7)', 6322, 456, 34),
	(77651, 'WORKWANA', 6322, 456, 34),
	(77652, 'YAMMUA (YAMUA/ARSO 6)', 6322, 456, 34),
	(77653, 'YAMTA (PIR 2)', 6322, 456, 34),
	(77654, 'YANAMA (PIR 1)', 6322, 456, 34),
	(77655, 'YATURAHARJA (ARSO 10)', 6322, 456, 34),
	(77656, 'YUWANAIM (ARSO 2)', 6322, 456, 34),
	(77657, 'KIBAY', 6323, 456, 34),
	(77658, 'KRIKU', 6323, 456, 34),
	(77659, 'PYAWI (WEMBI LUAR)', 6323, 456, 34),
	(77660, 'SANGKE', 6323, 456, 34),
	(77661, 'SKOFRO', 6323, 456, 34),
	(77662, 'SUSKUN (WAMBES DALAM)', 6323, 456, 34),
	(77663, 'WAMBES LUAR', 6323, 456, 34),
	(77664, 'WEMBI DALAM', 6323, 456, 34),
	(77665, 'WONOREJO (PIR 4)', 6323, 456, 34),
	(77666, 'YAMARA (PIR 5)', 6323, 456, 34),
	(77667, 'YETTY', 6323, 456, 34),
	(77668, 'MOLOF', 6324, 456, 34),
	(77669, 'SENGGI', 6324, 456, 34),
	(77670, 'USKU', 6324, 456, 34),
	(77671, 'WARLEF', 6324, 456, 34),
	(77672, 'WOSLAY', 6324, 456, 34),
	(77673, 'YABANDA (JABANDA)', 6324, 456, 34),
	(77674, 'ARSOPURA', 6325, 456, 34),
	(77675, 'INTAIMELYAN (INTAIMILYAN)', 6325, 456, 34),
	(77676, 'JAIFURI', 6325, 456, 34),
	(77677, 'NARAMBEN', 6325, 456, 34),
	(77678, 'SKANTO (SKAMTO)', 6325, 456, 34),
	(77679, 'TRAIMELYAN (TRAIMILYAN)', 6325, 456, 34),
	(77680, 'WIYANTRE (WIANTIE)', 6325, 456, 34),
	(77681, 'WULUKBUBUN', 6325, 456, 34),
	(77682, 'BIAS', 6326, 456, 34),
	(77683, 'LULES', 6326, 456, 34),
	(77684, 'MILKY (MILKI)', 6326, 456, 34),
	(77685, 'TEFANMA', 6326, 456, 34),
	(77686, 'TERFONES / TERPONES', 6326, 456, 34),
	(77687, 'TOWE ATAS', 6326, 456, 34),
	(77688, 'TOWE HITAM', 6326, 456, 34),
	(77689, 'AMPAS', 6327, 456, 34),
	(77690, 'BANDA', 6327, 456, 34),
	(77691, 'KALIFAM', 6327, 456, 34),
	(77692, 'KALIMO', 6327, 456, 34),
	(77693, 'PUND', 6327, 456, 34),
	(77694, 'YUWAINDA', 6327, 456, 34),
	(77695, 'AMGOTRO', 6328, 456, 34),
	(77696, 'DUBU', 6328, 456, 34),
	(77697, 'EMBI', 6328, 456, 34),
	(77698, 'SEMOGRAFI', 6328, 456, 34),
	(77699, 'UMUAF', 6328, 456, 34),
	(77700, 'YURUF', 6328, 456, 34),
	(77701, 'AITIRI', 6329, 457, 34),
	(77702, 'BORAI', 6329, 457, 34),
	(77703, 'KABUAENA', 6329, 457, 34),
	(77704, 'KAINUI I', 6329, 457, 34),
	(77705, 'KAINUI II', 6329, 457, 34),
	(77706, 'KONTIUNAI (KONTINUAI)', 6329, 457, 34),
	(77707, 'MANANAYAM', 6329, 457, 34),
	(77708, 'MENAWI', 6329, 457, 34),
	(77709, 'RAMBAI', 6329, 457, 34),
	(77710, 'RANSARNONI (RAMSAMONI)', 6329, 457, 34),
	(77711, 'ROIPI', 6329, 457, 34),
	(77712, 'ROIPI DUA', 6329, 457, 34),
	(77713, 'SANAYOKA', 6329, 457, 34),
	(77714, 'WADAPI', 6329, 457, 34),
	(77715, 'WANAMPOMPI', 6329, 457, 34),
	(77716, 'WANIWON', 6329, 457, 34),
	(77717, 'WAWUTI', 6329, 457, 34),
	(77718, 'YAPANANI', 6329, 457, 34),
	(77719, 'ADIWIPI', 6330, 457, 34),
	(77720, 'AIWARAGGANI', 6330, 457, 34),
	(77721, 'AMBAI I', 6330, 457, 34),
	(77722, 'AMBAI II', 6330, 457, 34),
	(77723, 'BAIREI', 6330, 457, 34),
	(77724, 'BAISORE', 6330, 457, 34),
	(77725, 'BAISORE', 6330, 457, 34),
	(77726, 'DORAU', 6330, 457, 34),
	(77727, 'IMBORIAWA', 6330, 457, 34),
	(77728, 'KAWIPI', 6330, 457, 34),
	(77729, 'MAMBAWI (AMBAI III)', 6330, 457, 34),
	(77730, 'MARAWI', 6330, 457, 34),
	(77731, 'NUNIANDEI', 6330, 457, 34),
	(77732, 'PEREA', 6330, 457, 34),
	(77733, 'RONDEPI', 6330, 457, 34),
	(77734, 'SAWERU', 6330, 457, 34),
	(77735, 'TOROA', 6330, 457, 34),
	(77736, 'UMANI', 6330, 457, 34),
	(77737, 'WAMORI', 6330, 457, 34),
	(77738, 'AMBAIDIRU', 6331, 457, 34),
	(77739, 'ARIEPI', 6331, 457, 34),
	(77740, 'ARIEPI DUA', 6331, 457, 34),
	(77741, 'AROMAREA', 6331, 457, 34),
	(77742, 'KAMANAP', 6331, 457, 34),
	(77743, 'KANAWA', 6331, 457, 34),
	(77744, 'MAMBO', 6331, 457, 34),
	(77745, 'MANAININ', 6331, 457, 34),
	(77746, 'MARIAROTU (MARIAR ROTU)', 6331, 457, 34),
	(77747, 'NUMAMAN', 6331, 457, 34),
	(77748, 'PANDUAMI (PANDOAMI/PANDUANI)', 6331, 457, 34),
	(77749, 'RAMANGKURANI', 6331, 457, 34),
	(77750, 'SARAWANDORI DUA', 6331, 457, 34),
	(77751, 'SERAWANDORI (SARAWANDORI)', 6331, 457, 34),
	(77752, 'TATUI', 6331, 457, 34),
	(77753, 'HUMBE AWAI', 6332, 457, 34),
	(77754, 'MAKIROAN (MOKIROAN)', 6332, 457, 34),
	(77755, 'NURAWI', 6332, 457, 34),
	(77756, 'POOM I', 6332, 457, 34),
	(77757, 'POOM II', 6332, 457, 34),
	(77758, 'RASIRI (WOISIRI)', 6332, 457, 34),
	(77759, 'SEREWEN', 6332, 457, 34),
	(77760, 'WORIOI (WARIORI)', 6332, 457, 34),
	(77761, 'ANDESARIA (ANDERSARIA)', 6333, 457, 34),
	(77762, 'DOREIAMINI (DOREIANMINI)', 6333, 457, 34),
	(77763, 'KAIPURI', 6333, 457, 34),
	(77764, 'KIRIMBRI', 6333, 457, 34),
	(77765, 'KURUDU', 6333, 457, 34),
	(77766, 'MANSESI', 6333, 457, 34),
	(77767, 'MANUSUNDU', 6333, 457, 34),
	(77768, 'MNUKWAR', 6333, 457, 34),
	(77769, 'AUSEM', 6334, 457, 34),
	(77770, 'JENIARI', 6334, 457, 34),
	(77771, 'MIOSNUM', 6334, 457, 34),
	(77772, 'UMPEKI', 6334, 457, 34),
	(77773, 'YEITUARAU', 6334, 457, 34),
	(77774, 'AISAU', 6335, 457, 34),
	(77775, 'BARAWAI', 6335, 457, 34),
	(77776, 'BARAWAI II', 6335, 457, 34),
	(77777, 'KOROROMPUI', 6335, 457, 34),
	(77778, 'SAWENDUI', 6335, 457, 34),
	(77779, 'SAWENUI (SAWENOI)', 6335, 457, 34),
	(77780, 'WAINDU', 6335, 457, 34),
	(77781, 'WODA', 6335, 457, 34),
	(77782, 'AMPIMOI', 6336, 457, 34),
	(77783, 'ARARENI', 6336, 457, 34),
	(77784, 'AYARI', 6336, 457, 34),
	(77785, 'BARERAIPI (BARERAI/BARERAIF)', 6336, 457, 34),
	(77786, 'KAROAIPI (KOROAPI)', 6336, 457, 34),
	(77787, 'RANDAWAYA', 6336, 457, 34),
	(77788, 'SIROMI', 6336, 457, 34),
	(77789, 'TAREI', 6336, 457, 34),
	(77790, 'WABUAYAR (WABUAYER)', 6336, 457, 34),
	(77791, 'WAITA', 6336, 457, 34),
	(77792, 'WARIRONI', 6336, 457, 34),
	(77793, 'AIBONDENI', 6337, 457, 34),
	(77794, 'AWADO', 6337, 457, 34),
	(77795, 'DUMANI', 6337, 457, 34),
	(77796, 'HAIHOREI', 6337, 457, 34),
	(77797, 'JAIMARIA', 6337, 457, 34),
	(77798, 'KANAKI', 6337, 457, 34),
	(77799, 'KAREMONI', 6337, 457, 34),
	(77800, 'REMBAI', 6337, 457, 34),
	(77801, 'WOINAP', 6337, 457, 34),
	(77802, 'WOOI', 6337, 457, 34),
	(77803, 'ANSUS', 6338, 457, 34),
	(77804, 'INOWA', 6338, 457, 34),
	(77805, 'KAIRAWI', 6338, 457, 34),
	(77806, 'MANIRI', 6338, 457, 34),
	(77807, 'MARAU', 6338, 457, 34),
	(77808, 'NAREI', 6338, 457, 34),
	(77809, 'NATABUI', 6338, 457, 34),
	(77810, 'NUIWIORA', 6338, 457, 34),
	(77811, 'PAPUMA (PAPUAMA)', 6338, 457, 34),
	(77812, 'SASAWA', 6338, 457, 34),
	(77813, 'TOWETA', 6338, 457, 34),
	(77814, 'WARABORI', 6338, 457, 34),
	(77815, 'WEBI', 6338, 457, 34),
	(77816, 'WIMONI', 6338, 457, 34),
	(77817, 'WOIWANI', 6338, 457, 34),
	(77818, 'YARORI', 6338, 457, 34),
	(77819, 'YENUSI MARAU', 6338, 457, 34),
	(77820, 'ANATAUREI (ANATOREI)', 6339, 457, 34),
	(77821, 'BANAWA', 6339, 457, 34),
	(77822, 'BARAWAIKAP', 6339, 457, 34),
	(77823, 'BAWAI', 6339, 457, 34),
	(77824, 'FAMBOAMAN (FAMBUAMAN)', 6339, 457, 34),
	(77825, 'IMANDOA', 6339, 457, 34),
	(77826, 'KANDOWARIRA', 6339, 457, 34),
	(77827, 'KETUAPI', 6339, 457, 34),
	(77828, 'MANAINI', 6339, 457, 34),
	(77829, 'MANTEMBU', 6339, 457, 34),
	(77830, 'MARIADEI (MARIADERI)', 6339, 457, 34),
	(77831, 'NUNDAWIPI', 6339, 457, 34),
	(77832, 'PASIR HITAM', 6339, 457, 34),
	(77833, 'PASIR PUTIH', 6339, 457, 34),
	(77834, 'SERUI JAYA', 6339, 457, 34),
	(77835, 'SERUI KOTA', 6339, 457, 34),
	(77836, 'SERUI LAUT', 6339, 457, 34),
	(77837, 'TARAU', 6339, 457, 34),
	(77838, 'TURU', 6339, 457, 34),
	(77839, 'WARARI', 6339, 457, 34),
	(77840, 'YAPAN', 6339, 457, 34),
	(77841, 'AWUNAWAI', 6340, 457, 34),
	(77842, 'DAWAI', 6340, 457, 34),
	(77843, 'DUAI', 6340, 457, 34),
	(77844, 'KOROMBOBI (KAROMBOBI)', 6340, 457, 34),
	(77845, 'MERERUNI (MASERUNI)', 6340, 457, 34),
	(77846, 'NUNSEMBAI (HUNSEMBAI)', 6340, 457, 34),
	(77847, 'NUNSIARI', 6340, 457, 34),
	(77848, 'SERE SERE', 6340, 457, 34),
	(77849, 'WABO', 6340, 457, 34),
	(77850, 'WABOMPI', 6340, 457, 34),
	(77851, 'WOINSUPI (WAISOPI)', 6340, 457, 34),
	(77852, 'KARAWI', 6341, 457, 34),
	(77853, 'KIRIYOU (KIRIYOW)', 6341, 457, 34),
	(77854, 'ROSWARI (ARTANENG)', 6341, 457, 34),
	(77855, 'SAMBRAWAI (SAMRAWAI)', 6341, 457, 34),
	(77856, 'SOROMASEN (SOROMASEM)', 6341, 457, 34),
	(77857, 'TINDARET', 6341, 457, 34),
	(77858, 'YOBI', 6341, 457, 34),
	(77859, 'ANITILA', 6342, 458, 34),
	(77860, 'BALIME', 6342, 458, 34),
	(77861, 'BALIME (BALINGGA)', 6342, 458, 34),
	(77862, 'BALIMNERI', 6342, 458, 34),
	(77863, 'BRUYUGU', 6342, 458, 34),
	(77864, 'EKABA', 6342, 458, 34),
	(77865, 'GUME', 6342, 458, 34),
	(77866, 'OGODOME (ODAGOME)', 6342, 458, 34),
	(77867, 'POPOME', 6342, 458, 34),
	(77868, 'TIKOME', 6342, 458, 34),
	(77869, 'TIMA', 6342, 458, 34),
	(77870, 'TINGGIPURA', 6342, 458, 34),
	(77871, 'WAME', 6342, 458, 34),
	(77872, 'YEYUGU (GILUBANDU)', 6342, 458, 34),
	(77873, 'YUGUME', 6342, 458, 34),
	(77874, 'DIMBA', 6343, 458, 34),
	(77875, 'KELULOME', 6343, 458, 34),
	(77876, 'LANGGIME', 6343, 458, 34),
	(77877, 'MAGEGOBAK', 6343, 458, 34),
	(77878, 'TOLOGI', 6343, 458, 34),
	(77879, 'WANGGAGOME', 6343, 458, 34),
	(77880, 'WIYAGI', 6343, 458, 34),
	(77881, 'WULUWA', 6343, 458, 34),
	(77882, 'YUGWA (YUKWA)', 6343, 458, 34),
	(77883, 'AYAFOFA', 6344, 458, 34),
	(77884, 'EKAPAME', 6344, 458, 34),
	(77885, 'GAMELIA', 6344, 458, 34),
	(77886, 'GUKOPI (GUKOP)', 6344, 458, 34),
	(77887, 'ODIKA', 6344, 458, 34),
	(77888, 'PINDALO', 6344, 458, 34),
	(77889, 'PIRAWUN', 6344, 458, 34),
	(77890, 'PIWUGUN', 6344, 458, 34),
	(77891, 'SALLEMO (SALEMO)', 6344, 458, 34),
	(77892, 'WUPI', 6344, 458, 34),
	(77893, 'YUDANI', 6344, 458, 34),
	(77894, 'YUGUMABUR (YUBUMABUR)', 6344, 458, 34),
	(77895, 'ANDUGUME', 6345, 458, 34),
	(77896, 'DUGU - DUGU', 6345, 458, 34),
	(77897, 'KUYAWAGE', 6345, 458, 34),
	(77898, 'LUAREM', 6345, 458, 34),
	(77899, 'MUME', 6345, 458, 34),
	(77900, 'NENGGEYA', 6345, 458, 34),
	(77901, 'TINIME', 6345, 458, 34),
	(77902, 'WAMITU', 6345, 458, 34),
	(77903, 'WAMLRU', 6345, 458, 34),
	(77904, 'WUPAGA', 6345, 458, 34),
	(77905, 'BONOM', 6346, 458, 34),
	(77906, 'EYUNI', 6346, 458, 34),
	(77907, 'GONDURA', 6346, 458, 34),
	(77908, 'GUBURNI', 6346, 458, 34),
	(77909, 'INDAWA', 6346, 458, 34),
	(77910, 'INDIGU (INDUGU)', 6346, 458, 34),
	(77911, 'JUTA', 6346, 458, 34),
	(77912, 'KAMULUME (KUMULUME)', 6346, 458, 34),
	(77913, 'KELOYAK (KOLAYAK)', 6346, 458, 34),
	(77914, 'KEMIRI', 6346, 458, 34),
	(77915, 'KIMBO', 6346, 458, 34),
	(77916, 'KONDENA', 6346, 458, 34),
	(77917, 'LABORA', 6346, 458, 34),
	(77918, 'LELAM', 6346, 458, 34),
	(77919, 'MAMIRI', 6346, 458, 34),
	(77920, 'PENGGIMA (PENGGIME)', 6346, 458, 34),
	(77921, 'PIRAMBOR', 6346, 458, 34),
	(77922, 'TEYIKO (SEIKO)', 6346, 458, 34),
	(77923, 'TIGIMA (TIGIME)', 6346, 458, 34),
	(77924, 'WANUGA', 6346, 458, 34),
	(77925, 'YANEKO', 6346, 458, 34),
	(77926, 'YUGUMIA', 6346, 458, 34),
	(77927, 'BAGI', 6347, 458, 34),
	(77928, 'MALAGAI', 6347, 458, 34),
	(77929, 'OKA', 6347, 458, 34),
	(77930, 'WABIRAGI', 6347, 458, 34),
	(77931, 'YIGEMILI', 6347, 458, 34),
	(77932, 'EKANOM', 6348, 458, 34),
	(77933, 'GOLO', 6348, 458, 34),
	(77934, 'GOLOMI', 6348, 458, 34),
	(77935, 'GOLOPURA', 6348, 458, 34),
	(77936, 'ILUNGGUME (ILUNGGAME/ILUNGGIME)', 6348, 458, 34),
	(77937, 'KARUNGGAME', 6348, 458, 34),
	(77938, 'KUGAME', 6348, 458, 34),
	(77939, 'KULIA', 6348, 458, 34),
	(77940, 'KUWANOM', 6348, 458, 34),
	(77941, 'LIBOME', 6348, 458, 34),
	(77942, 'LUBUTINI', 6348, 458, 34),
	(77943, 'MELENDIK', 6348, 458, 34),
	(77944, 'MILIMBO', 6348, 458, 34),
	(77945, 'NAMBUME', 6348, 458, 34),
	(77946, 'NANIM', 6348, 458, 34),
	(77947, 'NILEME', 6348, 458, 34),
	(77948, 'PIRIME', 6348, 458, 34),
	(77949, 'TAKOBAK', 6348, 458, 34),
	(77950, 'TEKUL', 6348, 458, 34),
	(77951, 'UMBANUME', 6348, 458, 34),
	(77952, 'WAMINDIK', 6348, 458, 34),
	(77953, 'WENAM', 6348, 458, 34),
	(77954, 'WIRINI', 6348, 458, 34),
	(77955, 'YALIPAK', 6348, 458, 34),
	(77956, 'YAMIGA', 6348, 458, 34),
	(77957, 'YIWILI', 6348, 458, 34),
	(77958, 'YUGUMABUR (YUGUMOBUR)', 6348, 458, 34),
	(77959, 'ABUA', 6349, 458, 34),
	(77960, 'BIGIPAGA', 6349, 458, 34),
	(77961, 'GIKUR', 6349, 458, 34),
	(77962, 'GIPURA', 6349, 458, 34),
	(77963, 'KANUMBUME', 6349, 458, 34),
	(77964, 'LUALO', 6349, 458, 34),
	(77965, 'LUGOBAK', 6349, 458, 34),
	(77966, 'MEGALUNIK', 6349, 458, 34),
	(77967, 'MUARA', 6349, 458, 34),
	(77968, 'POGA', 6349, 458, 34),
	(77969, 'BOKON', 6350, 458, 34),
	(77970, 'DUGUME', 6350, 458, 34),
	(77971, 'DURA', 6350, 458, 34),
	(77972, 'GIWAN', 6350, 458, 34),
	(77973, 'GOLIKME (GOLIME)', 6350, 458, 34),
	(77974, 'GUNINGGAME', 6350, 458, 34),
	(77975, 'GURIKA', 6350, 458, 34),
	(77976, 'KONIKME', 6350, 458, 34),
	(77977, 'KUMULUK', 6350, 458, 34),
	(77978, 'KWENUKWI', 6350, 458, 34),
	(77979, 'LONGGALO', 6350, 458, 34),
	(77980, 'LUGOM (LOGOM)', 6350, 458, 34),
	(77981, 'MOKONI', 6350, 458, 34),
	(77982, 'NAME', 6350, 458, 34),
	(77983, 'NINABUA', 6350, 458, 34),
	(77984, 'NOBO (GUBO)', 6350, 458, 34),
	(77985, 'OLUME', 6350, 458, 34),
	(77986, 'OYI', 6350, 458, 34),
	(77987, 'WULUNDIA', 6350, 458, 34),
	(77988, 'YIRENE', 6350, 458, 34),
	(77989, 'YOGOBAK (YAKOBAK)', 6350, 458, 34),
	(77990, 'ARGENERI (ARIGENERI/ARGINERI)', 6351, 458, 34),
	(77991, 'BONANIP', 6351, 458, 34),
	(77992, 'GUARI (GIARI/GUARE)', 6351, 458, 34),
	(77993, 'GUBO (NOBO)', 6351, 458, 34),
	(77994, 'KUABAGA (KUOBAGA)', 6351, 458, 34),
	(77995, 'KUKEPAKE', 6351, 458, 34),
	(77996, 'MABUME (NABUNE)', 6351, 458, 34),
	(77997, 'MILINGGAME (MALIMNERI)', 6351, 458, 34),
	(77998, 'PONALO (PANALO)', 6351, 458, 34),
	(77999, 'TABUKEKER', 6351, 458, 34),
	(78000, 'TIMI', 6351, 458, 34)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 77")
}
