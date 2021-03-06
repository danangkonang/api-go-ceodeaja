package kel70

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Kel72() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(72001, 'NDUNGGA', 5749, 418, 32),
	(72002, 'REWARANGGA', 5749, 418, 32),
	(72003, 'REWARANGGA SELATAN', 5749, 418, 32),
	(72004, 'TIWUTEWA', 5749, 418, 32),
	(72005, 'BOROKANDA', 5750, 418, 32),
	(72006, 'EMBUNDOA', 5750, 418, 32),
	(72007, 'GHEOGHOMA', 5750, 418, 32),
	(72008, 'KOTA RAJA', 5750, 418, 32),
	(72009, 'KOTA RATU', 5750, 418, 32),
	(72010, 'MBOMBA', 5750, 418, 32),
	(72011, 'RATERU', 5750, 418, 32),
	(72012, 'ROWORENA', 5750, 418, 32),
	(72013, 'ROWORENA BARAT', 5750, 418, 32),
	(72014, 'WATUSIPI', 5750, 418, 32),
	(72015, 'KOANARA', 5751, 418, 32),
	(72016, 'NDUARIA', 5751, 418, 32),
	(72017, 'NUAMURI', 5751, 418, 32),
	(72018, 'NUAMURI BARAT', 5751, 418, 32),
	(72019, 'PEMO', 5751, 418, 32),
	(72020, 'WOLOARA', 5751, 418, 32),
	(72021, 'CIKAMPEK UTARA', 5752, 418, 32),
	(72022, 'HANGALANDE', 5752, 418, 32),
	(72023, 'JOMIN BARAT', 5752, 418, 32),
	(72024, 'JOMIN TIMUR', 5752, 418, 32),
	(72025, 'KOTABARU', 5752, 418, 32),
	(72026, 'LISELANDE', 5752, 418, 32),
	(72027, 'LOBONIKI', 5752, 418, 32),
	(72028, 'NDONDO', 5752, 418, 32),
	(72029, 'NIOPANDA', 5752, 418, 32),
	(72030, 'PANGULAH BARU', 5752, 418, 32),
	(72031, 'PANGULAH SELATAN', 5752, 418, 32),
	(72032, 'PANGULAH UTARA', 5752, 418, 32),
	(72033, 'PUCUNG', 5752, 418, 32),
	(72034, 'RANGALAKA', 5752, 418, 32),
	(72035, 'SARIMULYA', 5752, 418, 32),
	(72036, 'TOU', 5752, 418, 32),
	(72037, 'TOU BARAT', 5752, 418, 32),
	(72038, 'TOU TIMUR', 5752, 418, 32),
	(72039, 'WANCIMEKAR', 5752, 418, 32),
	(72040, 'DETUARA', 5753, 418, 32),
	(72041, 'KURU', 5753, 418, 32),
	(72042, 'KURU SARE', 5753, 418, 32),
	(72043, 'LISE KURU', 5753, 418, 32),
	(72044, 'MUKUREKU', 5753, 418, 32),
	(72045, 'MUKUREKU SA ATE', 5753, 418, 32),
	(72046, 'NDENGGA RONGGE', 5753, 418, 32),
	(72047, 'NDIKOSAPU', 5753, 418, 32),
	(72048, 'NGGUMBELAKA', 5753, 418, 32),
	(72049, 'RUTU JEJA', 5753, 418, 32),
	(72050, 'TANALANGI', 5753, 418, 32),
	(72051, 'TANIWODA', 5753, 418, 32),
	(72052, 'TIWUSORA', 5753, 418, 32),
	(72053, 'WOLOGAI TIMUR', 5753, 418, 32),
	(72054, 'BU TANALAGU', 5754, 418, 32),
	(72055, 'DETUPERA', 5754, 418, 32),
	(72056, 'FATAMARI', 5754, 418, 32),
	(72057, 'HUBATUWA/HOBATUA', 5754, 418, 32),
	(72058, 'LIABEKE', 5754, 418, 32),
	(72059, 'MBEWAWORA', 5754, 418, 32),
	(72060, 'NUA LIMA', 5754, 418, 32),
	(72061, 'RANGGATALO', 5754, 418, 32),
	(72062, 'TANAROGA', 5754, 418, 32),
	(72063, 'WATUNESO', 5754, 418, 32),
	(72064, 'WOLOARO/WOLOARA', 5754, 418, 32),
	(72065, 'WOLOLELEA (WOLOLELA A)', 5754, 418, 32),
	(72066, 'WOLOSAMBI', 5754, 418, 32),
	(72067, 'BOAFEO', 5755, 418, 32),
	(72068, 'KAMUMBHEKA/KAMUBHEKA', 5755, 418, 32),
	(72069, 'KEBIRANGGA', 5755, 418, 32),
	(72070, 'KEBIRANGGA SELATAN', 5755, 418, 32),
	(72071, 'KEBIRANGGA TENGAH', 5755, 418, 32),
	(72072, 'KOBALEBA', 5755, 418, 32),
	(72073, 'KOLIPAKA', 5755, 418, 32),
	(72074, 'MAGEKAPA', 5755, 418, 32),
	(72075, 'MUNDINGGASA', 5755, 418, 32),
	(72076, 'NABE', 5755, 418, 32),
	(72077, 'NATANANGGE', 5755, 418, 32),
	(72078, 'AEWORA', 5756, 418, 32),
	(72079, 'DETUWALA', 5756, 418, 32),
	(72080, 'KELIWUMBU', 5756, 418, 32),
	(72081, 'MAUROLE', 5756, 418, 32),
	(72082, 'MAUSAMBI', 5756, 418, 32),
	(72083, 'NGALUKOJA', 5756, 418, 32),
	(72084, 'NIRANUSA', 5756, 418, 32),
	(72085, 'OTOGEDU', 5756, 418, 32),
	(72086, 'RANAKOLO', 5756, 418, 32),
	(72087, 'RANOKOLO SELATAN', 5756, 418, 32),
	(72088, 'ULUDALA', 5756, 418, 32),
	(72089, 'WATUKAMBA', 5756, 418, 32),
	(72090, 'WOLOAU', 5756, 418, 32),
	(72091, 'ANARAJA', 5757, 418, 32),
	(72092, 'BHERAMARI', 5757, 418, 32),
	(72093, 'EMBUZOZO', 5757, 418, 32),
	(72094, 'JEGHARONGGO', 5757, 418, 32),
	(72095, 'JEMBUREA', 5757, 418, 32),
	(72096, 'KEKANDERE', 5757, 418, 32),
	(72097, 'KERIREA', 5757, 418, 32),
	(72098, 'MALAWARU', 5757, 418, 32),
	(72099, 'MBOBHENGA', 5757, 418, 32),
	(72100, 'NDETUREA', 5757, 418, 32),
	(72101, 'NDORUREA', 5757, 418, 32),
	(72102, 'NDORUREA I', 5757, 418, 32),
	(72103, 'NGGOREA', 5757, 418, 32),
	(72104, 'ONDOREA', 5757, 418, 32),
	(72105, 'ONDOREA BARAT', 5757, 418, 32),
	(72106, 'PENGGAJAWA', 5757, 418, 32),
	(72107, 'RAPORENDU', 5757, 418, 32),
	(72108, 'RAPOWAWO', 5757, 418, 32),
	(72109, 'ROMAREA', 5757, 418, 32),
	(72110, 'SANGGARHORHO', 5757, 418, 32),
	(72111, 'TANAZOZO', 5757, 418, 32),
	(72112, 'TENDA ONDO', 5757, 418, 32),
	(72113, 'TENDAMBEPA', 5757, 418, 32),
	(72114, 'TENDAREA', 5757, 418, 32),
	(72115, 'TIMBARIA', 5757, 418, 32),
	(72116, 'TITWEREA', 5757, 418, 32),
	(72117, 'UZUZOZO', 5757, 418, 32),
	(72118, 'WATUMITE', 5757, 418, 32),
	(72119, 'ZOZOZEA', 5757, 418, 32),
	(72120, 'KEKASEWA', 5758, 418, 32),
	(72121, 'KELIKIKU', 5758, 418, 32),
	(72122, 'LOKOBOKO', 5758, 418, 32),
	(72123, 'MANULONDO (MANULANDO)', 5758, 418, 32),
	(72124, 'NANGANESA', 5758, 418, 32),
	(72125, 'NGALUPOLO', 5758, 418, 32),
	(72126, 'NGAUROGA/NGALUROGA', 5758, 418, 32),
	(72127, 'NILA', 5758, 418, 32),
	(72128, 'ONELAKO', 5758, 418, 32),
	(72129, 'PUUTUGA', 5758, 418, 32),
	(72130, 'REKA', 5758, 418, 32),
	(72131, 'WOLOKOTA', 5758, 418, 32),
	(72132, 'WOLOTOPO', 5758, 418, 32),
	(72133, 'WOLOTOPO TIMUR', 5758, 418, 32),
	(72134, 'DEMULAKA', 5759, 418, 32),
	(72135, 'KURULIMBU', 5759, 418, 32),
	(72136, 'KURULIMBU SELATAN', 5759, 418, 32),
	(72137, 'NGALUPOLO/NAGLUPOLO', 5759, 418, 32),
	(72138, 'NGGUWA', 5759, 418, 32),
	(72139, 'ROGA', 5759, 418, 32),
	(72140, 'SOKORIA', 5759, 418, 32),
	(72141, 'AEBARA', 5760, 418, 32),
	(72142, 'KELISAMBA', 5760, 418, 32),
	(72143, 'LUNGARIA', 5760, 418, 32),
	(72144, 'MAUBASA', 5760, 418, 32),
	(72145, 'MAUBASA BARAT', 5760, 418, 32),
	(72146, 'MAUBASA TIMUR', 5760, 418, 32),
	(72147, 'MOLE', 5760, 418, 32),
	(72148, 'RATEMANGGA', 5760, 418, 32),
	(72149, 'SERANDORI', 5760, 418, 32),
	(72150, 'WONDA', 5760, 418, 32),
	(72151, 'AEJETI', 5761, 418, 32),
	(72152, 'KAZO KAPO', 5761, 418, 32),
	(72153, 'NDORIWOI', 5761, 418, 32),
	(72154, 'PADERAPE', 5761, 418, 32),
	(72155, 'PUUTARA', 5761, 418, 32),
	(72156, 'REDODORI', 5761, 418, 32),
	(72157, 'RENDO RATERUA', 5761, 418, 32),
	(72158, 'RENGA MENGE', 5761, 418, 32),
	(72159, 'RORURANGGA (RENDO RANGA)', 5761, 418, 32),
	(72160, 'AE NDOKO', 5762, 418, 32),
	(72161, 'AELIPO', 5762, 418, 32),
	(72162, 'AELIPO I', 5762, 418, 32),
	(72163, 'AEMURI', 5762, 418, 32),
	(72164, 'DETUBELA', 5762, 418, 32),
	(72165, 'EKOAE', 5762, 418, 32),
	(72166, 'FATAATU', 5762, 418, 32),
	(72167, 'FATAATU TIMUR', 5762, 418, 32),
	(72168, 'KELITEMBU', 5762, 418, 32),
	(72169, 'MAUTENDA', 5762, 418, 32),
	(72170, 'MAUTENDA BARAT', 5762, 418, 32),
	(72171, 'MBOTULAKA', 5762, 418, 32),
	(72172, 'MUKUSAKI', 5762, 418, 32),
	(72173, 'NUANGENDA', 5762, 418, 32),
	(72174, 'NUMBA', 5762, 418, 32),
	(72175, 'RATEWATI', 5762, 418, 32),
	(72176, 'RATEWATI SELATAN', 5762, 418, 32),
	(72177, 'TANALI', 5762, 418, 32),
	(72178, 'WAKA', 5762, 418, 32),
	(72179, 'WELAMOSA', 5762, 418, 32),
	(72180, 'WEWARIA', 5762, 418, 32),
	(72181, 'WOLOOJA', 5762, 418, 32),
	(72182, 'NGGELA', 5763, 418, 32),
	(72183, 'NUAMULU', 5763, 418, 32),
	(72184, 'PORA', 5763, 418, 32),
	(72185, 'TENDA', 5763, 418, 32),
	(72186, 'WIWIPEMO', 5763, 418, 32),
	(72187, 'WOLOJITA', 5763, 418, 32),
	(72188, 'BOKASAPE', 5764, 418, 32),
	(72189, 'BOKASAPE TIMUR', 5764, 418, 32),
	(72190, 'JOPU', 5764, 418, 32),
	(72191, 'LIKANAKA', 5764, 418, 32),
	(72192, 'LISE LOWOBORA', 5764, 418, 32),
	(72193, 'LISE PU\\U', 5764, 418, 32),
	(72194, 'LISEDETU', 5764, 418, 32),
	(72195, 'MBULI LOO', 5764, 418, 32),
	(72196, 'MBULIWARALAU', 5764, 418, 32),
	(72197, 'MBULIWARALAU UTARA', 5764, 418, 32),
	(72198, 'NAKAMBARA', 5764, 418, 32),
	(72199, 'NIROMESI', 5764, 418, 32),
	(72200, 'NUALISE', 5764, 418, 32),
	(72201, 'RINDIWAWO', 5764, 418, 32),
	(72202, 'TANA LO\\O', 5764, 418, 32),
	(72203, 'WOLOKOLI', 5764, 418, 32),
	(72204, 'WOLOSOKO', 5764, 418, 32),
	(72205, 'ADONARA', 5765, 419, 32),
	(72206, 'KOLILANANG', 5765, 419, 32),
	(72207, 'KOLIMASANG', 5765, 419, 32),
	(72208, 'KOLIPETUNG', 5765, 419, 32),
	(72209, 'LAMAHODA', 5765, 419, 32),
	(72210, 'NISANULAN', 5765, 419, 32),
	(72211, 'SAGU', 5765, 419, 32),
	(72212, 'TIKATUKANG', 5765, 419, 32),
	(72213, 'BUGALIMA', 5766, 419, 32),
	(72214, 'BUKIT SEBURI I', 5766, 419, 32),
	(72215, 'BUKIT SEBURI II', 5766, 419, 32),
	(72216, 'DANIBAO', 5766, 419, 32),
	(72217, 'DUANUR (DUWANUR)', 5766, 419, 32),
	(72218, 'HOMA', 5766, 419, 32),
	(72219, 'HURUNG', 5766, 419, 32),
	(72220, 'ILEPATI', 5766, 419, 32),
	(72221, 'KIMA KAMAK', 5766, 419, 32),
	(72222, 'NIMUN DANI BAO', 5766, 419, 32),
	(72223, 'PAJINIAN', 5766, 419, 32),
	(72224, 'RIANGPADU', 5766, 419, 32),
	(72225, 'TONUWOTAN', 5766, 419, 32),
	(72226, 'WAI TUKAN', 5766, 419, 32),
	(72227, 'WAIWADAN', 5766, 419, 32),
	(72228, 'WATO BAYA', 5766, 419, 32),
	(72229, 'WOLOKIBANG', 5766, 419, 32),
	(72230, 'WUREH', 5766, 419, 32),
	(72231, 'BAYA', 5767, 419, 32),
	(72232, 'BIDARA', 5767, 419, 32),
	(72233, 'HOKO HOROWURA', 5767, 419, 32),
	(72234, 'HOROWURA', 5767, 419, 32),
	(72235, 'KENOTAN', 5767, 419, 32),
	(72236, 'KOKOTOBO', 5767, 419, 32),
	(72237, 'LEWOBELE', 5767, 419, 32),
	(72238, 'LEWOPAO', 5767, 419, 32),
	(72239, 'LITE', 5767, 419, 32),
	(72240, 'NUBALEMA', 5767, 419, 32),
	(72241, 'NUBALEMA DUA', 5767, 419, 32),
	(72242, 'OE SAYANG', 5767, 419, 32),
	(72243, 'WEWIT', 5767, 419, 32),
	(72244, 'BELOTA', 5768, 419, 32),
	(72245, 'BILAL', 5768, 419, 32),
	(72246, 'DAWATAA', 5768, 419, 32),
	(72247, 'GELONG', 5768, 419, 32),
	(72248, 'IPIEBANG', 5768, 419, 32),
	(72249, 'KARINGLAMALOUK', 5768, 419, 32),
	(72250, 'KIWANG ONA', 5768, 419, 32),
	(72251, 'KWAELAGA LAMAWATO', 5768, 419, 32),
	(72252, 'LAMAHALA JAYA', 5768, 419, 32),
	(72253, 'LAMALATA', 5768, 419, 32),
	(72254, 'LAMATEWELU', 5768, 419, 32),
	(72255, 'LEWOBUNGA', 5768, 419, 32),
	(72256, 'LIBUBARU', 5768, 419, 32),
	(72257, 'NARASAOSINA', 5768, 419, 32),
	(72258, 'PUHU', 5768, 419, 32),
	(72259, 'TAPOBALI', 5768, 419, 32),
	(72260, 'TERONG', 5768, 419, 32),
	(72261, 'TUAWOLO', 5768, 419, 32),
	(72262, 'WAIBURAK', 5768, 419, 32),
	(72263, 'WAIWERANG KOTA', 5768, 419, 32),
	(72264, 'WATANPAO', 5768, 419, 32),
	(72265, 'BAMA', 5769, 419, 32),
	(72266, 'BLEPANAWA', 5769, 419, 32),
	(72267, 'KAWALELO', 5769, 419, 32),
	(72268, 'LAMIKA', 5769, 419, 32),
	(72269, 'LEWOKLUOK', 5769, 419, 32),
	(72270, 'LEWOMUDA', 5769, 419, 32),
	(72271, 'WATOTIKA ILE', 5769, 419, 32),
	(72272, 'BAYUNTAA', 5770, 419, 32),
	(72273, 'BEDALEWUN', 5770, 419, 32),
	(72274, 'BOLENG', 5770, 419, 32),
	(72275, 'BUNGALAWAN', 5770, 419, 32),
	(72276, 'DOKENG', 5770, 419, 32),
	(72277, 'DUABLOLONG', 5770, 419, 32),
	(72278, 'HARUBALA', 5770, 419, 32),
	(72279, 'HELANLANGOWUYO', 5770, 419, 32),
	(72280, 'LAMABAYUNG', 5770, 419, 32),
	(72281, 'LAMAWOLO', 5770, 419, 32),
	(72282, 'LEBANUBA', 5770, 419, 32),
	(72283, 'LEWAT', 5770, 419, 32),
	(72284, 'LEWO KELENG', 5770, 419, 32),
	(72285, 'LEWOPAO', 5770, 419, 32),
	(72286, 'NELE LAMAWANGI', 5770, 419, 32),
	(72287, 'NELEBLOLONG', 5770, 419, 32),
	(72288, 'NELELAMADIKEN', 5770, 419, 32),
	(72289, 'NELELAMAWANGI DUA', 5770, 419, 32),
	(72290, 'NELERERENG', 5770, 419, 32),
	(72291, 'NOBO', 5770, 419, 32),
	(72292, 'RIAWALE/RIANWALE', 5770, 419, 32),
	(72293, 'BIRAWAN', 5771, 419, 32),
	(72294, 'DULIPALI', 5771, 419, 32),
	(72295, 'LEWOAWANG', 5771, 419, 32),
	(72296, 'NOBOKONGA', 5771, 419, 32),
	(72297, 'NURRI', 5771, 419, 32),
	(72298, 'RIANG RITA', 5771, 419, 32),
	(72299, 'RIANGBURA', 5771, 419, 32),
	(72300, 'HALAKODANUAN', 5772, 419, 32),
	(72301, 'LEWOHALA', 5772, 419, 32),
	(72302, 'LEWOLOBA', 5772, 419, 32),
	(72303, 'MUDAKEPUTU', 5772, 419, 32),
	(72304, 'RIANG KEMIE', 5772, 419, 32),
	(72305, 'TIWATOBI', 5772, 419, 32),
	(72306, 'WAILOLONG', 5772, 419, 32),
	(72307, 'WATOTUTU', 5772, 419, 32),
	(72308, 'ADOBALA', 5773, 419, 32),
	(72309, 'HINGA', 5773, 419, 32),
	(72310, 'HORINARA', 5773, 419, 32),
	(72311, 'KELUWAIN', 5773, 419, 32),
	(72312, 'LAMABUNGA', 5773, 419, 32),
	(72313, 'LAMAPAHA', 5773, 419, 32),
	(72314, 'MANGAALENG', 5773, 419, 32),
	(72315, 'MUDA', 5773, 419, 32),
	(72316, 'NISAKARANG', 5773, 419, 32),
	(72317, 'PEPAK KELU', 5773, 419, 32),
	(72318, 'REDONTENA', 5773, 419, 32),
	(72319, 'SUKUTOKAN', 5773, 419, 32),
	(72320, 'AMAGARAPATI', 5774, 419, 32),
	(72321, 'BALELA', 5774, 419, 32),
	(72322, 'EKASAPTA', 5774, 419, 32),
	(72323, 'LAMAWALANG', 5774, 419, 32),
	(72324, 'LARANTUKA', 5774, 419, 32),
	(72325, 'LEWOLERE', 5774, 419, 32),
	(72326, 'LOHAYONG', 5774, 419, 32),
	(72327, 'LOKEA', 5774, 419, 32),
	(72328, 'MOKANTARAK', 5774, 419, 32),
	(72329, 'PANTAI BESAR', 5774, 419, 32),
	(72330, 'POHON BAO', 5774, 419, 32),
	(72331, 'POHON SIRIH', 5774, 419, 32),
	(72332, 'POSTOH', 5774, 419, 32),
	(72333, 'PUKEN TOBI WANGI BAO', 5774, 419, 32),
	(72334, 'SAROTARI', 5774, 419, 32),
	(72335, 'SAROTARI TENGAH', 5774, 419, 32),
	(72336, 'SAROTARI TIMUR', 5774, 419, 32),
	(72337, 'WAIBALUN', 5774, 419, 32),
	(72338, 'WAIHALI', 5774, 419, 32),
	(72339, 'WERI', 5774, 419, 32),
	(72340, 'BALUKHERING', 5775, 419, 32),
	(72341, 'BANTALA', 5775, 419, 32),
	(72342, 'ILE PADUNG', 5775, 419, 32),
	(72343, 'LEWOBELEN', 5775, 419, 32),
	(72344, 'PAINAPANG', 5775, 419, 32),
	(72345, 'RIANGKOTEK', 5775, 419, 32),
	(72346, 'SINAR HADING', 5775, 419, 32),
	(72347, 'BALAWELING I', 5776, 419, 32),
	(72348, 'BALAWELING II', 5776, 419, 32),
	(72349, 'DANI WATO', 5776, 419, 32),
	(72350, 'KALELU', 5776, 419, 32),
	(72351, 'KARAWATUNG', 5776, 419, 32),
	(72352, 'LAMAOLE', 5776, 419, 32),
	(72353, 'LAMAWALANG (LAMAWOHONG)', 5776, 419, 32),
	(72354, 'LEWONAMA', 5776, 419, 32),
	(72355, 'LEWOTANAH OLE', 5776, 419, 32),
	(72356, 'NUSADANI', 5776, 419, 32),
	(72357, 'ONGALERENG', 5776, 419, 32),
	(72358, 'PAMAKAYO', 5776, 419, 32),
	(72359, 'RITAEBANG', 5776, 419, 32),
	(72360, 'TANAH LEIN', 5776, 419, 32),
	(72361, 'TITEHENA', 5776, 419, 32),
	(72362, 'BUBU ATAGAMU', 5777, 419, 32),
	(72363, 'KALIKE / KELIKE', 5777, 419, 32),
	(72364, 'KALIKE / KELIKE AIMATAN', 5777, 419, 32),
	(72365, 'KENERE', 5777, 419, 32),
	(72366, 'LEMANU', 5777, 419, 32),
	(72367, 'LEWOGRARAN', 5777, 419, 32),
	(72368, 'SULENGWASENG', 5777, 419, 32),
	(72369, 'KWUTA (KAWUTA)', 5778, 419, 32),
	(72370, 'LABELEN', 5778, 419, 32),
	(72371, 'LAMAWAI', 5778, 419, 32),
	(72372, 'LEBAO', 5778, 419, 32),
	(72373, 'LEWOGEKA', 5778, 419, 32),
	(72374, 'LEWOHEDO', 5778, 419, 32),
	(72375, 'LIWO', 5778, 419, 32),
	(72376, 'LOHAYONG I', 5778, 419, 32),
	(72377, 'LOHAYONG II', 5778, 419, 32),
	(72378, 'MENANGA', 5778, 419, 32),
	(72379, 'MOTONWUTUN', 5778, 419, 32),
	(72380, 'TANAH WERANG', 5778, 419, 32),
	(72381, 'WATANHURA', 5778, 419, 32),
	(72382, 'WATANHURA II', 5778, 419, 32),
	(72383, 'WATOBUKU', 5778, 419, 32),
	(72384, 'WATOHARI', 5778, 419, 32),
	(72385, 'WULUBLOLONG', 5778, 419, 32),
	(72386, 'ARANSINA', 5779, 419, 32),
	(72387, 'BAHINGA', 5779, 419, 32),
	(72388, 'BANDONA', 5779, 419, 32),
	(72389, 'GEKENGDERAN', 5779, 419, 32),
	(72390, 'KOLAKA', 5779, 419, 32),
	(72391, 'LAMANABI', 5779, 419, 32),
	(72392, 'LAMATUTU', 5779, 419, 32),
	(72393, 'LATON LIWO', 5779, 419, 32),
	(72394, 'LATONLIWO DUA', 5779, 419, 32),
	(72395, 'LEWOBUNGA', 5779, 419, 32),
	(72396, 'NUSANIPA', 5779, 419, 32),
	(72397, 'PATISIRAWALANG', 5779, 419, 32),
	(72398, 'RATU LODONG', 5779, 419, 32),
	(72399, 'SINA HADIGALA', 5779, 419, 32),
	(72400, 'SINAMALAKA', 5779, 419, 32),
	(72401, 'WAIBAO', 5779, 419, 32),
	(72402, 'ADABANG', 5780, 419, 32),
	(72403, 'BOKANG WOLOMATANG', 5780, 419, 32),
	(72404, 'DULI JAYA', 5780, 419, 32),
	(72405, 'DUN TANA LEWOINGU', 5780, 419, 32),
	(72406, 'ILE GERONG', 5780, 419, 32),
	(72407, 'KOBASOMA', 5780, 419, 32),
	(72408, 'KONGA', 5780, 419, 32),
	(72409, 'LERABOLENG', 5780, 419, 32),
	(72410, 'LEWOINGU', 5780, 419, 32),
	(72411, 'LEWOLAGA', 5780, 419, 32),
	(72412, 'SERINUHO', 5780, 419, 32),
	(72413, 'TENAWAHANG', 5780, 419, 32),
	(72414, 'TUAKEPA', 5780, 419, 32),
	(72415, 'WATOWARA', 5780, 419, 32),
	(72416, 'BALAWELING', 5781, 419, 32),
	(72417, 'BALAWELING NOTEN', 5781, 419, 32),
	(72418, 'BAOBAGE', 5781, 419, 32),
	(72419, 'LAMABLAWA', 5781, 419, 32),
	(72420, 'LAMALEKA', 5781, 419, 32),
	(72421, 'LEWOPULO', 5781, 419, 32),
	(72422, 'ORINGBELE', 5781, 419, 32),
	(72423, 'PLEDO', 5781, 419, 32),
	(72424, 'RIANGDULI', 5781, 419, 32),
	(72425, 'SANDOSI', 5781, 419, 32),
	(72426, 'TOBITIKA', 5781, 419, 32),
	(72427, 'TUWAGOETOBI', 5781, 419, 32),
	(72428, 'WAIWURING', 5781, 419, 32),
	(72429, 'WATOLOLONG', 5781, 419, 32),
	(72430, 'WATOONE', 5781, 419, 32),
	(72431, 'WERANGGERE', 5781, 419, 32),
	(72432, 'BELIKO (BLIKO)', 5782, 419, 32),
	(72433, 'KAWELA', 5782, 419, 32),
	(72434, 'KLUKENGNUKING', 5782, 419, 32),
	(72435, 'NAYUBAYA', 5782, 419, 32),
	(72436, 'OYANGBARANG', 5782, 419, 32),
	(72437, 'PANDAI', 5782, 419, 32),
	(72438, 'SAMASOGE', 5782, 419, 32),
	(72439, 'TANAH TUKAN', 5782, 419, 32),
	(72440, 'TOBILOTA', 5782, 419, 32),
	(72441, 'WAILEBE', 5782, 419, 32),
	(72442, 'WOTANULUMADO', 5782, 419, 32),
	(72443, 'BAYUNTAA', 5783, 419, 32),
	(72444, 'BORU', 5783, 419, 32),
	(72445, 'BORU KEDANG', 5783, 419, 32),
	(72446, 'DOKENG', 5783, 419, 32),
	(72447, 'HEWA', 5783, 419, 32),
	(72448, 'HOKENG JAYA', 5783, 419, 32),
	(72449, 'KLATANLO', 5783, 419, 32),
	(72450, 'LEWAT', 5783, 419, 32),
	(72451, 'NAWOKOTE', 5783, 419, 32),
	(72452, 'NILEKNOHENG', 5783, 419, 32),
	(72453, 'NOBO', 5783, 419, 32),
	(72454, 'OJANDETUN', 5783, 419, 32),
	(72455, 'PANTAI OA', 5783, 419, 32),
	(72456, 'PULULERA', 5783, 419, 32),
	(72457, 'WAIULA', 5783, 419, 32),
	(72458, 'ALAK', 5784, 420, 32),
	(72459, 'BATUPLAT', 5784, 420, 32),
	(72460, 'FATUFETO', 5784, 420, 32),
	(72461, 'MANTASI', 5784, 420, 32),
	(72462, 'MANULAI II', 5784, 420, 32),
	(72463, 'MANUTAPEN', 5784, 420, 32),
	(72464, 'NAIONI', 5784, 420, 32),
	(72465, 'NAMOSAIN', 5784, 420, 32),
	(72466, 'NUNBAUN DELHA', 5784, 420, 32),
	(72467, 'NUNBAUN SABU', 5784, 420, 32),
	(72468, 'NUNHILA', 5784, 420, 32),
	(72469, 'PENKASE OELETA', 5784, 420, 32),
	(72470, 'FATUKANUTU (FATUKNUTU)', 5785, 420, 32),
	(72471, 'FATUTETA', 5785, 420, 32),
	(72472, 'KAIRANE', 5785, 420, 32),
	(72473, 'KUANHEUM', 5785, 420, 32),
	(72474, 'NIUNBAUN', 5785, 420, 32),
	(72475, 'OELETO (OEFETO)', 5785, 420, 32),
	(72476, 'RAKNAMO', 5785, 420, 32),
	(72477, 'ENOLANAN', 5786, 420, 32),
	(72478, 'MUKE', 5786, 420, 32),
	(72479, 'NUNMAFO', 5786, 420, 32),
	(72480, 'OEMOFA', 5786, 420, 32),
	(72481, 'OEMOLO', 5786, 420, 32),
	(72482, 'OENAUNU', 5786, 420, 32),
	(72483, 'OENIKO', 5786, 420, 32),
	(72484, 'OENUNTONO (OENUNUTONO)', 5786, 420, 32),
	(72485, 'PATHAU', 5786, 420, 32),
	(72486, 'SEKI', 5786, 420, 32),
	(72487, 'APREN', 5787, 420, 32),
	(72488, 'KOTABES', 5787, 420, 32),
	(72489, 'NONBES', 5787, 420, 32),
	(72490, 'OENONI', 5787, 420, 32),
	(72491, 'OENONI DUA', 5787, 420, 32),
	(72492, 'OESENA', 5787, 420, 32),
	(72493, 'PONAIN', 5787, 420, 32),
	(72494, 'TESBATAN', 5787, 420, 32),
	(72495, 'TESBATAN II', 5787, 420, 32),
	(72496, 'ERBAUN', 5788, 420, 32),
	(72497, 'MERBAUN', 5788, 420, 32),
	(72498, 'NEKBAUN', 5788, 420, 32),
	(72499, 'NIUKBAUN', 5788, 420, 32),
	(72500, 'SOBA', 5788, 420, 32),
	(72501, 'TEUNBAUN', 5788, 420, 32),
	(72502, 'TOOBAUN', 5788, 420, 32),
	(72503, 'TUNBAUN', 5788, 420, 32),
	(72504, 'BURAEN', 5789, 420, 32),
	(72505, 'NEKMESE', 5789, 420, 32),
	(72506, 'RETRAEN', 5789, 420, 32),
	(72507, 'SAHRAEN', 5789, 420, 32),
	(72508, 'SONRAEN', 5789, 420, 32),
	(72509, 'ENORAEN', 5790, 420, 32),
	(72510, 'OEBESI', 5790, 420, 32),
	(72511, 'PAKUBAUN', 5790, 420, 32),
	(72512, 'RABEKA', 5790, 420, 32),
	(72513, 'BIOBA BARU (BARUTAEN)', 5791, 420, 32),
	(72514, 'LETKOLE', 5791, 420, 32),
	(72515, 'MANUBELON', 5791, 420, 32),
	(72516, 'NEFONEUT', 5791, 420, 32),
	(72517, 'FAUMES', 5792, 420, 32),
	(72518, 'HONUK', 5792, 420, 32),
	(72519, 'OELFATU', 5792, 420, 32),
	(72520, 'SAUKIBE', 5792, 420, 32),
	(72521, 'SOLIU', 5792, 420, 32),
	(72522, 'TIMAU', 5792, 420, 32),
	(72523, 'FATUMETAN', 5793, 420, 32),
	(72524, 'FATUSUKI', 5793, 420, 32),
	(72525, 'LELOBOKO', 5793, 420, 32),
	(72526, 'LELOGAMA', 5793, 420, 32),
	(72527, 'OELBANU', 5793, 420, 32),
	(72528, 'OHAEM', 5793, 420, 32),
	(72529, 'OHAEM DUA', 5793, 420, 32),
	(72530, 'BINAFUN', 5794, 420, 32),
	(72531, 'BITOBE', 5794, 420, 32),
	(72532, 'BONMUTI', 5794, 420, 32),
	(72533, 'FATUMONAS', 5794, 420, 32),
	(72534, 'KIFU', 5795, 420, 32),
	(72535, 'NATEMNANU (NATUMMANU)', 5795, 420, 32),
	(72536, 'NETEMNANU SELATAN', 5795, 420, 32),
	(72537, 'NETEMNANU UTARA', 5795, 420, 32),
	(72538, 'NUNUANAH', 5795, 420, 32),
	(72539, 'AFOAN', 5796, 420, 32),
	(72540, 'BAKUIN', 5796, 420, 32),
	(72541, 'FATUNAUS', 5796, 420, 32),
	(72542, 'KOLABE', 5796, 420, 32),
	(72543, 'LILMUS', 5796, 420, 32),
	(72544, 'NAIKLIU', 5796, 420, 32),
	(72545, 'CAMPLONG I', 5797, 420, 32),
	(72546, 'CAMPLONG II', 5797, 420, 32),
	(72547, 'EKATETA', 5797, 420, 32),
	(72548, 'KIUONI', 5797, 420, 32),
	(72549, 'KUIMASI', 5797, 420, 32),
	(72550, 'NAUNU', 5797, 420, 32),
	(72551, 'OEBOLA', 5797, 420, 32),
	(72552, 'OEBOLA DALAM', 5797, 420, 32),
	(72553, 'SILLU', 5797, 420, 32),
	(72554, 'TULNAKU', 5797, 420, 32),
	(72555, 'KALALI', 5798, 420, 32),
	(72556, 'NAITAE', 5798, 420, 32),
	(72557, 'NUATAUS', 5798, 420, 32),
	(72558, 'POTO', 5798, 420, 32),
	(72559, 'TUAKAU', 5798, 420, 32),
	(72560, 'NONBAUN', 5799, 420, 32),
	(72561, 'NUNSAEN', 5799, 420, 32),
	(72562, 'OELBITENO', 5799, 420, 32),
	(72563, 'PASSI', 5799, 420, 32),
	(72564, 'KELAPA LIMA', 5800, 420, 32),
	(72565, 'LASIANA', 5800, 420, 32),
	(72566, 'OESAPA', 5800, 420, 32),
	(72567, 'OESAPA BARAT', 5800, 420, 32),
	(72568, 'OESAPA SELATAN', 5800, 420, 32),
	(72569, 'AIR MATA', 5801, 420, 32),
	(72570, 'BONIPOI', 5801, 420, 32),
	(72571, 'FATUBESI', 5801, 420, 32),
	(72572, 'LAI LAI BISI KOPAN', 5801, 420, 32),
	(72573, 'MERDEKA', 5801, 420, 32),
	(72574, 'NEFONAEK', 5801, 420, 32),
	(72575, 'OEBA', 5801, 420, 32),
	(72576, 'PASIR PANJANG', 5801, 420, 32),
	(72577, 'SOLOR', 5801, 420, 32),
	(72578, 'TODE KISAR', 5801, 420, 32),
	(72579, 'AIR NONA', 5802, 420, 32),
	(72580, 'BAKUNASE', 5802, 420, 32),
	(72581, 'FONTEIN', 5802, 420, 32),
	(72582, 'KUANINO', 5802, 420, 32),
	(72583, 'NAIKOTEN I (SATU)', 5802, 420, 32),
	(72584, 'NAIKOTEN II (DUA)', 5802, 420, 32),
	(72585, 'NUNLEU', 5802, 420, 32),
	(72586, 'BATAKTE', 5803, 420, 32),
	(72587, 'BOLOK', 5803, 420, 32),
	(72588, 'KUANHEUN', 5803, 420, 32),
	(72589, 'LIFULEO', 5803, 420, 32),
	(72590, 'MANULAI I', 5803, 420, 32),
	(72591, 'NITNEO', 5803, 420, 32),
	(72592, 'OEMATANUNU', 5803, 420, 32),
	(72593, 'OENAEK', 5803, 420, 32),
	(72594, 'OENESU', 5803, 420, 32),
	(72595, 'SUMLILI', 5803, 420, 32),
	(72596, 'TABLOLONG', 5803, 420, 32),
	(72597, 'TESABELA', 5803, 420, 32),
	(72598, 'MATA AIR', 5804, 420, 32),
	(72599, 'NOELBAKI', 5804, 420, 32),
	(72600, 'OEBELO', 5804, 420, 32),
	(72601, 'OELNASI', 5804, 420, 32),
	(72602, 'OELPUAH', 5804, 420, 32),
	(72603, 'PENFUI TIMUR', 5804, 420, 32),
	(72604, 'TANAH MERAH', 5804, 420, 32),
	(72605, 'TARUS', 5804, 420, 32),
	(72606, 'BABAU', 5805, 420, 32),
	(72607, 'MANUSAK', 5805, 420, 32),
	(72608, 'MERDEKA', 5805, 420, 32),
	(72609, 'NAIBONAT', 5805, 420, 32),
	(72610, 'NUNKURUS', 5805, 420, 32),
	(72611, 'OEFAFI', 5805, 420, 32),
	(72612, 'OELATIMO', 5805, 420, 32),
	(72613, 'OESAO', 5805, 420, 32),
	(72614, 'OESAO II', 5805, 420, 32),
	(72615, 'PUKDALE', 5805, 420, 32),
	(72616, 'TANAH PUTIH', 5805, 420, 32),
	(72617, 'TUAPUKAN', 5805, 420, 32),
	(72618, 'TUATUKA', 5805, 420, 32),
	(72619, 'BELLO', 5806, 420, 32),
	(72620, 'FATUKOA', 5806, 420, 32),
	(72621, 'KOLHUA', 5806, 420, 32),
	(72622, 'MAULAFA', 5806, 420, 32),
	(72623, 'NAIKOLAN', 5806, 420, 32),
	(72624, 'NAIMATA', 5806, 420, 32),
	(72625, 'OEPURA', 5806, 420, 32),
	(72626, 'PENFUI', 5806, 420, 32),
	(72627, 'SIKUMANA', 5806, 420, 32),
	(72628, 'BISMARK (BISMARAK)', 5807, 420, 32),
	(72629, 'BONE', 5807, 420, 32),
	(72630, 'OBEN', 5807, 420, 32),
	(72631, 'OELOMIN', 5807, 420, 32),
	(72632, 'OEMASI', 5807, 420, 32),
	(72633, 'OENIF', 5807, 420, 32),
	(72634, 'OEPAHA', 5807, 420, 32),
	(72635, 'TALOITAN (TALOETAN)', 5807, 420, 32),
	(72636, 'TASIKONA', 5807, 420, 32),
	(72637, 'TUNFEU', 5807, 420, 32),
	(72638, 'USAPISONBAI', 5807, 420, 32),
	(72639, 'BAKUNASE DUA', 5808, 420, 32),
	(72640, 'FATULULI', 5808, 420, 32),
	(72641, 'KAYU PUTIH', 5808, 420, 32),
	(72642, 'LILIBA', 5808, 420, 32),
	(72643, 'OEBOBO', 5808, 420, 32),
	(72644, 'OEBUFU', 5808, 420, 32),
	(72645, 'OETETE', 5808, 420, 32),
	(72646, 'TUAK DAUN MERAH', 5808, 420, 32),
	(72647, 'BATUINAN', 5809, 420, 32),
	(72648, 'BOKONUSAN', 5809, 420, 32),
	(72649, 'HANSISI', 5809, 420, 32),
	(72650, 'HUILELOT', 5809, 420, 32),
	(72651, 'LETBAUN', 5809, 420, 32),
	(72652, 'OTAN', 5809, 420, 32),
	(72653, 'UIASA', 5809, 420, 32),
	(72654, 'UITAO', 5809, 420, 32),
	(72655, 'AKLE', 5810, 420, 32),
	(72656, 'NAIKEN (NAIKEAN)', 5810, 420, 32),
	(72657, 'ONANSILA (OENANSILA)', 5810, 420, 32),
	(72658, 'UIBOA', 5810, 420, 32),
	(72659, 'UITIUH ANA', 5810, 420, 32),
	(72660, 'UITIUH TUAN', 5810, 420, 32),
	(72661, 'BIPOLO', 5811, 420, 32),
	(72662, 'OETETA', 5811, 420, 32),
	(72663, 'PANTAI BERINGIN', 5811, 420, 32),
	(72664, 'PANTULAN', 5811, 420, 32),
	(72665, 'PARITI', 5811, 420, 32),
	(72666, 'PITAI', 5811, 420, 32),
	(72667, 'SULAMU', 5811, 420, 32),
	(72668, 'BAUMATA', 5812, 420, 32),
	(72669, 'BAUMATA BARAT', 5812, 420, 32),
	(72670, 'BAUMATA TIMUR', 5812, 420, 32),
	(72671, 'BAUMATA UTARA', 5812, 420, 32),
	(72672, 'BOKONG', 5812, 420, 32),
	(72673, 'KUAKLALO', 5812, 420, 32),
	(72674, 'OELETSALA', 5812, 420, 32),
	(72675, 'OELTUA', 5812, 420, 32),
	(72676, 'BENU', 5813, 420, 32),
	(72677, 'FATUKONA', 5813, 420, 32),
	(72678, 'HUEKNUTU', 5813, 420, 32),
	(72679, 'KAUNIKI', 5813, 420, 32),
	(72680, 'NOELMINA', 5813, 420, 32),
	(72681, 'OELNAINENO', 5813, 420, 32),
	(72682, 'OESUSU', 5813, 420, 32),
	(72683, 'TAKARI', 5813, 420, 32),
	(72684, 'TANINI', 5813, 420, 32),
	(72685, 'TUAPANAF', 5813, 420, 32),
	(72686, 'ATAKORE', 5814, 421, 32),
	(72687, 'DORI PEWUT', 5814, 421, 32),
	(72688, 'DULIR', 5814, 421, 32),
	(72689, 'ILE KERBAU', 5814, 421, 32),
	(72690, 'ILE KIMOK', 5814, 421, 32),
	(72691, 'KATAKEJA', 5814, 421, 32),
	(72692, 'LEBAATA', 5814, 421, 32),
	(72693, 'LEREK', 5814, 421, 32),
	(72694, 'LEWOGROMA', 5814, 421, 32),
	(72695, 'LUSILAME', 5814, 421, 32),
	(72696, 'NOGO DONI', 5814, 421, 32),
	(72697, 'NUBA ATALOJO', 5814, 421, 32),
	(72698, 'NUBABOLI', 5814, 421, 32),
	(72699, 'NUBAHAERAKA', 5814, 421, 32),
	(72700, 'TUBUK RAJAN', 5814, 421, 32),
	(72701, 'ATU WA\\LUPANG', 5815, 421, 32),
	(72702, 'ATULALENG', 5815, 421, 32),
	(72703, 'BARENG', 5815, 421, 32),
	(72704, 'BEAN', 5815, 421, 32),
	(72705, 'BENIHADING', 5815, 421, 32),
	(72706, 'BENIHADING II', 5815, 421, 32),
	(72707, 'BURIWUTUNG', 5815, 421, 32),
	(72708, 'KALIKUR', 5815, 421, 32),
	(72709, 'KALIKUR WL', 5815, 421, 32),
	(72710, 'KAOHUA', 5815, 421, 32),
	(72711, 'LEUBURI', 5815, 421, 32),
	(72712, 'LEUWOHUNG', 5815, 421, 32),
	(72713, 'LOYOBOHOR', 5815, 421, 32),
	(72714, 'MAMPIR', 5815, 421, 32),
	(72715, 'PANAMA', 5815, 421, 32),
	(72716, 'ROHO', 5815, 421, 32),
	(72717, 'RUMANG', 5815, 421, 32),
	(72718, 'TOBOTANI', 5815, 421, 32),
	(72719, 'TUBUNG WALANG', 5815, 421, 32),
	(72720, 'UMALEU', 5815, 421, 32),
	(72721, 'AMAKAKA', 5816, 421, 32),
	(72722, 'BEUTARAN', 5816, 421, 32),
	(72723, 'BUNGAMUDA', 5816, 421, 32),
	(72724, 'DULITUKAN', 5816, 421, 32),
	(72725, 'KOLIPADAN', 5816, 421, 32),
	(72726, 'KOLONTOBO', 5816, 421, 32),
	(72727, 'LAMAWARA', 5816, 421, 32),
	(72728, 'LARANWUNTUN', 5816, 421, 32),
	(72729, 'MURUONA', 5816, 421, 32),
	(72730, 'NAPASABOK', 5816, 421, 32),
	(72731, 'PETUNTAWA', 5816, 421, 32),
	(72732, 'PLILOLON', 5816, 421, 32),
	(72733, 'TAGAWITI', 5816, 421, 32),
	(72734, 'TANJUNG BATU', 5816, 421, 32),
	(72735, 'WAOWALA', 5816, 421, 32),
	(72736, 'WATODIRI', 5816, 421, 32),
	(72737, 'AULESA', 5817, 421, 32),
	(72738, 'BAO LALI DULI', 5817, 421, 32),
	(72739, 'JONTONA', 5817, 421, 32),
	(72740, 'LAMAAU', 5817, 421, 32),
	(72741, 'LAMAGUTE', 5817, 421, 32),
	(72742, 'LAMATOKAN', 5817, 421, 32),
	(72743, 'LAMAWOLO', 5817, 421, 32),
	(72744, 'TODANARA', 5817, 421, 32),
	(72745, 'WAIMATAN', 5817, 421, 32),
	(72746, 'ATAKOWA', 5818, 421, 32),
	(72747, 'BANITOBO', 5818, 421, 32),
	(72748, 'BAOPAMA (BAOPANA)', 5818, 421, 32),
	(72749, 'BELOREBONG (BALUREBONG)', 5818, 421, 32),
	(72750, 'DIKESARE', 5818, 421, 32),
	(72751, 'HADAKEWA', 5818, 421, 32),
	(72752, 'LAMADALE', 5818, 421, 32),
	(72753, 'LAMALELA', 5818, 421, 32),
	(72754, 'LAMATUKA', 5818, 421, 32),
	(72755, 'LERAHINGA', 5818, 421, 32),
	(72756, 'LEWOELENG', 5818, 421, 32),
	(72757, 'LODOTODOKOWA', 5818, 421, 32),
	(72758, 'MERDEKA', 5818, 421, 32),
	(72759, 'SERANGGORANG', 5818, 421, 32),
	(72760, 'TAPOBARAN', 5818, 421, 32),
	(72761, 'TAPOLANGU', 5818, 421, 32),
	(72762, 'WAIENGA', 5818, 421, 32),
	(72763, 'ATAWAI', 5819, 421, 32),
	(72764, 'BABOKERONG', 5819, 421, 32),
	(72765, 'BAOBOLAK', 5819, 421, 32),
	(72766, 'BELABAJA', 5819, 421, 32),
	(72767, 'BOLI BEAN', 5819, 421, 32),
	(72768, 'DUAWUTUN', 5819, 421, 32),
	(72769, 'ILE BOLI', 5819, 421, 32),
	(72770, 'LABALIMUT', 5819, 421, 32),
	(72771, 'LOLONG', 5819, 421, 32),
	(72772, 'LUSIDUAWUTUN', 5819, 421, 32),
	(72773, 'PASIR PUTIH', 5819, 421, 32),
	(72774, 'PENIKENE', 5819, 421, 32),
	(72775, 'TEWAOWUTUNG', 5819, 421, 32),
	(72776, 'WUAKERONG', 5819, 421, 32),
	(72777, 'BAKALEREK', 5820, 421, 32),
	(72778, 'BAOLANGU', 5820, 421, 32),
	(72779, 'BELOBATANG', 5820, 421, 32),
	(72780, 'LEWOLEBA', 5820, 421, 32),
	(72781, 'LEWOLEBA BARAT', 5820, 421, 32),
	(72782, 'LEWOLEBA SELATAN', 5820, 421, 32),
	(72783, 'LEWOLEBA TENGAH', 5820, 421, 32),
	(72784, 'LEWOLEBA TIMUR', 5820, 421, 32),
	(72785, 'LEWOLEBA UTARA', 5820, 421, 32),
	(72786, 'LITE ULUMADO', 5820, 421, 32),
	(72787, 'NUBA MADO', 5820, 421, 32),
	(72788, 'PADA', 5820, 421, 32),
	(72789, 'PAUBOKOL (PAOBOKOL)', 5820, 421, 32),
	(72790, 'SELANDORO', 5820, 421, 32),
	(72791, 'UDAK MELOMATA', 5820, 421, 32),
	(72792, 'WAIJARANG', 5820, 421, 32),
	(72793, 'WATOKOBU', 5820, 421, 32),
	(72794, 'ARAMENGI', 5821, 421, 32),
	(72795, 'BALAURING', 5821, 421, 32),
	(72796, 'DOLULOLONG', 5821, 421, 32),
	(72797, 'HINGALAMAMENGI', 5821, 421, 32),
	(72798, 'HOELEA', 5821, 421, 32),
	(72799, 'HOELEA II', 5821, 421, 32),
	(72800, 'LEBEWALA', 5821, 421, 32),
	(72801, 'LEUBATANG', 5821, 421, 32),
	(72802, 'LEUDANUNG', 5821, 421, 32),
	(72803, 'LEUWAYANG', 5821, 421, 32),
	(72804, 'MAHAL', 5821, 421, 32),
	(72805, 'MAHAL II', 5821, 421, 32),
	(72806, 'MELUWITING', 5821, 421, 32),
	(72807, 'NILANAPO', 5821, 421, 32),
	(72808, 'NORMAL', 5821, 421, 32),
	(72809, 'NORMAL I', 5821, 421, 32),
	(72810, 'PEUSAWAH', 5821, 421, 32),
	(72811, 'ROMA', 5821, 421, 32),
	(72812, 'WAILOLONG', 5821, 421, 32),
	(72813, 'WALANGSAWAH', 5821, 421, 32),
	(72814, 'WOWONG', 5821, 421, 32),
	(72815, 'ALAP ATADEI', 5822, 421, 32),
	(72816, 'ATAILI', 5822, 421, 32),
	(72817, 'ATAKERA', 5822, 421, 32),
	(72818, 'BELOBAD (BELOBAO)', 5822, 421, 32),
	(72819, 'IMULOLONG', 5822, 421, 32),
	(72820, 'LAMALERA A', 5822, 421, 32),
	(72821, 'LAMALERA B', 5822, 421, 32),
	(72822, 'LELATA', 5822, 421, 32),
	(72823, 'LEWORAJA', 5822, 421, 32),
	(72824, 'PANTAI HARAPAN', 5822, 421, 32),
	(72825, 'POSIWATU', 5822, 421, 32),
	(72826, 'PUOR', 5822, 421, 32),
	(72827, 'PUOR B', 5822, 421, 32),
	(72828, 'TAPOBALI', 5822, 421, 32),
	(72829, 'WULANDONI', 5822, 421, 32),
	(72830, 'BARANG', 5823, 422, 32),
	(72831, 'BEA MESE', 5823, 422, 32),
	(72832, 'GAPONG', 5823, 422, 32),
	(72833, 'GOLO', 5823, 422, 32),
	(72834, 'GOLO NCUANG', 5823, 422, 32),
	(72835, 'KENTOL', 5823, 422, 32),
	(72836, 'LADUR', 5823, 422, 32),
	(72837, 'LANDO', 5823, 422, 32),
	(72838, 'LANGKAS', 5823, 422, 32),
	(72839, 'NENU', 5823, 422, 32),
	(72840, 'PAGAL', 5823, 422, 32),
	(72841, 'PERAK', 5823, 422, 32),
	(72842, 'PINGGANG', 5823, 422, 32),
	(72843, 'RADO', 5823, 422, 32),
	(72844, 'RIUNG', 5823, 422, 32),
	(72845, 'WELU', 5823, 422, 32),
	(72846, 'WUDI', 5823, 422, 32),
	(72847, 'BANGKA ARA', 5824, 422, 32),
	(72848, 'BERE', 5824, 422, 32),
	(72849, 'COMPANG CIBAL', 5824, 422, 32),
	(72850, 'GOLO LANAK', 5824, 422, 32),
	(72851, 'GOLO WOI', 5824, 422, 32),
	(72852, 'LATUNG', 5824, 422, 32),
	(72853, 'LENDA', 5824, 422, 32),
	(72854, 'TIMBU', 5824, 422, 32),
	(72855, 'WAE CODI', 5824, 422, 32),
	(72856, 'WAE RENCA', 5824, 422, 32),
	(72857, 'CAREP', 5825, 422, 32),
	(72858, 'GOLO DUKAL', 5825, 422, 32),
	(72859, 'KAROT', 5825, 422, 32),
	(72860, 'LAWIR', 5825, 422, 32),
	(72861, 'MBAU MUKU', 5825, 422, 32),
	(72862, 'PAU', 5825, 422, 32),
	(72863, 'PITAK', 5825, 422, 32),
	(72864, 'TENDA', 5825, 422, 32),
	(72865, 'WALI', 5825, 422, 32),
	(72866, 'WASO', 5825, 422, 32),
	(72867, 'WATU', 5825, 422, 32),
	(72868, 'BANGKA DESE', 5826, 422, 32),
	(72869, 'BANGKA LELAK', 5826, 422, 32),
	(72870, 'BANGKA TONGGUR', 5826, 422, 32),
	(72871, 'GELONG', 5826, 422, 32),
	(72872, 'KETANG', 5826, 422, 32),
	(72873, 'LENTANG', 5826, 422, 32),
	(72874, 'NATI', 5826, 422, 32),
	(72875, 'NDIWAR', 5826, 422, 32),
	(72876, 'PONG OMPU', 5826, 422, 32),
	(72877, 'URANG', 5826, 422, 32),
	(72878, 'BANGKA AJANG', 5827, 422, 32),
	(72879, 'BANGKA RUANG', 5827, 422, 32),
	(72880, 'BENTENG TUBI', 5827, 422, 32),
	(72881, 'BUAR', 5827, 422, 32),
	(72882, 'COMPANG DARI', 5827, 422, 32),
	(72883, 'DIMPONG', 5827, 422, 32),
	(72884, 'GOLO LANGKOK', 5827, 422, 32),
	(72885, 'LIANG BUA', 5827, 422, 32),
	(72886, 'MANONG', 5827, 422, 32),
	(72887, 'PONG LENGOR', 5827, 422, 32),
	(72888, 'TENGKU LESE', 5827, 422, 32),
	(72889, 'WAE MANTANG', 5827, 422, 32),
	(72890, 'BAJAK', 5828, 422, 32),
	(72891, 'BARU', 5828, 422, 32),
	(72892, 'MATA AIR', 5828, 422, 32),
	(72893, 'REO', 5828, 422, 32),
	(72894, 'ROBEK', 5828, 422, 32),
	(72895, 'RUIS', 5828, 422, 32),
	(72896, 'SALAMA', 5828, 422, 32),
	(72897, 'WANGKUNG', 5828, 422, 32),
	(72898, 'WATU BAUR', 5828, 422, 32),
	(72899, 'WATU TANGO', 5828, 422, 32),
	(72900, 'LANTE', 5829, 422, 32),
	(72901, 'LEMARANG', 5829, 422, 32),
	(72902, 'LOCE', 5829, 422, 32),
	(72903, 'NGGALAK', 5829, 422, 32),
	(72904, 'PARA LANDO', 5829, 422, 32),
	(72905, 'RURA', 5829, 422, 32),
	(72906, 'SAMBI', 5829, 422, 32),
	(72907, 'TOE', 5829, 422, 32),
	(72908, 'TORONG KOE', 5829, 422, 32),
	(72909, 'WAE KAJONG', 5829, 422, 32),
	(72910, 'BANGKA LAO', 5830, 422, 32),
	(72911, 'BELANG TURI', 5830, 422, 32),
	(72912, 'BENTENG KUWU', 5830, 422, 32),
	(72913, 'BEO RAHONG', 5830, 422, 32),
	(72914, 'BEOKAKOR', 5830, 422, 32),
	(72915, 'BULAN', 5830, 422, 32),
	(72916, 'COMPANG DALO', 5830, 422, 32),
	(72917, 'COMPANG NAMUT', 5830, 422, 32),
	(72918, 'CUMBI', 5830, 422, 32),
	(72919, 'GOLO WOROK', 5830, 422, 32),
	(72920, 'KAKOR', 5830, 422, 32),
	(72921, 'MELER', 5830, 422, 32),
	(72922, 'POCO LIKING', 5830, 422, 32),
	(72923, 'PONG LALE', 5830, 422, 32),
	(72924, 'PONG LAO', 5830, 422, 32),
	(72925, 'PONG LEKO', 5830, 422, 32),
	(72926, 'PONG MURUNG', 5830, 422, 32),
	(72927, 'RAI', 5830, 422, 32),
	(72928, 'WAE BELANG', 5830, 422, 32),
	(72929, 'GARA', 5831, 422, 32),
	(72930, 'GOLO LAMBO', 5831, 422, 32),
	(72931, 'GOLO MUNTAS', 5831, 422, 32),
	(72932, 'ITENG', 5831, 422, 32),
	(72933, 'JAONG', 5831, 422, 32),
	(72934, 'KOAK', 5831, 422, 32),
	(72935, 'LANGGO', 5831, 422, 32),
	(72936, 'LEGU', 5831, 422, 32),
	(72937, 'LOLANG', 5831, 422, 32),
	(72938, 'LUNGAR', 5831, 422, 32),
	(72939, 'MOCOK', 5831, 422, 32),
	(72940, 'NGKAER', 5831, 422, 32),
	(72941, 'PAKA', 5831, 422, 32),
	(72942, 'PAPANG', 5831, 422, 32),
	(72943, 'PONGGEOK', 5831, 422, 32),
	(72944, 'PONGKOR', 5831, 422, 32),
	(72945, 'SATAR LOUNG', 5831, 422, 32),
	(72946, 'TADO', 5831, 422, 32),
	(72947, 'TAL', 5831, 422, 32),
	(72948, 'ULU BELANG', 5831, 422, 32),
	(72949, 'UMUNG', 5831, 422, 32),
	(72950, 'WAE AJANG', 5831, 422, 32),
	(72951, 'WEWO', 5831, 422, 32),
	(72952, 'BEA KONDO', 5832, 422, 32),
	(72953, 'BORIK', 5832, 422, 32),
	(72954, 'CAMBIR LECA', 5832, 422, 32),
	(72955, 'CEKA LUJU', 5832, 422, 32),
	(72956, 'CIRENG', 5832, 422, 32),
	(72957, 'GOLO ROPONG', 5832, 422, 32),
	(72958, 'GULUNG', 5832, 422, 32),
	(72959, 'HILIHINTIR', 5832, 422, 32),
	(72960, 'KOLE', 5832, 422, 32),
	(72961, 'LIA', 5832, 422, 32),
	(72962, 'LING', 5832, 422, 32),
	(72963, 'MATA WAE', 5832, 422, 32),
	(72964, 'NAO', 5832, 422, 32),
	(72965, 'NUCA MOLAS', 5832, 422, 32),
	(72966, 'POPO', 5832, 422, 32),
	(72967, 'RENDA', 5832, 422, 32),
	(72968, 'RUANG', 5832, 422, 32),
	(72969, 'SATAR LENDA', 5832, 422, 32),
	(72970, 'SATAR LUJU', 5832, 422, 32),
	(72971, 'SATAR RUWUK', 5832, 422, 32),
	(72972, 'TERONG', 5832, 422, 32),
	(72973, 'TODO', 5832, 422, 32),
	(72974, 'WONGKA', 5832, 422, 32),
	(72975, 'BANGKA KENDA', 5833, 422, 32),
	(72976, 'BANGKAJONG', 5833, 422, 32),
	(72977, 'COMPANG NDEHES', 5833, 422, 32),
	(72978, 'GOLO CADOR', 5833, 422, 32),
	(72979, 'GOLO MENDO', 5833, 422, 32),
	(72980, 'GOLO NEJANG', 5833, 422, 32),
	(72981, 'GOLO WATU/WUTU', 5833, 422, 32),
	(72982, 'GOLO WUA (GOLO WUAS)', 5833, 422, 32),
	(72983, 'LALONG', 5833, 422, 32),
	(72984, 'LONGKO', 5833, 422, 32),
	(72985, 'NDEHES', 5833, 422, 32),
	(72986, 'POCO', 5833, 422, 32),
	(72987, 'RANAKA/RENAKA', 5833, 422, 32),
	(72988, 'RANGGI', 5833, 422, 32),
	(72989, 'SATAR NGKELING', 5833, 422, 32),
	(72990, 'WAE MULU', 5833, 422, 32),
	(72991, 'WAE RII', 5833, 422, 32),
	(72992, 'BATU TIGA', 5834, 423, 32),
	(72993, 'GOLO KETAK', 5834, 423, 32),
	(72994, 'GOLO LUJANG', 5834, 423, 32),
	(72995, 'GOLO SEPANG', 5834, 423, 32),
	(72996, 'MBUIT', 5834, 423, 32),
	(72997, 'PONTIANAK (TANJUNG PONTIANAK)', 5834, 423, 32),
	(72998, 'POTA WANGKA', 5834, 423, 32),
	(72999, 'SEPANG', 5834, 423, 32),
	(73000, 'TANJUNG BOLENG', 5834, 423, 32)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 72")
}
