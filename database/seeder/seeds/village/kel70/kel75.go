package kel70

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel75() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(75001, 'SAUBEBA', 6026, 438, 33),
	(75002, 'SAYRO', 6026, 438, 33),
	(75003, 'SINGGIMEBA', 6026, 438, 33),
	(75004, 'TANAH RUBU', 6026, 438, 33),
	(75005, 'TELUK MUBRI', 6026, 438, 33),
	(75006, 'WARBEFOR', 6026, 438, 33),
	(75007, 'YONGGAM', 6026, 438, 33),
	(75008, 'YOOM I', 6026, 438, 33),
	(75009, 'YOOM II', 6026, 438, 33),
	(75010, 'ASKA', 6027, 438, 33),
	(75011, 'AURMIOS', 6027, 438, 33),
	(75012, 'BOWI SUBUR', 6027, 438, 33),
	(75013, 'IGOR', 6027, 438, 33),
	(75014, 'INYEI', 6027, 438, 33),
	(75015, 'JOWEN', 6027, 438, 33),
	(75016, 'KALI MERAH', 6027, 438, 33),
	(75017, 'KOYANI (KOYAMI)', 6027, 438, 33),
	(75018, 'MACUAN', 6027, 438, 33),
	(75019, 'MAKWAN', 6027, 438, 33),
	(75020, 'MANSABURI (SMARYAM)', 6027, 438, 33),
	(75021, 'MANTEDI', 6027, 438, 33),
	(75022, 'MASNI', 6027, 438, 33),
	(75023, 'MEIFORGA', 6027, 438, 33),
	(75024, 'MEMBOWI', 6027, 438, 33),
	(75025, 'MEREJEMEG (MEREJNEG)', 6027, 438, 33),
	(75026, 'MEYERUK', 6027, 438, 33),
	(75027, 'MEYOF II', 6027, 438, 33),
	(75028, 'MOWBJA (MOUBJA/MOBJA)', 6027, 438, 33),
	(75029, 'MUARA PRAFI', 6027, 438, 33),
	(75030, 'MUARA WARIORI', 6027, 438, 33),
	(75031, 'PRAFI BARAT', 6027, 438, 33),
	(75032, 'RIRINFOS', 6027, 438, 33),
	(75033, 'SEMBAB', 6027, 438, 33),
	(75034, 'SIBUNI', 6027, 438, 33),
	(75035, 'SUMBER BOGA', 6027, 438, 33),
	(75036, 'UNDI', 6027, 438, 33),
	(75037, 'UREY', 6027, 438, 33),
	(75038, 'WAMFOURA (WAM BOLA)', 6027, 438, 33),
	(75039, 'WARIORI', 6027, 438, 33),
	(75040, 'WARIORI INDAH', 6027, 438, 33),
	(75041, 'YENSUM', 6027, 438, 33),
	(75042, 'YONSORIBO (YEN SORIBO)', 6027, 438, 33),
	(75043, 'AIMASI', 6028, 438, 33),
	(75044, 'BENDIP MATOA (DEBIT MOTOA)', 6028, 438, 33),
	(75045, 'BOGOR', 6028, 438, 33),
	(75046, 'DESAY', 6028, 438, 33),
	(75047, 'INGUISI (INGKWOISI)', 6028, 438, 33),
	(75048, 'KALI AMIN', 6028, 438, 33),
	(75049, 'KRENU (KERENU/KRENUI)', 6028, 438, 33),
	(75050, 'LISMAUNGU (LISMANGGU)', 6028, 438, 33),
	(75051, 'MEBIJI / MEBJI', 6028, 438, 33),
	(75052, 'PRAFI MULIA/MULYA', 6028, 438, 33),
	(75053, 'SEMI (SOMI)', 6028, 438, 33),
	(75054, 'UDAPI HILIR (UDAP HILIR)', 6028, 438, 33),
	(75055, 'UHYEHEKBRIG (OGYEHEK/OYEHEK BRIG)', 6028, 438, 33),
	(75056, 'UMBUI (UMBUY)', 6028, 438, 33),
	(75057, 'WANGGONOMA', 6028, 438, 33),
	(75058, 'WASEGI POP (WASEKI)', 6028, 438, 33),
	(75059, 'WASEGI/WASEKI INDAH', 6028, 438, 33),
	(75060, 'KAIRONI', 6029, 438, 33),
	(75061, 'KASI (KASSI)', 6029, 438, 33),
	(75062, 'MANGGUPI', 6029, 438, 33),
	(75063, 'MEYOF I', 6029, 438, 33),
	(75064, 'SARAY', 6029, 438, 33),
	(75065, 'SIDEY', 6029, 438, 33),
	(75066, 'SIDEY BARU', 6029, 438, 33),
	(75067, 'SIDEY JAYA', 6029, 438, 33),
	(75068, 'SIDEY MAKMUR', 6029, 438, 33),
	(75069, 'WAMNOWI (WOMNOWI)', 6029, 438, 33),
	(75070, 'WARAMUI (WARAMOI)', 6029, 438, 33),
	(75071, 'WARIKI', 6029, 438, 33),
	(75072, 'AYAWI', 6030, 438, 33),
	(75073, 'CUYEHEP', 6030, 438, 33),
	(75074, 'HANGHOUW', 6030, 438, 33),
	(75075, 'IMBEISIKA', 6030, 438, 33),
	(75076, 'IMBOISIKA', 6030, 438, 33),
	(75077, 'IMBOITI', 6030, 438, 33),
	(75078, 'IMHASUMA', 6030, 438, 33),
	(75079, 'INDIBO', 6030, 438, 33),
	(75080, 'MBATMA', 6030, 438, 33),
	(75081, 'MENYUMFOKU', 6030, 438, 33),
	(75082, 'MIRONI', 6030, 438, 33),
	(75083, 'MISABUGOID', 6030, 438, 33),
	(75084, 'NINGDIP', 6030, 438, 33),
	(75085, 'REMBUY', 6030, 438, 33),
	(75086, 'UKEMBOISI', 6030, 438, 33),
	(75087, 'UKOPTI', 6030, 438, 33),
	(75088, 'UMNUM', 6030, 438, 33),
	(75089, 'URWAMBEI', 6030, 438, 33),
	(75090, 'WARAMI', 6030, 438, 33),
	(75091, 'WARIARI', 6030, 438, 33),
	(75092, 'WARKAPI', 6030, 438, 33),
	(75093, 'WARMAWAI', 6030, 438, 33),
	(75094, 'WARNYETI', 6030, 438, 33),
	(75095, 'WEDONI', 6030, 438, 33),
	(75096, 'DINDEY', 6031, 438, 33),
	(75097, 'DUWIN (DUWIN UMSINI)', 6031, 438, 33),
	(75098, 'GUEINTUY (GUENTUY)', 6031, 438, 33),
	(75099, 'HINGK', 6031, 438, 33),
	(75100, 'IBUWAU', 6031, 438, 33),
	(75101, 'IMBOISRATI (IBOISRATI)', 6031, 438, 33),
	(75102, 'INDISEY', 6031, 438, 33),
	(75103, 'MADRAT', 6031, 438, 33),
	(75104, 'MENIY (MENY)', 6031, 438, 33),
	(75105, 'NGUNGGUEN', 6031, 438, 33),
	(75106, 'NIMBAY', 6031, 438, 33),
	(75107, 'SNAIMBOY', 6031, 438, 33),
	(75108, 'SOTEA', 6031, 438, 33),
	(75109, 'SRAYNDABEY (SRAINDABEY)', 6031, 438, 33),
	(75110, 'SUBSAY', 6031, 438, 33),
	(75111, 'TANAH MERAH', 6031, 438, 33),
	(75112, 'UMCEN (UMCEM)', 6031, 438, 33),
	(75113, 'WARMARE', 6031, 438, 33),
	(75114, 'DIBERA', 6032, 439, 33),
	(75115, 'DISIHU (DIHISU)', 6032, 439, 33),
	(75116, 'DISRA (DESRA)', 6032, 439, 33),
	(75117, 'DUHUGESA', 6032, 439, 33),
	(75118, 'HUGAMOD (HUGOMOT)', 6032, 439, 33),
	(75119, 'INYUARA', 6032, 439, 33),
	(75120, 'ISIM', 6032, 439, 33),
	(75121, 'MINDERMES', 6032, 439, 33),
	(75122, 'SIBJO', 6032, 439, 33),
	(75123, 'TAHOSTA (TOHOSTA)', 6032, 439, 33),
	(75124, 'TUBES', 6032, 439, 33),
	(75125, 'UMUHOUSI (UMOUSI)', 6032, 439, 33),
	(75126, 'DEMBEK', 6033, 439, 33),
	(75127, 'DEMINI', 6033, 439, 33),
	(75128, 'GAYA BARU', 6033, 439, 33),
	(75129, 'NENY PANTAI (NENEI PANTAI/NIJ)', 6033, 439, 33),
	(75130, 'SIWI', 6033, 439, 33),
	(75131, 'WAREN', 6033, 439, 33),
	(75132, 'YEKWANDI', 6033, 439, 33),
	(75133, 'ARYAWEN MOHO (ARYAWENM)', 6034, 439, 33),
	(75134, 'BENYAS', 6034, 439, 33),
	(75135, 'DISI', 6034, 439, 33),
	(75136, 'HIYOU', 6034, 439, 33),
	(75137, 'NENEI (NENEY)', 6034, 439, 33),
	(75138, 'SESUM', 6034, 439, 33),
	(75139, 'WAMA', 6034, 439, 33),
	(75140, 'AKEJU', 6035, 439, 33),
	(75141, 'MARGOMULYO', 6035, 439, 33),
	(75142, 'MARGORUKUN', 6035, 439, 33),
	(75143, 'MASABUI DUA (II)', 6035, 439, 33),
	(75144, 'MASABUI SATU (I)', 6035, 439, 33),
	(75145, 'MUARI', 6035, 439, 33),
	(75146, 'ORANSBARI', 6035, 439, 33),
	(75147, 'SIDOMULYO', 6035, 439, 33),
	(75148, 'SINDANG JAYA (SIDANG JAYA)', 6035, 439, 33),
	(75149, 'WAMIGTI', 6035, 439, 33),
	(75150, 'WANDOKI', 6035, 439, 33),
	(75151, 'WARBIADI', 6035, 439, 33),
	(75152, 'WARKWANDI', 6035, 439, 33),
	(75153, 'WAROSER', 6035, 439, 33),
	(75154, 'WATARIRI', 6035, 439, 33),
	(75155, 'ABRESSO (ABRESO)', 6036, 439, 33),
	(75156, 'BAMAHA', 6036, 439, 33),
	(75157, 'HAMAWI', 6036, 439, 33),
	(75158, 'HAMOR', 6036, 439, 33),
	(75159, 'KOBREY', 6036, 439, 33),
	(75160, 'MAMBREMA', 6036, 439, 33),
	(75161, 'NUHUWAI (NUHUEI/NUHUWEI)', 6036, 439, 33),
	(75162, 'RANSIKI KOTA', 6036, 439, 33),
	(75163, 'SABRI', 6036, 439, 33),
	(75164, 'SUSMOROF', 6036, 439, 33),
	(75165, 'SWER', 6036, 439, 33),
	(75166, 'TOBOO (TOBOUW/TOBOU)', 6036, 439, 33),
	(75167, 'WAMCEI', 6036, 439, 33),
	(75168, 'YAMBOY (YAMBOI)', 6036, 439, 33),
	(75169, 'KAPRUS', 6037, 439, 33),
	(75170, 'REYOB', 6037, 439, 33),
	(75171, 'SEIMEBA', 6037, 439, 33),
	(75172, 'YERMATUM', 6037, 439, 33),
	(75173, 'AINOD', 6038, 440, 33),
	(75174, 'AISYO', 6038, 440, 33),
	(75175, 'BORI', 6038, 440, 33),
	(75176, 'BORI TIMUR', 6038, 440, 33),
	(75177, 'FAITMAYAF', 6038, 440, 33),
	(75178, 'FAITMAYAT BARAT', 6038, 440, 33),
	(75179, 'FAITSAWE', 6038, 440, 33),
	(75180, 'FRAWEBO', 6038, 440, 33),
	(75181, 'FUTON', 6038, 440, 33),
	(75182, 'KOCUAS UTARA', 6038, 440, 33),
	(75183, 'KOCUWER', 6038, 440, 33),
	(75184, 'KOCUWER SELATAN', 6038, 440, 33),
	(75185, 'KOKAS', 6038, 440, 33),
	(75186, 'KONJA', 6038, 440, 33),
	(75187, 'KUMURKEK', 6038, 440, 33),
	(75188, 'KUMURKEK BARAT', 6038, 440, 33),
	(75189, 'MARTAIM', 6038, 440, 33),
	(75190, 'SAMPIKA', 6038, 440, 33),
	(75191, 'SIKOF', 6038, 440, 33),
	(75192, 'SUNEI', 6038, 440, 33),
	(75193, 'SUSUMUK', 6038, 440, 33),
	(75194, 'TAHAHITE', 6038, 440, 33),
	(75195, 'WERJAYA', 6038, 440, 33),
	(75196, 'ASIAF ZAMAN', 6039, 440, 33),
	(75197, 'AWET MAIN', 6039, 440, 33),
	(75198, 'FUOG', 6039, 440, 33),
	(75199, 'FUOG SELATAN', 6039, 440, 33),
	(75200, 'HORA IEK', 6039, 440, 33),
	(75201, 'IMSUN', 6039, 440, 33),
	(75202, 'KAITANA', 6039, 440, 33),
	(75203, 'KISOR', 6039, 440, 33),
	(75204, 'KRUS', 6039, 440, 33),
	(75205, 'ROMA', 6039, 440, 33),
	(75206, 'SABAH', 6039, 440, 33),
	(75207, 'SAME RAKATOR', 6039, 440, 33),
	(75208, 'SORRY', 6039, 440, 33),
	(75209, 'TAHSIMARA', 6039, 440, 33),
	(75210, 'TOLAK', 6039, 440, 33),
	(75211, 'YEEK', 6039, 440, 33),
	(75212, 'AIKRER', 6040, 440, 33),
	(75213, 'AISA', 6040, 440, 33),
	(75214, 'AITREM', 6040, 440, 33),
	(75215, 'BUOH SA', 6040, 440, 33),
	(75216, 'FUOG', 6040, 440, 33),
	(75217, 'SABAH', 6040, 440, 33),
	(75218, 'SANEM', 6040, 440, 33),
	(75219, 'SAWIN', 6040, 440, 33),
	(75220, 'WAKOM', 6040, 440, 33),
	(75221, 'AIMAU', 6041, 440, 33),
	(75222, 'AINESRA', 6041, 440, 33),
	(75223, 'AKINGGING', 6041, 440, 33),
	(75224, 'FRAMNEWAY', 6041, 440, 33),
	(75225, 'MESYAM', 6041, 440, 33),
	(75226, 'SASIOR NAFE', 6041, 440, 33),
	(75227, 'TIEFROMEN', 6041, 440, 33),
	(75228, 'MAKIRI', 6042, 440, 33),
	(75229, 'SRUMATE', 6042, 440, 33),
	(75230, 'WARBA', 6042, 440, 33),
	(75231, 'WINUNI', 6042, 440, 33),
	(75232, 'WOMBA', 6042, 440, 33),
	(75233, 'WORMU', 6042, 440, 33),
	(75234, 'AIFAM', 6043, 440, 33),
	(75235, 'AIKUS', 6043, 440, 33),
	(75236, 'AIWESA', 6043, 440, 33),
	(75237, 'ASSEM', 6043, 440, 33),
	(75238, 'AYATA', 6043, 440, 33),
	(75239, 'FAAN', 6043, 440, 33),
	(75240, 'FRAMBU', 6043, 440, 33),
	(75241, 'KAMAT', 6043, 440, 33),
	(75242, 'MUPAS', 6043, 440, 33),
	(75243, 'PITOR', 6043, 440, 33),
	(75244, 'SAUD', 6043, 440, 33),
	(75245, 'TIAM', 6043, 440, 33),
	(75246, 'AYAWASI', 6044, 440, 33),
	(75247, 'AYAWASI SELATAN', 6044, 440, 33),
	(75248, 'AYAWASI TIMUR', 6044, 440, 33),
	(75249, 'FONATU', 6044, 440, 33),
	(75250, 'HAENKANES', 6044, 440, 33),
	(75251, 'HOWAIT', 6044, 440, 33),
	(75252, 'IRATA', 6044, 440, 33),
	(75253, 'KONJA', 6044, 440, 33),
	(75254, 'KONJA BARAT', 6044, 440, 33),
	(75255, 'MAN', 6044, 440, 33),
	(75256, 'MOSUN', 6044, 440, 33),
	(75257, 'MOSUN TIMUR', 6044, 440, 33),
	(75258, 'MOSUN UTARA', 6044, 440, 33),
	(75259, 'MOWES', 6044, 440, 33),
	(75260, 'NESET', 6044, 440, 33),
	(75261, 'RAHA', 6044, 440, 33),
	(75262, 'SUSAI', 6044, 440, 33),
	(75263, 'WAYANE', 6044, 440, 33),
	(75264, 'YARAT', 6044, 440, 33),
	(75265, 'YARAT TIMUR', 6044, 440, 33),
	(75266, 'AFKREM', 6045, 440, 33),
	(75267, 'AITINYO', 6045, 440, 33),
	(75268, 'BOFAIT', 6045, 440, 33),
	(75269, 'IBASUF', 6045, 440, 33),
	(75270, 'IHORE', 6045, 440, 33),
	(75271, 'IROHMRAR', 6045, 440, 33),
	(75272, 'KAMRO', 6045, 440, 33),
	(75273, 'KAMRO SELATAN', 6045, 440, 33),
	(75274, 'KARSU', 6045, 440, 33),
	(75275, 'KOROM', 6045, 440, 33),
	(75276, 'SIRA AYA', 6045, 440, 33),
	(75277, 'SOWAI SAU', 6045, 440, 33),
	(75278, 'SRIS', 6045, 440, 33),
	(75279, 'SUBIN', 6045, 440, 33),
	(75280, 'SUMANIS', 6045, 440, 33),
	(75281, 'TEHAK KECIL', 6045, 440, 33),
	(75282, 'WIHO', 6045, 440, 33),
	(75283, 'FATASE', 6046, 440, 33),
	(75284, 'FATEM', 6046, 440, 33),
	(75285, 'HASWEH', 6046, 440, 33),
	(75286, 'HOSYO ATA', 6046, 440, 33),
	(75287, 'KAMBUFATEM', 6046, 440, 33),
	(75288, 'KAMBUFATEM UTARA', 6046, 440, 33),
	(75289, 'SIYO', 6046, 440, 33),
	(75290, 'WAYBOMATAH', 6046, 440, 33),
	(75291, 'FAITHOWES', 6047, 440, 33),
	(75292, 'IKUF', 6047, 440, 33),
	(75293, 'IKUF UTARA', 6047, 440, 33),
	(75294, 'ISIR', 6047, 440, 33),
	(75295, 'JITMAU', 6047, 440, 33),
	(75296, 'JITMAU TIMUR', 6047, 440, 33),
	(75297, 'KAMBUSABO', 6047, 440, 33),
	(75298, 'SARIMO', 6047, 440, 33),
	(75299, 'SRIRTABAM', 6047, 440, 33),
	(75300, 'YUMAME', 6047, 440, 33),
	(75301, 'ASNAIF', 6048, 440, 33),
	(75302, 'AWIT', 6048, 440, 33),
	(75303, 'EROKWERO', 6048, 440, 33),
	(75304, 'EWAY', 6048, 440, 33),
	(75305, 'FRAMBOH', 6048, 440, 33),
	(75306, 'IROH SOHSER', 6048, 440, 33),
	(75307, 'ITIGAH', 6048, 440, 33),
	(75308, 'SABUN', 6048, 440, 33),
	(75309, 'SIRA', 6048, 440, 33),
	(75310, 'SIRA TEE', 6048, 440, 33),
	(75311, 'TOHMRI', 6048, 440, 33),
	(75312, 'WAY U', 6048, 440, 33),
	(75313, 'WRAIT U', 6048, 440, 33),
	(75314, 'YAKSORO', 6048, 440, 33),
	(75315, 'ASMURUF TEE', 6049, 440, 33),
	(75316, 'ASMURUF U', 6049, 440, 33),
	(75317, 'BOHWAT', 6049, 440, 33),
	(75318, 'FAN', 6049, 440, 33),
	(75319, 'FATEGOMI', 6049, 440, 33),
	(75320, 'FRAMAFIR', 6049, 440, 33),
	(75321, 'GOHSAMES', 6049, 440, 33),
	(75322, 'INTA', 6049, 440, 33),
	(75323, 'MIRAFAN', 6049, 440, 33),
	(75324, 'SUBRIT', 6049, 440, 33),
	(75325, 'TEHAK BESAR', 6049, 440, 33),
	(75326, 'TEHAK TEE', 6049, 440, 33),
	(75327, 'AFES', 6050, 440, 33),
	(75328, 'AYAMARU', 6050, 440, 33),
	(75329, 'FRAHARO', 6050, 440, 33),
	(75330, 'FRAMU', 6050, 440, 33),
	(75331, 'MEFKAJIM II', 6050, 440, 33),
	(75332, 'SMUSUWIOH', 6050, 440, 33),
	(75333, 'TUSO', 6050, 440, 33),
	(75334, 'TWER', 6050, 440, 33),
	(75335, 'CHALIAT', 6051, 440, 33),
	(75336, 'FANSE', 6051, 440, 33),
	(75337, 'SEHU', 6051, 440, 33),
	(75338, 'SFACKO', 6051, 440, 33),
	(75339, 'SFAKRAKA', 6051, 440, 33),
	(75340, 'SIEN', 6051, 440, 33),
	(75341, 'SOROWAN', 6051, 440, 33),
	(75342, 'TBO', 6051, 440, 33),
	(75343, 'ADOH', 6052, 440, 33),
	(75344, 'ORAIN', 6052, 440, 33),
	(75345, 'ORSU', 6052, 440, 33),
	(75346, 'RAWAS', 6052, 440, 33),
	(75347, 'SEGIOR', 6052, 440, 33),
	(75348, 'SOAN', 6052, 440, 33),
	(75349, 'SOSIAN', 6052, 440, 33),
	(75350, 'TEMEL', 6052, 440, 33),
	(75351, 'WARBO', 6052, 440, 33),
	(75352, 'WOMAN', 6052, 440, 33),
	(75353, 'HAWIOH', 6053, 440, 33),
	(75354, 'KANISABAR', 6053, 440, 33),
	(75355, 'KOMA KOMA', 6053, 440, 33),
	(75356, 'LEMAUK KLIT', 6053, 440, 33),
	(75357, 'SAGRIM', 6053, 440, 33),
	(75358, 'SAUF', 6053, 440, 33),
	(75359, 'SEMBARO', 6053, 440, 33),
	(75360, 'SENEH', 6053, 440, 33),
	(75361, 'SFARARE', 6053, 440, 33),
	(75362, 'SIMIYAH', 6053, 440, 33),
	(75363, 'ARUS', 6054, 440, 33),
	(75364, 'ASSES', 6054, 440, 33),
	(75365, 'FAIT NGGRE', 6054, 440, 33),
	(75366, 'FAITSIMAR', 6054, 440, 33),
	(75367, 'ISNUM', 6054, 440, 33),
	(75368, 'KOFAIT', 6054, 440, 33),
	(75369, 'SUFU', 6054, 440, 33),
	(75370, 'BAWY', 6055, 440, 33),
	(75371, 'FIANE', 6055, 440, 33),
	(75372, 'HUFIOH', 6055, 440, 33),
	(75373, 'ISME', 6055, 440, 33),
	(75374, 'KARTAPURA', 6055, 440, 33),
	(75375, 'MEN', 6055, 440, 33),
	(75376, 'RINDU', 6055, 440, 33),
	(75377, 'SEMU', 6055, 440, 33),
	(75378, 'TUT', 6055, 440, 33),
	(75379, 'YOHWER', 6055, 440, 33),
	(75380, 'FAITMAJIN', 6056, 440, 33),
	(75381, 'FAITSIUR', 6056, 440, 33),
	(75382, 'HUBERITA', 6056, 440, 33),
	(75383, 'INSAS', 6056, 440, 33),
	(75384, 'ISMAYU', 6056, 440, 33),
	(75385, 'KAMBUAYA', 6056, 440, 33),
	(75386, 'KEYUM', 6056, 440, 33),
	(75387, 'SEFAYIT', 6056, 440, 33),
	(75388, 'FAITWOSUR', 6057, 440, 33),
	(75389, 'ISUSU', 6057, 440, 33),
	(75390, 'KAMBUIFA', 6057, 440, 33),
	(75391, 'KAMBUSKATO', 6057, 440, 33),
	(75392, 'KAMBUSKATO UTARA', 6057, 440, 33),
	(75393, 'MANO', 6057, 440, 33),
	(75394, 'SIPAT', 6057, 440, 33),
	(75395, 'ASMURUF TEE', 6058, 440, 33),
	(75396, 'ASMURUF U', 6058, 440, 33),
	(75397, 'BOHWAT', 6058, 440, 33),
	(75398, 'FAN', 6058, 440, 33),
	(75399, 'FATEGOMI', 6058, 440, 33),
	(75400, 'FRAMAFIR', 6058, 440, 33),
	(75401, 'GOHSAMES', 6058, 440, 33),
	(75402, 'INTA', 6058, 440, 33),
	(75403, 'MIRAFAN', 6058, 440, 33),
	(75404, 'SUBRIT', 6058, 440, 33),
	(75405, 'TEHAK BESAR', 6058, 440, 33),
	(75406, 'TEHAK TEE', 6058, 440, 33),
	(75407, 'FRABO', 6059, 440, 33),
	(75408, 'KARFA', 6059, 440, 33),
	(75409, 'KONA', 6059, 440, 33),
	(75410, 'KOSAH', 6059, 440, 33),
	(75411, 'MAPURA', 6059, 440, 33),
	(75412, 'SUWIAM', 6059, 440, 33),
	(75413, 'TOMASE', 6059, 440, 33),
	(75414, 'BAKRABI', 6060, 440, 33),
	(75415, 'BATU GADING', 6060, 440, 33),
	(75416, 'CEGE', 6060, 440, 33),
	(75417, 'DATA', 6060, 440, 33),
	(75418, 'KADAI', 6060, 440, 33),
	(75419, 'KARELLA', 6060, 440, 33),
	(75420, 'KOMBIF', 6060, 440, 33),
	(75421, 'LAKUKANG', 6060, 440, 33),
	(75422, 'LAPASA', 6060, 440, 33),
	(75423, 'LAPPA UPANG', 6060, 440, 33),
	(75424, 'MAHOS', 6060, 440, 33),
	(75425, 'MARIO', 6060, 440, 33),
	(75426, 'MATTAMPA WALIE', 6060, 440, 33),
	(75427, 'MATTIRO WALIE', 6060, 440, 33),
	(75428, 'NAFASI', 6060, 440, 33),
	(75429, 'PADAELO', 6060, 440, 33),
	(75430, 'PATTIRO', 6060, 440, 33),
	(75431, 'RUFASES', 6060, 440, 33),
	(75432, 'SAWO', 6060, 440, 33),
	(75433, 'SEYA', 6060, 440, 33),
	(75434, 'SUMALING', 6060, 440, 33),
	(75435, 'SUSWA', 6060, 440, 33),
	(75436, 'TELLONGENG', 6060, 440, 33),
	(75437, 'TELLU BOCCOE', 6060, 440, 33),
	(75438, 'UJUNG SALANGKETO', 6060, 440, 33),
	(75439, 'UJUNG TANAH', 6060, 440, 33),
	(75440, 'WABAM', 6060, 440, 33),
	(75441, 'BATU ZAMAN', 6061, 440, 33),
	(75442, 'KANGGILO', 6061, 440, 33),
	(75443, 'MOPI', 6061, 440, 33),
	(75444, 'NONOMI', 6061, 440, 33),
	(75445, 'SARAFAMBAI', 6061, 440, 33),
	(75446, 'SAWARA JAYA', 6061, 440, 33),
	(75447, 'URI', 6061, 440, 33),
	(75448, 'WAREN I', 6061, 440, 33),
	(75449, 'WAREN II', 6061, 440, 33),
	(75450, 'AYAU (AYAUBEY)', 6062, 441, 33),
	(75451, 'HUNGKU', 6062, 441, 33),
	(75452, 'IGMBAI (IGEMBAI/IMBEISBA)', 6062, 441, 33),
	(75453, 'IMBAI', 6062, 441, 33),
	(75454, 'IRAIWERI', 6062, 441, 33),
	(75455, 'MENTUBEI (MENTUBEY)', 6062, 441, 33),
	(75456, 'PAMAHA', 6062, 441, 33),
	(75457, 'SPIYOUGUP', 6062, 441, 33),
	(75458, 'SRUBEY', 6062, 441, 33),
	(75459, 'SUSI', 6062, 441, 33),
	(75460, 'SUTEIBEI (SUTEIBEY)', 6062, 441, 33),
	(75461, 'TESTEGA PAMAHA', 6062, 441, 33),
	(75462, 'ULONG (ULLONG)', 6062, 441, 33),
	(75463, 'UPER', 6062, 441, 33),
	(75464, 'ITKAU', 6063, 441, 33),
	(75465, 'KROBUT (KIWOT)', 6063, 441, 33),
	(75466, 'NGISROW (INGISROW)', 6063, 441, 33),
	(75467, 'SAKUMI (SAKUNI/SAKURAI)', 6063, 441, 33),
	(75468, 'SIBIOGUD', 6063, 441, 33),
	(75469, 'TOMBROK', 6063, 441, 33),
	(75470, 'TUBYAM (TUABYAM)', 6063, 441, 33),
	(75471, 'AIGA', 6064, 441, 33),
	(75472, 'BINGRAYUD (BINGGARUD)', 6064, 441, 33),
	(75473, 'BINGWAIMOD (BIGWAIMUD)', 6064, 441, 33),
	(75474, 'CATUBOUW', 6064, 441, 33),
	(75475, 'COIJUT (COIJUD)', 6064, 441, 33),
	(75476, 'IDEMAI', 6064, 441, 33),
	(75477, 'IJIGREK (IJIGREG)', 6064, 441, 33),
	(75478, 'IMANDRIGO', 6064, 441, 33),
	(75479, 'JIM', 6064, 441, 33),
	(75480, 'KAUNGWAM', 6064, 441, 33),
	(75481, 'MANGGESUK', 6064, 441, 33),
	(75482, 'MIEICOMTI (MIEY COMTY/MIEYCOMTI)', 6064, 441, 33),
	(75483, 'MIHIJ', 6064, 441, 33),
	(75484, 'MINHOU (MIHOU)', 6064, 441, 33),
	(75485, 'MINMO', 6064, 441, 33),
	(75486, 'NDABOUW', 6064, 441, 33),
	(75487, 'SLOMIOU', 6064, 441, 33),
	(75488, 'SOUDIN', 6064, 441, 33),
	(75489, 'SUGEMAH (SUGEMEH)', 6064, 441, 33),
	(75490, 'TIMTOU (TIMTOUW)', 6064, 441, 33),
	(75491, 'UNTI', 6064, 441, 33),
	(75492, 'CIGERA', 6065, 441, 33),
	(75493, 'CIRNOHU', 6065, 441, 33),
	(75494, 'DANMOU', 6065, 441, 33),
	(75495, 'DEBETIK (DIBETIK)', 6065, 441, 33),
	(75496, 'DEMDAMEI', 6065, 441, 33),
	(75497, 'DESRA (DISRA)', 6065, 441, 33),
	(75498, 'DUGRIMOG', 6065, 441, 33),
	(75499, 'GEDEIRA', 6065, 441, 33),
	(75500, 'IRANMEBA', 6065, 441, 33),
	(75501, 'KUSMENAU', 6065, 441, 33),
	(75502, 'MISEDA', 6065, 441, 33),
	(75503, 'SIRGEMEIH (SERGEMEH)', 6065, 441, 33),
	(75504, 'SNERANER (SNEREMER)', 6065, 441, 33),
	(75505, 'TOMSIR', 6065, 441, 33),
	(75506, 'ARION (ARYON)', 6066, 441, 33),
	(75507, 'CANGOISI', 6066, 441, 33),
	(75508, 'DEMUNTI', 6066, 441, 33),
	(75509, 'GUWEUTI (GUEIUTI/GUWEI UTI)', 6066, 441, 33),
	(75510, 'GWEIPINGBAI (GUEIPINBAI)', 6066, 441, 33),
	(75511, 'HEATIE BOUW (HAKTIEBOU)', 6066, 441, 33),
	(75512, 'HUMEYSI (HUMEISI)', 6066, 441, 33),
	(75513, 'IKIMABOW (IKIMABOU)', 6066, 441, 33),
	(75514, 'KISAP', 6066, 441, 33),
	(75515, 'KWAIYEHEP', 6066, 441, 33),
	(75516, 'KWOK I', 6066, 441, 33),
	(75517, 'KWOK II (KWOKWISNGUNG)', 6066, 441, 33),
	(75518, 'LEIHEAK (LIEHAK)', 6066, 441, 33),
	(75519, 'MBEGAU', 6066, 441, 33),
	(75520, 'MBRANDE', 6066, 441, 33),
	(75521, 'MINGGOT (MANGGOT)', 6066, 441, 33),
	(75522, 'MINYEIMEMUT (MINYEIMEMUD)', 6066, 441, 33),
	(75523, 'MONUT', 6066, 441, 33),
	(75524, 'NGIMOUBRE (NGIMOUBRI)', 6066, 441, 33),
	(75525, 'NTAP', 6066, 441, 33),
	(75526, 'NUNGKIMOR', 6066, 441, 33),
	(75527, 'PENIBUT', 6066, 441, 33),
	(75528, 'PUNGGUNG (PUNGUG)', 6066, 441, 33),
	(75529, 'SOPNYAI', 6066, 441, 33),
	(75530, 'TIGAU COMTI', 6066, 441, 33),
	(75531, 'TINGKWOIKIO (TINGWOIKIOU)', 6066, 441, 33),
	(75532, 'TUMBEIBEHEI', 6066, 441, 33),
	(75533, 'UMCEP', 6066, 441, 33),
	(75534, 'URWONG', 6066, 441, 33),
	(75535, 'IMBEISABA', 6067, 441, 33),
	(75536, 'INYEBOW (INYEBOUW)', 6067, 441, 33),
	(75537, 'KONEY', 6067, 441, 33),
	(75538, 'MEMBEY', 6067, 441, 33),
	(75539, 'MEMTI', 6067, 441, 33),
	(75540, 'SISRANG (SESERAN)', 6067, 441, 33),
	(75541, 'USTI', 6067, 441, 33),
	(75542, 'ADUER', 6068, 441, 33),
	(75543, 'AIWOU (AYWOU)', 6068, 441, 33),
	(75544, 'AMBER', 6068, 441, 33),
	(75545, 'ANDANG', 6068, 441, 33),
	(75546, 'ANGGRA', 6068, 441, 33),
	(75547, 'APUI', 6068, 441, 33),
	(75548, 'ASARBEY (ASARBEI)', 6068, 441, 33),
	(75549, 'AWIBEHEY (AWIBEHEI)', 6068, 441, 33),
	(75550, 'AYAU', 6068, 441, 33),
	(75551, 'BAHAMYENTI', 6068, 441, 33),
	(75552, 'BINGWOYUD (BINGYWOYUT)', 6068, 441, 33),
	(75553, 'BUEIBEI (DUEIBEI)', 6068, 441, 33),
	(75554, 'COISI', 6068, 441, 33),
	(75555, 'DEMAISI', 6068, 441, 33),
	(75556, 'DRIYE', 6068, 441, 33),
	(75557, 'FIGOUD', 6068, 441, 33),
	(75558, 'HANDUK', 6068, 441, 33),
	(75559, 'IMBENTI', 6068, 441, 33),
	(75560, 'IMBONGGUN', 6068, 441, 33),
	(75561, 'IMBREKTI', 6068, 441, 33),
	(75562, 'INDABRI', 6068, 441, 33),
	(75563, 'INDONBEY (INDOBEY/INDONDEY)', 6068, 441, 33),
	(75564, 'INGGRAHIM (INGGRAMHIM)', 6068, 441, 33),
	(75565, 'INJUAR', 6068, 441, 33),
	(75566, 'IPINGOISI (IPINGOSI)', 6068, 441, 33),
	(75567, 'KIBAUW (KIPOUW)', 6068, 441, 33),
	(75568, 'KWAU', 6068, 441, 33),
	(75569, 'MAINDA', 6068, 441, 33),
	(75570, 'MBIGMA (MBINGMA)', 6068, 441, 33),
	(75571, 'MEMANKER (MEMANGKER)', 6068, 441, 33),
	(75572, 'MENYAMBOUW (MINYAMBOUW)', 6068, 441, 33),
	(75573, 'MICADIWOR', 6068, 441, 33),
	(75574, 'MICONTI (MICOMTI)', 6068, 441, 33),
	(75575, 'MINGGRE (MINGRE)', 6068, 441, 33),
	(75576, 'MISADWER', 6068, 441, 33),
	(75577, 'MISAPNGOISI (MISANGOISI)', 6068, 441, 33),
	(75578, 'MITIEDE', 6068, 441, 33),
	(75579, 'MNOUBEI (MINOGBEI/MINOGBEY/MINBEY)', 6068, 441, 33),
	(75580, 'MOKWAM', 6068, 441, 33),
	(75581, 'NIMBIAU', 6068, 441, 33),
	(75582, 'NINSIMOI', 6068, 441, 33),
	(75583, 'NJUAR', 6068, 441, 33),
	(75584, 'PIYAUSI (PINYAUSI)', 6068, 441, 33),
	(75585, 'SIGIM', 6068, 441, 33),
	(75586, 'SIMERBEI', 6068, 441, 33),
	(75587, 'SINAITOUSI (SINAYTOISI/SINITOUSI)', 6068, 441, 33),
	(75588, 'SIYAU (SYOU)', 6068, 441, 33),
	(75589, 'SMAINGGI (SMAINGEI/SMANGGEI)', 6068, 441, 33),
	(75590, 'UGJEHEK (UJEHEG)', 6068, 441, 33),
	(75591, 'UMPUG', 6068, 441, 33),
	(75592, 'UNGGA', 6068, 441, 33),
	(75593, 'WAMINDA (WAMMINDA)', 6068, 441, 33),
	(75594, 'ANUK', 6069, 441, 33),
	(75595, 'DUGUHANI (DUGAHANI)', 6069, 441, 33),
	(75596, 'INYAUB', 6069, 441, 33),
	(75597, 'KOBREY', 6069, 441, 33),
	(75598, 'KOPO', 6069, 441, 33),
	(75599, 'KOSTERA', 6069, 441, 33),
	(75600, 'MENESRIJ', 6069, 441, 33),
	(75601, 'SAUGEMEBA', 6069, 441, 33),
	(75602, 'SUNGEDES (SUNGUDES)', 6069, 441, 33),
	(75603, 'SURUREY (SURURAI)', 6069, 441, 33),
	(75604, 'TOMSTERA', 6069, 441, 33),
	(75605, 'TUHUBEA (DUHUBEYA)', 6069, 441, 33),
	(75606, 'ASUM (ANSUM)', 6070, 441, 33),
	(75607, 'AWIGAU (AWAIKUM/AWAIGAU)', 6070, 441, 33),
	(75608, 'DEROUBU (DEUROBA)', 6070, 441, 33),
	(75609, 'DISURA', 6070, 441, 33),
	(75610, 'GENYU', 6070, 441, 33),
	(75611, 'HORETA', 6070, 441, 33),
	(75612, 'IRBOS', 6070, 441, 33),
	(75613, 'SISKEDOUW (SISKEDOWO)', 6070, 441, 33),
	(75614, 'TAIGA (TAIGE)', 6070, 441, 33),
	(75615, 'TRIDAGA', 6070, 441, 33),
	(75616, 'UBEISA', 6070, 441, 33),
	(75617, 'ASAI II (ASAY II)', 6071, 441, 33),
	(75618, 'CENING', 6071, 441, 33),
	(75619, 'DEMOURA (DEMUORA)', 6071, 441, 33),
	(75620, 'DUMBRE (DUMBREY)', 6071, 441, 33),
	(75621, 'IBA', 6071, 441, 33),
	(75622, 'JIGJA (JIJGA)', 6071, 441, 33),
	(75623, 'MEIDOGDA', 6071, 441, 33),
	(75624, 'MEIFEKENI', 6071, 441, 33),
	(75625, 'MEIFOWOSKA', 6071, 441, 33),
	(75626, 'MEIMERSA', 6071, 441, 33),
	(75627, 'MEKSI', 6071, 441, 33),
	(75628, 'MENGEHENA (MEIGEHENAWU)', 6071, 441, 33),
	(75629, 'MENJUGIJI (MEJUGIJI)', 6071, 441, 33),
	(75630, 'MOFOUKEDA', 6071, 441, 33),
	(75631, 'MORUMFEYI', 6071, 441, 33),
	(75632, 'TESTEGA', 6071, 441, 33),
	(75633, 'ABIDON', 6072, 442, 33),
	(75634, 'MEOSBEKWAN', 6072, 442, 33),
	(75635, 'RENI', 6072, 442, 33),
	(75636, 'RUTUM', 6072, 442, 33),
	(75637, 'AMDUI', 6073, 442, 33),
	(75638, 'WAILEBET', 6073, 442, 33),
	(75639, 'WAIMAN', 6073, 442, 33),
	(75640, 'YENANAS', 6073, 442, 33),
	(75641, 'AREFI SELATAN', 6074, 442, 33),
	(75642, 'AREFI TIMUR', 6074, 442, 33),
	(75643, 'YENSAWAI BARAT', 6074, 442, 33),
	(75644, 'YENSAWAI TIMUR', 6074, 442, 33),
	(75645, 'DOREKER (DOREKAR)', 6075, 442, 33),
	(75646, 'MEOSBEKWAN', 6075, 442, 33),
	(75647, 'RENI', 6075, 442, 33),
	(75648, 'RUTUM', 6075, 442, 33),
	(75649, 'YENKAWIR', 6075, 442, 33),
	(75650, 'PULAU TIKUS', 6076, 442, 33),
	(75651, 'SATUKURANO', 6076, 442, 33),
	(75652, 'WEJIM BARAT', 6076, 442, 33),
	(75653, 'WEJIM TIMUR', 6076, 442, 33),
	(75654, 'AWAT', 6077, 442, 33),
	(75655, 'DEER', 6077, 442, 33),
	(75656, 'DIBALAL', 6077, 442, 33),
	(75657, 'MIKIRAN', 6077, 442, 33),
	(75658, 'TOLABI (TOLOBI)', 6077, 442, 33),
	(75659, 'BONWAKIR', 6078, 442, 33),
	(75660, 'SAPORDANCO', 6078, 442, 33),
	(75661, 'WAISAI', 6078, 442, 33),
	(75662, 'WARMASEN', 6078, 442, 33),
	(75663, 'ARBOREK', 6079, 442, 33),
	(75664, 'KABUY (KABUI)', 6079, 442, 33),
	(75665, 'KAPISAWAR', 6079, 442, 33),
	(75666, 'KURKAPA', 6079, 442, 33),
	(75667, 'SAWANDAREK', 6079, 442, 33),
	(75668, 'SAWINGGRAI', 6079, 442, 33),
	(75669, 'YEMBEKWAN', 6079, 442, 33),
	(75670, 'YENBUBA', 6079, 442, 33),
	(75671, 'YENWAUPNOOR (YENWAUPNOR)', 6079, 442, 33),
	(75672, 'FOLLEY', 6080, 442, 33),
	(75673, 'LENMALAS', 6080, 442, 33),
	(75674, 'SATUKORANO (SATO KURANA)', 6080, 442, 33),
	(75675, 'TOMOLOL', 6080, 442, 33),
	(75676, 'WEJIM', 6080, 442, 33),
	(75677, 'BIGA', 6081, 442, 33),
	(75678, 'GAMTA', 6081, 442, 33),
	(75679, 'KAPATCOL', 6081, 442, 33),
	(75680, 'LILINTA', 6081, 442, 33),
	(75681, 'MAGEY', 6081, 442, 33),
	(75682, 'DABATAN', 6082, 442, 33),
	(75683, 'FAFANLAP', 6082, 442, 33),
	(75684, 'HARAPAN JAYA', 6082, 442, 33),
	(75685, 'USAHA JAYA', 6082, 442, 33),
	(75686, 'YELLU', 6082, 442, 33),
	(75687, 'AUDAM', 6083, 442, 33),
	(75688, 'FOLLEY', 6083, 442, 33),
	(75689, 'LIMALAS BARAT', 6083, 442, 33),
	(75690, 'LIMALAS TIMUR', 6083, 442, 33),
	(75691, 'TOMOLOL', 6083, 442, 33),
	(75692, 'USAHA JAYA', 6083, 442, 33),
	(75693, 'KALIAM', 6084, 442, 33),
	(75694, 'KALWAL', 6084, 442, 33),
	(75695, 'SOLOL', 6084, 442, 33),
	(75696, 'WAIBON', 6084, 442, 33),
	(75697, 'KALOBO', 6085, 442, 33),
	(75698, 'SAKABU', 6085, 442, 33),
	(75699, 'WAIBU', 6085, 442, 33),
	(75700, 'WAIJAN', 6085, 442, 33),
	(75701, 'WAILABU', 6085, 442, 33),
	(75702, 'WAILEN', 6085, 442, 33),
	(75703, 'WAIMECI', 6085, 442, 33),
	(75704, 'JEFMAN BARAT', 6086, 442, 33),
	(75705, 'JEFMAN TIMUR', 6086, 442, 33),
	(75706, 'KAPATLAP', 6086, 442, 33),
	(75707, 'SAMATE', 6086, 442, 33),
	(75708, 'WAIDIM', 6086, 442, 33),
	(75709, 'WAMEGA', 6086, 442, 33),
	(75710, 'DUBER', 6087, 442, 33),
	(75711, 'KAPADIRI', 6087, 442, 33),
	(75712, 'RAUKI', 6087, 442, 33),
	(75713, 'URAI', 6087, 442, 33),
	(75714, 'KALITOKO', 6088, 442, 33),
	(75715, 'LOPINTOL', 6088, 442, 33),
	(75716, 'MUMES', 6088, 442, 33),
	(75717, 'WARSAMDIN', 6088, 442, 33),
	(75718, 'ARWAY', 6089, 442, 33),
	(75719, 'BEO', 6089, 442, 33),
	(75720, 'GO', 6089, 442, 33),
	(75721, 'KABILOL', 6089, 442, 33),
	(75722, 'WAIFOI', 6089, 442, 33),
	(75723, 'WARIMAK', 6089, 442, 33),
	(75724, 'BIANCI', 6090, 442, 33),
	(75725, 'MUTUS', 6090, 442, 33),
	(75726, 'SALEO', 6090, 442, 33),
	(75727, 'SELPELE', 6090, 442, 33),
	(75728, 'WAISILIP', 6090, 442, 33),
	(75729, 'GAG', 6091, 442, 33),
	(75730, 'MANYAIFUN', 6091, 442, 33),
	(75731, 'MEOSMANGGARA', 6091, 442, 33),
	(75732, 'PAM', 6091, 442, 33),
	(75733, 'SAUKABU', 6091, 442, 33),
	(75734, 'SAUPAPIR', 6091, 442, 33),
	(75735, 'FRIWEN', 6092, 442, 33),
	(75736, 'SAONEK', 6092, 442, 33),
	(75737, 'SAPORKREN', 6092, 442, 33),
	(75738, 'WAWIYAI', 6092, 442, 33),
	(75739, 'YENBESER', 6092, 442, 33),
	(75740, 'PUPER', 6093, 442, 33),
	(75741, 'URBINASOPEN', 6093, 442, 33),
	(75742, 'YEMBEKAKI', 6093, 442, 33),
	(75743, 'YENSENER (YESNER)', 6093, 442, 33),
	(75744, 'ANDEY', 6094, 442, 33),
	(75745, 'ASUKWERI', 6094, 442, 33),
	(75746, 'BONSAYOR', 6094, 442, 33),
	(75747, 'DARUMBAB', 6094, 442, 33),
	(75748, 'KABARE', 6094, 442, 33),
	(75749, 'KALISADE', 6094, 442, 33),
	(75750, 'BONI', 6095, 442, 33),
	(75751, 'MNIER', 6095, 442, 33),
	(75752, 'WARKORI', 6095, 442, 33),
	(75753, 'WARMANAI', 6095, 442, 33),
	(75754, 'AIMAS', 6096, 443, 33),
	(75755, 'KLABINAIN', 6096, 443, 33),
	(75756, 'KLAIGIT', 6096, 443, 33),
	(75757, 'MAIBO', 6096, 443, 33),
	(75758, 'MALASOM', 6096, 443, 33),
	(75759, 'MALAWELE', 6096, 443, 33),
	(75760, 'MALAWILI', 6096, 443, 33),
	(75761, 'MARIAT GUNUNG', 6096, 443, 33),
	(75762, 'MARIAT PANTAI', 6096, 443, 33),
	(75763, 'WARMON', 6096, 443, 33),
	(75764, 'BAGUNG', 6097, 443, 33),
	(75765, 'DISFRA', 6097, 443, 33),
	(75766, 'KAAS', 6097, 443, 33),
	(75767, 'KLABRA', 6097, 443, 33),
	(75768, 'KLARION', 6097, 443, 33),
	(75769, 'MUMPI', 6097, 443, 33),
	(75770, 'WANURIAN', 6097, 443, 33),
	(75771, 'WENSI', 6097, 443, 33),
	(75772, 'WISBIAK', 6097, 443, 33),
	(75773, 'BUK', 6098, 443, 33),
	(75774, 'INDIWI', 6098, 443, 33),
	(75775, 'KLABOT', 6098, 443, 33),
	(75776, 'KLAIS', 6098, 443, 33),
	(75777, 'MIMPE', 6098, 443, 33),
	(75778, 'MLAKHAN', 6098, 443, 33),
	(75779, 'MLAT', 6098, 443, 33),
	(75780, 'MLAWES', 6098, 443, 33),
	(75781, 'GISIM DARAT', 6099, 443, 33),
	(75782, 'KELADUM', 6099, 443, 33),
	(75783, 'KLAGULU', 6099, 443, 33),
	(75784, 'KLALOMON', 6099, 443, 33),
	(75785, 'KLAMANO', 6099, 443, 33),
	(75786, 'KLAMUGUN', 6099, 443, 33),
	(75787, 'KLASMAN', 6099, 443, 33),
	(75788, 'KLAWANA', 6099, 443, 33),
	(75789, 'MALAJAPA', 6099, 443, 33),
	(75790, 'MALASIGIT', 6099, 443, 33),
	(75791, 'MLAIS', 6099, 443, 33),
	(75792, 'WARIAYAU', 6099, 443, 33),
	(75793, 'WONOSARI', 6099, 443, 33),
	(75794, 'KLAMUGUN', 6100, 443, 33),
	(75795, 'MISKUM', 6100, 443, 33),
	(75796, 'SAENGKEDUK', 6100, 443, 33),
	(75797, 'SELEKOBO', 6100, 443, 33),
	(75798, 'SIWIS', 6100, 443, 33),
	(75799, 'AYOLOKLE', 6101, 443, 33),
	(75800, 'BRIANLO', 6101, 443, 33),
	(75801, 'HOBARD', 6101, 443, 33),
	(75802, 'KLAMNE', 6101, 443, 33),
	(75803, 'MLAWEN', 6101, 443, 33),
	(75804, 'SAS', 6101, 443, 33),
	(75805, 'TARSA', 6101, 443, 33),
	(75806, 'TOTJIN', 6101, 443, 33),
	(75807, 'WILTY', 6101, 443, 33),
	(75808, 'KLASOWOH', 6102, 443, 33),
	(75809, 'KLAWULUH', 6102, 443, 33),
	(75810, 'KLAYALI', 6102, 443, 33),
	(75811, 'KWAKEIK', 6102, 443, 33),
	(75812, 'MALAKOBUTU', 6102, 443, 33),
	(75813, 'MALALILIS', 6102, 443, 33),
	(75814, 'ASBAKEN', 6103, 443, 33),
	(75815, 'BAINKETE', 6103, 443, 33),
	(75816, 'BATU LUBANG/LOBANG', 6103, 443, 33),
	(75817, 'BATU LUBANG/LOBANG PANTAI', 6103, 443, 33),
	(75818, 'KUADAS (KWADAS)', 6103, 443, 33),
	(75819, 'MAKBON', 6103, 443, 33),
	(75820, 'MALAUMKARTA', 6103, 443, 33),
	(75821, 'SAWATUK', 6103, 443, 33),
	(75822, 'TELUK DORE', 6103, 443, 33),
	(75823, 'KLABEN', 6104, 443, 33),
	(75824, 'KLAMALU', 6104, 443, 33),
	(75825, 'KLAMASEN', 6104, 443, 33),
	(75826, 'KLASUKUK', 6104, 443, 33),
	(75827, 'MARIYAI', 6104, 443, 33),
	(75828, 'DASRI', 6105, 443, 33),
	(75829, 'KETAWAS', 6105, 443, 33),
	(75830, 'KLAOS', 6105, 443, 33),
	(75831, 'KLATIM', 6105, 443, 33),
	(75832, 'LUWELALA', 6105, 443, 33),
	(75833, 'MASOS', 6105, 443, 33),
	(75834, 'SULUH', 6105, 443, 33),
	(75835, 'WARBO', 6105, 443, 33),
	(75836, 'ARAR', 6106, 443, 33),
	(75837, 'JEFLIO (YEFLIO/YEFILIO)', 6106, 443, 33),
	(75838, 'KLALIN (KLAIN)', 6106, 443, 33),
	(75839, 'KLASMELEK', 6106, 443, 33),
	(75840, 'MAKBALIM', 6106, 443, 33),
	(75841, 'MAKBUSUN (MAKBUSUM)', 6106, 443, 33),
	(75842, 'MAKOTYAMSA', 6106, 443, 33),
	(75843, 'WEN', 6106, 443, 33),
	(75844, 'KLAFDALIM', 6107, 443, 33),
	(75845, 'KLAFORO', 6107, 443, 33),
	(75846, 'KLASARIN', 6107, 443, 33),
	(75847, 'KLASOF', 6107, 443, 33),
	(75848, 'KLAWOTON', 6107, 443, 33),
	(75849, 'NINJEMUR', 6107, 443, 33),
	(75850, 'WONOSOBO', 6107, 443, 33),
	(75851, 'KATINIM', 6108, 443, 33),
	(75852, 'MAJARAN', 6108, 443, 33),
	(75853, 'MAJENER', 6108, 443, 33),
	(75854, 'MALAUS', 6108, 443, 33),
	(75855, 'MATAWOLOT', 6108, 443, 33),
	(75856, 'RAWA SUGI', 6108, 443, 33),
	(75857, 'WALAL', 6108, 443, 33),
	(75858, 'DULBATAN', 6109, 443, 33),
	(75859, 'DURIANKARI', 6109, 443, 33),
	(75860, 'KATLOL (KLOTLOF/KATLOF)', 6109, 443, 33),
	(75861, 'MANOKET', 6109, 443, 33),
	(75862, 'MARALOL', 6109, 443, 33),
	(75863, 'SAILOLOF', 6109, 443, 33),
	(75864, 'SAKAPUL', 6109, 443, 33),
	(75865, 'WALIAM', 6109, 443, 33),
	(75866, 'BATU PAYUNG', 6110, 443, 33),
	(75867, 'DASRI', 6110, 443, 33),
	(75868, 'KLATIN (KLATIM)', 6110, 443, 33),
	(75869, 'LUWELALA (SUWELALA)', 6110, 443, 33),
	(75870, 'MALADOFOK', 6110, 443, 33),
	(75871, 'SAILALA', 6110, 443, 33),
	(75872, 'SALUK', 6110, 443, 33),
	(75873, 'KASIM', 6111, 443, 33),
	(75874, 'KASIMLE', 6111, 443, 33),
	(75875, 'KLAYAS', 6111, 443, 33),
	(75876, 'MALABAM (MALABAN)', 6111, 443, 33),
	(75877, 'SEGET', 6111, 443, 33),
	(75878, 'WAJIN', 6111, 443, 33),
	(75879, 'WASINGSAN', 6111, 443, 33),
	(75880, 'WAYENKEDE', 6111, 443, 33),
	(75881, 'GISIM', 6112, 443, 33),
	(75882, 'KLASIGON (KLASEGUN)', 6112, 443, 33),
	(75883, 'MAJEMAU', 6112, 443, 33),
	(75884, 'SEGUN', 6112, 443, 33),
	(75885, 'WAIMON', 6112, 443, 33),
	(75886, 'WAINLABAT', 6112, 443, 33),
	(75887, 'KAMPUNG BARU', 6113, 443, 33),
	(75888, 'KLADEMAK', 6113, 443, 33),
	(75889, 'KLAKUBLIK', 6113, 443, 33),
	(75890, 'KLASUUR', 6113, 443, 33),
	(75891, 'REMU UTARA', 6113, 443, 33),
	(75892, 'KLABALA', 6114, 443, 33),
	(75893, 'KLAWASI', 6114, 443, 33),
	(75894, 'RUFEI', 6114, 443, 33),
	(75895, 'SAOKA', 6114, 443, 33),
	(75896, 'TANJUNG KASUARI', 6114, 443, 33),
	(75897, 'DUM BARAT (DOOM BARAT)', 6115, 443, 33),
	(75898, 'DUM TIMUR (DOOM TIMUR)', 6115, 443, 33),
	(75899, 'RAAM', 6115, 443, 33),
	(75900, 'SOOP', 6115, 443, 33),
	(75901, 'KLALIGI', 6116, 443, 33),
	(75902, 'KLASABI', 6116, 443, 33),
	(75903, 'MALABUTOR', 6116, 443, 33),
	(75904, 'MALAWEI', 6116, 443, 33),
	(75905, 'REMU SELATAN', 6116, 443, 33),
	(75906, 'GIWU', 6117, 443, 33),
	(75907, 'KLABLIM', 6117, 443, 33),
	(75908, 'KLAMANA', 6117, 443, 33),
	(75909, 'KLASAMAN', 6117, 443, 33),
	(75910, 'KLASUAT', 6117, 443, 33),
	(75911, 'KLAWALU', 6117, 443, 33),
	(75912, 'KLAWUYUK', 6117, 443, 33),
	(75913, 'KLAGETE', 6118, 443, 33),
	(75914, 'MALAINGKEDI', 6118, 443, 33),
	(75915, 'MALANU', 6118, 443, 33),
	(75916, 'MATALAMAGI (MATALAMANGI)', 6118, 443, 33),
	(75917, 'SAWAGUMU', 6118, 443, 33),
	(75918, 'BEMUS', 6119, 444, 33),
	(75919, 'PASIR PUTIH', 6119, 444, 33),
	(75920, 'WANDUM', 6119, 444, 33),
	(75921, 'WELEK', 6119, 444, 33),
	(75922, 'ISOGO', 6120, 444, 33),
	(75923, 'MATE', 6120, 444, 33),
	(75924, 'MOGIBI', 6120, 444, 33),
	(75925, 'ODEARE', 6120, 444, 33),
	(75926, 'SERKOS', 6120, 444, 33),
	(75927, 'SIBAE', 6120, 444, 33),
	(75928, 'SIRI-SIRI', 6120, 444, 33),
	(75929, 'SOLTA BARU', 6120, 444, 33),
	(75930, 'WADOI', 6120, 444, 33),
	(75931, 'BENAWA I', 6121, 444, 33),
	(75932, 'HAEMARAN (KURANA)', 6121, 444, 33),
	(75933, 'IKANA', 6121, 444, 33),
	(75934, 'KAIS', 6121, 444, 33),
	(75935, 'MAKARORO', 6121, 444, 33),
	(75936, 'MOGOTEMIN (MOGATEMIN)', 6121, 444, 33),
	(75937, 'MUKAMAT', 6121, 444, 33),
	(75938, 'ONIMSEFA', 6121, 444, 33),
	(75939, 'SIRANGGO', 6121, 444, 33),
	(75940, 'SUMANO', 6121, 444, 33),
	(75941, 'TAPURI', 6121, 444, 33),
	(75942, 'YAHADIAN', 6121, 444, 33),
	(75943, 'ARBASINA', 6122, 444, 33),
	(75944, 'BIRAWAKU', 6122, 444, 33),
	(75945, 'DAIMAR', 6122, 444, 33),
	(75946, 'DAUBAK', 6122, 444, 33),
	(75947, 'KASUWERI', 6122, 444, 33),
	(75948, 'KOREWATARA', 6122, 444, 33),
	(75949, 'MIGORI', 6122, 444, 33),
	(75950, 'MIGRITO', 6122, 444, 33),
	(75951, 'NAYAKORE', 6122, 444, 33),
	(75952, 'NEGERI BESAR', 6122, 444, 33),
	(75953, 'SIWATORI', 6122, 444, 33),
	(75954, 'TAMBANI', 6122, 444, 33),
	(75955, 'TAPAS', 6122, 444, 33),
	(75956, 'TAROF', 6122, 444, 33),
	(75957, 'TOPDAN', 6122, 444, 33),
	(75958, 'TOTONA', 6122, 444, 33),
	(75959, 'ADONA', 6123, 444, 33),
	(75960, 'ATORI', 6123, 444, 33),
	(75961, 'BENAWA II', 6123, 444, 33),
	(75962, 'BUBUKO', 6123, 444, 33),
	(75963, 'KAMUNDAN DUA', 6123, 444, 33),
	(75964, 'KAMUNDAN SATU', 6123, 444, 33),
	(75965, 'KARIRIF', 6123, 444, 33),
	(75966, 'KAYUBIRO', 6123, 444, 33),
	(75967, 'UDAGAGA', 6123, 444, 33),
	(75968, 'ALEBO', 6124, 444, 33),
	(75969, 'AMBOLOLI', 6124, 444, 33),
	(75970, 'AMOHALO', 6124, 444, 33),
	(75971, 'AONUPE', 6124, 444, 33),
	(75972, 'BARIAT', 6124, 444, 33),
	(75973, 'CIALAM JAYA', 6124, 444, 33),
	(75974, 'KONDA', 6124, 444, 33),
	(75975, 'KONDA', 6124, 444, 33),
	(75976, 'LALOWIU', 6124, 444, 33),
	(75977, 'LAMBUSA', 6124, 444, 33),
	(75978, 'LAMOMEA', 6124, 444, 33),
	(75979, 'LAWOILA', 6124, 444, 33),
	(75980, 'LEBO JAYA', 6124, 444, 33),
	(75981, 'MANELEK', 6124, 444, 33),
	(75982, 'MASAGENA', 6124, 444, 33),
	(75983, 'MOROME (MORAME)', 6124, 444, 33),
	(75984, 'NAGNA', 6124, 444, 33),
	(75985, 'PAMBULAA JAYA', 6124, 444, 33),
	(75986, 'PUOSU JAYA', 6124, 444, 33),
	(75987, 'RANOWILA', 6124, 444, 33),
	(75988, 'TANEA', 6124, 444, 33),
	(75989, 'WAMARGEGE', 6124, 444, 33),
	(75990, 'WONUA', 6124, 444, 33),
	(75991, 'BEDARE', 6125, 444, 33),
	(75992, 'MUGIM', 6125, 444, 33),
	(75993, 'NUSA', 6125, 444, 33),
	(75994, 'PURAGI', 6125, 444, 33),
	(75995, 'SAGA', 6125, 444, 33),
	(75996, 'TAWANGGIRE', 6125, 444, 33),
	(75997, 'BUMI AJO', 6126, 444, 33),
	(75998, 'HARARO', 6126, 444, 33),
	(75999, 'HASIK JAYA', 6126, 444, 33),
	(76000, 'JOSHIRO', 6126, 444, 33)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 75")
}
