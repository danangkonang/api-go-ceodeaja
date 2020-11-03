package kel10

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Kel17() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
		(17001, 'KUNIR', 1372, 73, 6),
		(17002, 'PIKATAN', 1372, 73, 6),
		(17003, 'REJOSARI', 1372, 73, 6),
		(17004, 'SALAM', 1372, 73, 6),
		(17005, 'TAWANGREJO', 1372, 73, 6),
		(17006, 'WONODADI', 1372, 73, 6),
		(17007, 'GUNUNG GEDE', 1373, 73, 6),
		(17008, 'KALIGRENJENG', 1373, 73, 6),
		(17009, 'NGADIPURO', 1373, 73, 6),
		(17010, 'NGENI', 1373, 73, 6),
		(17011, 'PASIRAMAN', 1373, 73, 6),
		(17012, 'SUMBERBOTO', 1373, 73, 6),
		(17013, 'TAMBAKREJO', 1373, 73, 6),
		(17014, 'WONOTIRTO', 1373, 73, 6),
		(17015, 'BALENREJO', 1374, 74, 6),
		(17016, 'BULAKLO', 1374, 74, 6),
		(17017, 'BULU', 1374, 74, 6),
		(17018, 'KABUNAN', 1374, 74, 6),
		(17019, 'KEDUNGBONDO', 1374, 74, 6),
		(17020, 'KEDUNGDOWO', 1374, 74, 6),
		(17021, 'KEMAMANG', 1374, 74, 6),
		(17022, 'KENEP', 1374, 74, 6),
		(17023, 'LENGKONG', 1374, 74, 6),
		(17024, 'MARGOMULYO', 1374, 74, 6),
		(17025, 'MAYANGKAWIS', 1374, 74, 6),
		(17026, 'MULYOAGUNG', 1374, 74, 6),
		(17027, 'MULYOREJO', 1374, 74, 6),
		(17028, 'NGADILUHUR', 1374, 74, 6),
		(17029, 'PENGANTEN', 1374, 74, 6),
		(17030, 'PILANG GEDE', 1374, 74, 6),
		(17031, 'POH BOGO', 1374, 74, 6),
		(17032, 'PRAMBATAN', 1374, 74, 6),
		(17033, 'SARIREJO', 1374, 74, 6),
		(17034, 'SEKARAN', 1374, 74, 6),
		(17035, 'SIDOBANDUNG', 1374, 74, 6),
		(17036, 'SOBONTORO', 1374, 74, 6),
		(17037, 'SUWALOH', 1374, 74, 6),
		(17038, 'BANJARAN', 1375, 74, 6),
		(17039, 'BANJARANYAR', 1375, 74, 6),
		(17040, 'BAURENO', 1375, 74, 6),
		(17041, 'BLONGSONG', 1375, 74, 6),
		(17042, 'BUMIAYU', 1375, 74, 6),
		(17043, 'DRAJAT', 1375, 74, 6),
		(17044, 'GAJAH', 1375, 74, 6),
		(17045, 'GUNUNGSARI', 1375, 74, 6),
		(17046, 'KADUNGREJO', 1375, 74, 6),
		(17047, 'KALISARI', 1375, 74, 6),
		(17048, 'KARANGDAYU', 1375, 74, 6),
		(17049, 'KAUMAN', 1375, 74, 6),
		(17050, 'LEBAKSARI', 1375, 74, 6),
		(17051, 'NGEMPLAK', 1375, 74, 6),
		(17052, 'PASINAN', 1375, 74, 6),
		(17053, 'POMAHAN', 1375, 74, 6),
		(17054, 'PUCANGARUM', 1375, 74, 6),
		(17055, 'SELOREJO', 1375, 74, 6),
		(17056, 'SEMBUNGLOR', 1375, 74, 6),
		(17057, 'SRATUREJO', 1375, 74, 6),
		(17058, 'SUMURAGUNG', 1375, 74, 6),
		(17059, 'TANGGUNGAN', 1375, 74, 6),
		(17060, 'TLOGOAGUNG', 1375, 74, 6),
		(17061, 'TROJALU', 1375, 74, 6),
		(17062, 'TULUNGAGUNG', 1375, 74, 6),
		(17063, 'BANJAREJO', 1376, 74, 6),
		(17064, 'CAMPURJO', 1376, 74, 6),
		(17065, 'JETAK', 1376, 74, 6),
		(17066, 'KADIPATEN', 1376, 74, 6),
		(17067, 'KALIREJO', 1376, 74, 6),
		(17068, 'KARANG PACAR', 1376, 74, 6),
		(17069, 'KAUMAN', 1376, 74, 6),
		(17070, 'KEPATIHAN', 1376, 74, 6),
		(17071, 'KLANGON', 1376, 74, 6),
		(17072, 'LEDOK KULON', 1376, 74, 6),
		(17073, 'LEDOK WETAN', 1376, 74, 6),
		(17074, 'MOJO KAMPUNG', 1376, 74, 6),
		(17075, 'MULYOAGUNG', 1376, 74, 6),
		(17076, 'NGROWO', 1376, 74, 6),
		(17077, 'PACUL', 1376, 74, 6),
		(17078, 'SEMANDING', 1376, 74, 6),
		(17079, 'SUKOREJO', 1376, 74, 6),
		(17080, 'SUMBANG', 1376, 74, 6),
		(17081, 'BUBULAN', 1377, 74, 6),
		(17082, 'CANCUNG', 1377, 74, 6),
		(17083, 'CLEBUNG', 1377, 74, 6),
		(17084, 'NGOROGUNUNG', 1377, 74, 6),
		(17085, 'SUMBER BENDO', 1377, 74, 6),
		(17086, 'DANDER', 1378, 74, 6),
		(17087, 'GROWOK', 1378, 74, 6),
		(17088, 'JATIBLIMBING', 1378, 74, 6),
		(17089, 'KARANGSONO', 1378, 74, 6),
		(17090, 'KUNCI', 1378, 74, 6),
		(17091, 'MOJORANU', 1378, 74, 6),
		(17092, 'NGABLAK', 1378, 74, 6),
		(17093, 'NGRASEH', 1378, 74, 6),
		(17094, 'NGULANAN', 1378, 74, 6),
		(17095, 'NGUMPAK DALEM', 1378, 74, 6),
		(17096, 'NGUNUT', 1378, 74, 6),
		(17097, 'SENDANGREJO', 1378, 74, 6),
		(17098, 'SUMBER AGUNG', 1378, 74, 6),
		(17099, 'SUMBER ARUM', 1378, 74, 6),
		(17100, 'SUMBER TLASEH', 1378, 74, 6),
		(17101, 'SUMODIKARAN', 1378, 74, 6),
		(17102, 'BRENGGOLO', 1379, 74, 6),
		(17103, 'GREBEGAN', 1379, 74, 6),
		(17104, 'KALITIDU', 1379, 74, 6),
		(17105, 'LERAN', 1379, 74, 6),
		(17106, 'MAYANGGENENG', 1379, 74, 6),
		(17107, 'MAYANGREJO', 1379, 74, 6),
		(17108, 'MLATEN', 1379, 74, 6),
		(17109, 'MOJO', 1379, 74, 6),
		(17110, 'MOJOSARI', 1379, 74, 6),
		(17111, 'NGRINGINREJO', 1379, 74, 6),
		(17112, 'NGUJO', 1379, 74, 6),
		(17113, 'PANJUNAN', 1379, 74, 6),
		(17114, 'PILANGSARI', 1379, 74, 6),
		(17115, 'PUNGPUNGAN', 1379, 74, 6),
		(17116, 'SUKOHARJO', 1379, 74, 6),
		(17117, 'SUMENGKO', 1379, 74, 6),
		(17118, 'TALOK', 1379, 74, 6),
		(17119, 'WOTAN NGARE', 1379, 74, 6),
		(17120, 'BAKUNG', 1380, 74, 6),
		(17121, 'BUNGUR', 1380, 74, 6),
		(17122, 'CANGAKAN', 1380, 74, 6),
		(17123, 'CARUBAN', 1380, 74, 6),
		(17124, 'GEDONGARUM', 1380, 74, 6),
		(17125, 'KABALAN', 1380, 74, 6),
		(17126, 'KANOR', 1380, 74, 6),
		(17127, 'KEDUNGPRIMPEN', 1380, 74, 6),
		(17128, 'NGLARANGAN', 1380, 74, 6),
		(17129, 'PALEMBON', 1380, 74, 6),
		(17130, 'PESEN', 1380, 74, 6),
		(17131, 'PILANG', 1380, 74, 6),
		(17132, 'PIYAK', 1380, 74, 6),
		(17133, 'PRIGI', 1380, 74, 6),
		(17134, 'SAMBERAN', 1380, 74, 6),
		(17135, 'SARANGAN', 1380, 74, 6),
		(17136, 'SEDENG', 1380, 74, 6),
		(17137, 'SEMAMBUNG', 1380, 74, 6),
		(17138, 'SIMBATAN', 1380, 74, 6),
		(17139, 'SIMOREJO', 1380, 74, 6),
		(17140, 'SROYO', 1380, 74, 6),
		(17141, 'SUMBERWANGI', 1380, 74, 6),
		(17142, 'TAMBAHREJO', 1380, 74, 6),
		(17143, 'TEJO', 1380, 74, 6),
		(17144, 'TEMU', 1380, 74, 6),
		(17145, 'BAKALAN', 1381, 74, 6),
		(17146, 'BANGILAN', 1381, 74, 6),
		(17147, 'BENDO', 1381, 74, 6),
		(17148, 'BOGO', 1381, 74, 6),
		(17149, 'KALIANYAR', 1381, 74, 6),
		(17150, 'KAPAS', 1381, 74, 6),
		(17151, 'KEDATON', 1381, 74, 6),
		(17152, 'KLAMPOK', 1381, 74, 6),
		(17153, 'KUMPUL REJO', 1381, 74, 6),
		(17154, 'MOJODESO', 1381, 74, 6),
		(17155, 'NGAMPEL', 1381, 74, 6),
		(17156, 'PADANG MENTOYO', 1381, 74, 6),
		(17157, 'PLESUNGAN', 1381, 74, 6),
		(17158, 'SAMBIROTO', 1381, 74, 6),
		(17159, 'SEMBUNG', 1381, 74, 6),
		(17160, 'SEMEN PINGGIR', 1381, 74, 6),
		(17161, 'SUKOWATI', 1381, 74, 6),
		(17162, 'TANJUNG HARJO', 1381, 74, 6),
		(17163, 'TAPELAN', 1381, 74, 6),
		(17164, 'TIKUSAN', 1381, 74, 6),
		(17165, 'WEDI', 1381, 74, 6),
		(17166, 'BATOKAN', 1382, 74, 6),
		(17167, 'BESAH', 1382, 74, 6),
		(17168, 'BETET', 1382, 74, 6),
		(17169, 'KASIMAN', 1382, 74, 6),
		(17170, 'NGAGLIK', 1382, 74, 6),
		(17171, 'SAMBENG', 1382, 74, 6),
		(17172, 'SEKARAN', 1382, 74, 6),
		(17173, 'SIDOMUKTI', 1382, 74, 6),
		(17174, 'TAMBAKMERAK', 1382, 74, 6),
		(17175, 'TEMBELING', 1382, 74, 6),
		(17176, 'BEJI', 1383, 74, 6),
		(17177, 'HARGOMULYO', 1383, 74, 6),
		(17178, 'KAWENGAN', 1383, 74, 6),
		(17179, 'KEDEWAN', 1383, 74, 6),
		(17180, 'WONOCOLO', 1383, 74, 6),
		(17181, 'BABAD', 1384, 74, 6),
		(17182, 'BALONGCABE', 1384, 74, 6),
		(17183, 'DAYUKIDUL', 1384, 74, 6),
		(17184, 'DROKILO', 1384, 74, 6),
		(17185, 'DUWEL', 1384, 74, 6),
		(17186, 'GEGER', 1384, 74, 6),
		(17187, 'JAMBEREJO', 1384, 74, 6),
		(17188, 'KEDUNGADEM', 1384, 74, 6),
		(17189, 'KEDUNGREJO', 1384, 74, 6),
		(17190, 'KENDUNG', 1384, 74, 6),
		(17191, 'KEPOH KIDUL', 1384, 74, 6),
		(17192, 'KESONGO', 1384, 74, 6),
		(17193, 'MEGALE', 1384, 74, 6),
		(17194, 'MLIDEG', 1384, 74, 6),
		(17195, 'MOJOREJO', 1384, 74, 6),
		(17196, 'NGRANDU', 1384, 74, 6),
		(17197, 'PANJANG', 1384, 74, 6),
		(17198, 'PEJOK', 1384, 74, 6),
		(17199, 'SIDOMULYO', 1384, 74, 6),
		(17200, 'SIDOREJO', 1384, 74, 6),
		(17201, 'TLOGOAGUNG', 1384, 74, 6),
		(17202, 'TONDOMULO', 1384, 74, 6),
		(17203, 'TUMBRASANOM', 1384, 74, 6),
		(17204, 'BALONGDOWO', 1385, 74, 6),
		(17205, 'BAYEMGEDE', 1385, 74, 6),
		(17206, 'BETET', 1385, 74, 6),
		(17207, 'BRANGKAL', 1385, 74, 6),
		(17208, 'BUMIREJO', 1385, 74, 6),
		(17209, 'CENGKIR', 1385, 74, 6),
		(17210, 'JIPO', 1385, 74, 6),
		(17211, 'KARANGAN', 1385, 74, 6),
		(17212, 'KEPOH', 1385, 74, 6),
		(17213, 'KRANGKONG', 1385, 74, 6),
		(17214, 'MOJOSARI', 1385, 74, 6),
		(17215, 'MUDUNG', 1385, 74, 6),
		(17216, 'NGLUMBER', 1385, 74, 6),
		(17217, 'NGRANGGON ANYAR', 1385, 74, 6),
		(17218, 'PEJOK', 1385, 74, 6),
		(17219, 'POHWATES', 1385, 74, 6),
		(17220, 'SIDOMUKTI', 1385, 74, 6),
		(17221, 'SIMOREJO', 1385, 74, 6),
		(17222, 'SUGIHWARAS', 1385, 74, 6),
		(17223, 'SUMBERAGUNG', 1385, 74, 6),
		(17224, 'SUMBERGEDE', 1385, 74, 6),
		(17225, 'SUMBEROTO', 1385, 74, 6),
		(17226, 'TLOGOREJO', 1385, 74, 6),
		(17227, 'TURIGEDE', 1385, 74, 6),
		(17228, 'WORO', 1385, 74, 6),
		(17229, 'BANARAN', 1386, 74, 6),
		(17230, 'DUKUH LOR', 1386, 74, 6),
		(17231, 'KACANGAN', 1386, 74, 6),
		(17232, 'KEDUNGREJO', 1386, 74, 6),
		(17233, 'KEMIRI', 1386, 74, 6),
		(17234, 'KETILENG', 1386, 74, 6),
		(17235, 'KLITEH', 1386, 74, 6),
		(17236, 'MALO', 1386, 74, 6),
		(17237, 'NGUJUNG', 1386, 74, 6),
		(17238, 'PETAK', 1386, 74, 6),
		(17239, 'RENDENG', 1386, 74, 6),
		(17240, 'SEMLARAN', 1386, 74, 6),
		(17241, 'SUDAH', 1386, 74, 6),
		(17242, 'SUKOREJO', 1386, 74, 6),
		(17243, 'SUMBEREJO', 1386, 74, 6),
		(17244, 'TAMBAKROMO', 1386, 74, 6),
		(17245, 'TANGGIR', 1386, 74, 6),
		(17246, 'TINAWUN', 1386, 74, 6),
		(17247, 'TREMBES', 1386, 74, 6),
		(17248, 'TULUNGAGUNG', 1386, 74, 6),
		(17249, 'GENENG', 1387, 74, 6),
		(17250, 'KALANGAN', 1387, 74, 6),
		(17251, 'MARGOMULYO', 1387, 74, 6),
		(17252, 'MEDURI', 1387, 74, 6),
		(17253, 'NGELO', 1387, 74, 6),
		(17254, 'SUMBEREJO', 1387, 74, 6),
		(17255, 'BONDOL', 1388, 74, 6),
		(17256, 'KARANGMANGU', 1388, 74, 6),
		(17257, 'NGAMBON', 1388, 74, 6),
		(17258, 'NGLAMPIN', 1388, 74, 6),
		(17259, 'SENGON', 1388, 74, 6),
		(17260, 'BANCER', 1389, 74, 6),
		(17261, 'BLIMBING GEDE', 1389, 74, 6),
		(17262, 'JUMOK', 1389, 74, 6),
		(17263, 'KALIREJO', 1389, 74, 6),
		(17264, 'KLEMPUN', 1389, 74, 6),
		(17265, 'LUWIHAJI', 1389, 74, 6),
		(17266, 'MOJOREJO', 1389, 74, 6),
		(17267, 'NGANTI', 1389, 74, 6),
		(17268, 'NGRAHO', 1389, 74, 6),
		(17269, 'PANDAN', 1389, 74, 6),
		(17270, 'PAYAMAN', 1389, 74, 6),
		(17271, 'SUGIH WARAS', 1389, 74, 6),
		(17272, 'SUMBERAGUNG', 1389, 74, 6),
		(17273, 'SUMBERARUM', 1389, 74, 6),
		(17274, 'TANGGUNGAN', 1389, 74, 6),
		(17275, 'TAPELAN', 1389, 74, 6),
		(17276, 'BANJARJO', 1390, 74, 6),
		(17277, 'CENDONO', 1390, 74, 6),
		(17278, 'DENGOK', 1390, 74, 6),
		(17279, 'KEBUNAGUNG', 1390, 74, 6),
		(17280, 'KENDUNG', 1390, 74, 6),
		(17281, 'KUNCEN', 1390, 74, 6),
		(17282, 'NGASINAN', 1390, 74, 6),
		(17283, 'NGEPER', 1390, 74, 6),
		(17284, 'NGRADIN', 1390, 74, 6),
		(17285, 'NGUKEN', 1390, 74, 6),
		(17286, 'PADANGAN', 1390, 74, 6),
		(17287, 'PRANGI', 1390, 74, 6),
		(17288, 'PURWOREJO', 1390, 74, 6),
		(17289, 'SIDOREJO', 1390, 74, 6),
		(17290, 'SONOREJO', 1390, 74, 6),
		(17291, 'TEBON', 1390, 74, 6),
		(17292, 'BARENG', 1391, 74, 6),
		(17293, 'BOBOL', 1391, 74, 6),
		(17294, 'DELING', 1391, 74, 6),
		(17295, 'KLINO', 1391, 74, 6),
		(17296, 'MIYONO', 1391, 74, 6),
		(17297, 'SEKAR', 1391, 74, 6),
		(17298, 'ALASGUNG', 1392, 74, 6),
		(17299, 'BALONGREJO', 1392, 74, 6),
		(17300, 'BARENG', 1392, 74, 6),
		(17301, 'BULU', 1392, 74, 6),
		(17302, 'DRENGES', 1392, 74, 6),
		(17303, 'GENJOR', 1392, 74, 6),
		(17304, 'GLAGAH WANGI', 1392, 74, 6),
		(17305, 'GLAGAHAN', 1392, 74, 6),
		(17306, 'JATITENGAH', 1392, 74, 6),
		(17307, 'KEDUNGDOWO', 1392, 74, 6),
		(17308, 'NGLAJANG', 1392, 74, 6),
		(17309, 'PANEMON', 1392, 74, 6),
		(17310, 'PANUNGGALAN', 1392, 74, 6),
		(17311, 'SIWALAN', 1392, 74, 6),
		(17312, 'SUGIHWARAS', 1392, 74, 6),
		(17313, 'TRATE', 1392, 74, 6),
		(17314, 'WEDORO', 1392, 74, 6),
		(17315, 'DUYUNGAN', 1393, 74, 6),
		(17316, 'JUMPUT', 1393, 74, 6),
		(17317, 'KALICILIK', 1393, 74, 6),
		(17318, 'KLEPEK', 1393, 74, 6),
		(17319, 'PACING', 1393, 74, 6),
		(17320, 'PURWOASRI', 1393, 74, 6),
		(17321, 'SEMAWOT', 1393, 74, 6),
		(17322, 'SEMEN KIDUL', 1393, 74, 6),
		(17323, 'SIDODADI', 1393, 74, 6),
		(17324, 'SIDOREJO', 1393, 74, 6),
		(17325, 'SITIAJI', 1393, 74, 6),
		(17326, 'SUKOSEWU', 1393, 74, 6),
		(17327, 'SUMBEREJO KIDUL', 1393, 74, 6),
		(17328, 'TEGALKODO', 1393, 74, 6),
		(17329, 'BANJARJO', 1394, 74, 6),
		(17330, 'BOGANGIN', 1394, 74, 6),
		(17331, 'BUTOH', 1394, 74, 6),
		(17332, 'DERU', 1394, 74, 6),
		(17333, 'JATIGEDE', 1394, 74, 6),
		(17334, 'KARANG DINOYO', 1394, 74, 6),
		(17335, 'KARANGDOWO', 1394, 74, 6),
		(17336, 'KAYULEMAH', 1394, 74, 6),
		(17337, 'KEDUNGREJO', 1394, 74, 6),
		(17338, 'MARGOAGUNG', 1394, 74, 6),
		(17339, 'MEJUWET', 1394, 74, 6),
		(17340, 'MLINJENG', 1394, 74, 6),
		(17341, 'NGAMPAL', 1394, 74, 6),
		(17342, 'PEJAMBON', 1394, 74, 6),
		(17343, 'PEKUWON', 1394, 74, 6),
		(17344, 'PRAYUNGAN', 1394, 74, 6),
		(17345, 'SAMBONGREJO', 1394, 74, 6),
		(17346, 'SENDANGAGUNG', 1394, 74, 6),
		(17347, 'SUMBERHARJO', 1394, 74, 6),
		(17348, 'SUMBERREJO', 1394, 74, 6),
		(17349, 'SUMURAGUNG', 1394, 74, 6),
		(17350, 'TALUN', 1394, 74, 6),
		(17351, 'TELENG', 1394, 74, 6),
		(17352, 'TLOGOHAJI', 1394, 74, 6),
		(17353, 'TULUNGREJO', 1394, 74, 6),
		(17354, 'WOTAN', 1394, 74, 6),
		(17355, 'BAKALAN', 1395, 74, 6),
		(17356, 'DOLOK GEDE', 1395, 74, 6),
		(17357, 'GADING', 1395, 74, 6),
		(17358, 'GAMONGAN', 1395, 74, 6),
		(17359, 'JATIMULYO', 1395, 74, 6),
		(17360, 'JAWIK', 1395, 74, 6),
		(17361, 'KACANGAN', 1395, 74, 6),
		(17362, 'KALISUMBER', 1395, 74, 6),
		(17363, 'MALING MATI', 1395, 74, 6),
		(17364, 'MULYOREJO', 1395, 74, 6),
		(17365, 'NAPIS', 1395, 74, 6),
		(17366, 'NGRANCANG', 1395, 74, 6),
		(17367, 'PENGKOL', 1395, 74, 6),
		(17368, 'SENDANGREJO', 1395, 74, 6),
		(17369, 'SUKOREJO', 1395, 74, 6),
		(17370, 'TAMBAK REJO', 1395, 74, 6),
		(17371, 'TANJUNG', 1395, 74, 6),
		(17372, 'TURI', 1395, 74, 6),
		(17373, 'BAKULAN', 1396, 74, 6),
		(17374, 'BELUN', 1396, 74, 6),
		(17375, 'BUNTALAN', 1396, 74, 6),
		(17376, 'JONO', 1396, 74, 6),
		(17377, 'KEDUNGSARI', 1396, 74, 6),
		(17378, 'KEDUNGSUMBER', 1396, 74, 6),
		(17379, 'NGUJUNG', 1396, 74, 6),
		(17380, 'PANCUR', 1396, 74, 6),
		(17381, 'PANDANTOYO', 1396, 74, 6),
		(17382, 'PAPRINGAN', 1396, 74, 6),
		(17383, 'SOKO', 1396, 74, 6),
		(17384, 'TEMAYANG', 1396, 74, 6),
		(17385, 'BANJARSARI', 1397, 74, 6),
		(17386, 'BERO', 1397, 74, 6),
		(17387, 'GADEN', 1397, 74, 6),
		(17388, 'GUYANGAN', 1397, 74, 6),
		(17389, 'JATIPURO', 1397, 74, 6),
		(17390, 'KALIKEBO', 1397, 74, 6),
		(17391, 'KANDANGAN', 1397, 74, 6),
		(17392, 'KANTEN', 1397, 74, 6),
		(17393, 'KARANGPAKEL', 1397, 74, 6),
		(17394, 'KRADENAN', 1397, 74, 6),
		(17395, 'MANDONG', 1397, 74, 6),
		(17396, 'MIRENG', 1397, 74, 6),
		(17397, 'MORI', 1397, 74, 6),
		(17398, 'PADANG', 1397, 74, 6),
		(17399, 'PAGERWESI', 1397, 74, 6),
		(17400, 'PALAR', 1397, 74, 6),
		(17401, 'PLANGGU', 1397, 74, 6),
		(17402, 'PULUHAN', 1397, 74, 6),
		(17403, 'PUNDUNGSARI', 1397, 74, 6),
		(17404, 'SABRANG LOR', 1397, 74, 6),
		(17405, 'SAJEN', 1397, 74, 6),
		(17406, 'SRANAK', 1397, 74, 6),
		(17407, 'SUMBANG TIMUN', 1397, 74, 6),
		(17408, 'SUMBER', 1397, 74, 6),
		(17409, 'SUMBERREJO', 1397, 74, 6),
		(17410, 'TRUCUK', 1397, 74, 6),
		(17411, 'TRUCUK', 1397, 74, 6),
		(17412, 'TULUNGREJO', 1397, 74, 6),
		(17413, 'WANGLU', 1397, 74, 6),
		(17414, 'WONOSARI (WONOSOSA)', 1397, 74, 6),
		(17415, 'BARATAN', 1398, 75, 6),
		(17416, 'BENDELAN', 1398, 75, 6),
		(17417, 'BINAKAL', 1398, 75, 6),
		(17418, 'GADINGSARI', 1398, 75, 6),
		(17419, 'JERUKSOKSOK', 1398, 75, 6),
		(17420, 'KEMBANGAN', 1398, 75, 6),
		(17421, 'SUMBERTENGAH', 1398, 75, 6),
		(17422, 'SUMBERWARU', 1398, 75, 6),
		(17423, 'BADEAN', 1399, 75, 6),
		(17424, 'BLINDUNGAN', 1399, 75, 6),
		(17425, 'DABASAH', 1399, 75, 6),
		(17426, 'KADEMANGAN', 1399, 75, 6),
		(17427, 'KEMBANG', 1399, 75, 6),
		(17428, 'KOTAKULON', 1399, 75, 6),
		(17429, 'NANGKAAN', 1399, 75, 6),
		(17430, 'PANCORAN', 1399, 75, 6),
		(17431, 'PEJATEN', 1399, 75, 6),
		(17432, 'SUKOWIRYO', 1399, 75, 6),
		(17433, 'TAMANSARI', 1399, 75, 6),
		(17434, 'BOTOLINGGO', 1400, 75, 6),
		(17435, 'GAYAM', 1400, 75, 6),
		(17436, 'GAYAM LOR', 1400, 75, 6),
		(17437, 'KLEKEHAN (KLEKEAN)', 1400, 75, 6),
		(17438, 'LANAS', 1400, 75, 6),
		(17439, 'LUMUTAN', 1400, 75, 6),
		(17440, 'PENANG', 1400, 75, 6),
		(17441, 'SUMBERCANTING', 1400, 75, 6),
		(17442, 'BAJURAN', 1401, 75, 6),
		(17443, 'BATU AMPAR', 1401, 75, 6),
		(17444, 'BATU SALANG', 1401, 75, 6),
		(17445, 'BERCAK', 1401, 75, 6),
		(17446, 'BERCAK ASRI', 1401, 75, 6),
		(17447, 'CERMEE', 1401, 75, 6),
		(17448, 'GRUJUGAN', 1401, 75, 6),
		(17449, 'JIREK MAS', 1401, 75, 6),
		(17450, 'KLADI', 1401, 75, 6),
		(17451, 'PELALANGAN', 1401, 75, 6),
		(17452, 'RAMBAN KULON', 1401, 75, 6),
		(17453, 'RAMBAN WETAN', 1401, 75, 6),
		(17454, 'SOLOR', 1401, 75, 6),
		(17455, 'SULING KULON', 1401, 75, 6),
		(17456, 'SULING WETAN', 1401, 75, 6),
		(17457, 'CURAHDAMI', 1402, 75, 6),
		(17458, 'CURAHPOH', 1402, 75, 6),
		(17459, 'JETIS', 1402, 75, 6),
		(17460, 'KUPANG', 1402, 75, 6),
		(17461, 'LOCARE', 1402, 75, 6),
		(17462, 'PAKUWESI', 1402, 75, 6),
		(17463, 'PENAMBANGAN', 1402, 75, 6),
		(17464, 'PETUNG', 1402, 75, 6),
		(17465, 'PONCOGATI', 1402, 75, 6),
		(17466, 'SILOLEMBU', 1402, 75, 6),
		(17467, 'SUMBER SALAK', 1402, 75, 6),
		(17468, 'SUMBERSUKO', 1402, 75, 6),
		(17469, 'DADAPAN', 1403, 75, 6),
		(17470, 'DAWUHAN', 1403, 75, 6),
		(17471, 'GRUJUGAN KIDUL', 1403, 75, 6),
		(17472, 'KABUARAN', 1403, 75, 6),
		(17473, 'KEJAWAN', 1403, 75, 6),
		(17474, 'PEKAUMAN', 1403, 75, 6),
		(17475, 'SUMBERPANDAN', 1403, 75, 6),
		(17476, 'TAMAN', 1403, 75, 6),
		(17477, 'TEGALMIJIN', 1403, 75, 6),
		(17478, 'WANISODO', 1403, 75, 6),
		(17479, 'WONOSARI', 1403, 75, 6),
		(17480, 'GRUJUGAN LOR', 1404, 75, 6),
		(17481, 'JAMBE ANOM', 1404, 75, 6),
		(17482, 'JAMBE SARI', 1404, 75, 6),
		(17483, 'PEJAGAN', 1404, 75, 6),
		(17484, 'PENGARANG', 1404, 75, 6),
		(17485, 'PUCANG ANOM', 1404, 75, 6),
		(17486, 'SUMBER ANYAR', 1404, 75, 6),
		(17487, 'SUMBER JERUK', 1404, 75, 6),
		(17488, 'TEGALPASIR', 1404, 75, 6),
		(17489, 'BESUK', 1405, 75, 6),
		(17490, 'BLIMBING', 1405, 75, 6),
		(17491, 'KARANG SENGON', 1405, 75, 6),
		(17492, 'KARANGANYAR', 1405, 75, 6),
		(17493, 'KLABANG', 1405, 75, 6),
		(17494, 'KLAMPOKAN', 1405, 75, 6),
		(17495, 'LEPRAK', 1405, 75, 6),
		(17496, 'PANDAK', 1405, 75, 6),
		(17497, 'SUMBERSUKO', 1405, 75, 6),
		(17498, 'WONOBOYO', 1405, 75, 6),
		(17499, 'WONOKERTO', 1405, 75, 6),
		(17500, 'GAMBANGAN', 1406, 75, 6),
		(17501, 'GUNUNG SARI', 1406, 75, 6),
		(17502, 'MAESAN', 1406, 75, 6),
		(17503, 'PAKUNIRAN', 1406, 75, 6),
		(17504, 'PENANGGUNGAN', 1406, 75, 6),
		(17505, 'PUJERBARU', 1406, 75, 6),
		(17506, 'SUCOLOR', 1406, 75, 6),
		(17507, 'SUGER LOR', 1406, 75, 6),
		(17508, 'SUMBER ANYAR', 1406, 75, 6),
		(17509, 'SUMBERPAKEM', 1406, 75, 6),
		(17510, 'SUMBERSARI', 1406, 75, 6),
		(17511, 'TANAHWULAN', 1406, 75, 6),
		(17512, 'ANDUNGSARI', 1407, 75, 6),
		(17513, 'ARDISAENG', 1407, 75, 6),
		(17514, 'CANDI BINANGUN', 1407, 75, 6),
		(17515, 'GADINGSARI', 1407, 75, 6),
		(17516, 'HARGO BINANGUN', 1407, 75, 6),
		(17517, 'HARJO BINANGUN', 1407, 75, 6),
		(17518, 'KUPANG', 1407, 75, 6),
		(17519, 'PAKEM', 1407, 75, 6),
		(17520, 'PAKEM BINANGUN', 1407, 75, 6),
		(17521, 'PATEMON', 1407, 75, 6),
		(17522, 'PETUNG', 1407, 75, 6),
		(17523, 'PURWO BINANGUN', 1407, 75, 6),
		(17524, 'SUMBER DUMPYONG', 1407, 75, 6),
		(17525, 'BANDILAN', 1408, 75, 6),
		(17526, 'CANGKRING', 1408, 75, 6),
		(17527, 'PRAJEKAN KIDUL', 1408, 75, 6),
		(17528, 'PRAJEKAN LOR', 1408, 75, 6),
		(17529, 'SEMPOL', 1408, 75, 6),
		(17530, 'TARUM', 1408, 75, 6),
		(17531, 'WALIDONO', 1408, 75, 6),
		(17532, 'ALASSUMUR', 1409, 75, 6),
		(17533, 'KEJAYAN', 1409, 75, 6),
		(17534, 'MANGLI', 1409, 75, 6),
		(17535, 'MASKUNING KULON', 1409, 75, 6),
		(17536, 'MASKUNING WETAN', 1409, 75, 6),
		(17537, 'MENGOK', 1409, 75, 6),
		(17538, 'PADASAN', 1409, 75, 6),
		(17539, 'RANDUCANGKRING', 1409, 75, 6),
		(17540, 'SUKODONO', 1409, 75, 6),
		(17541, 'SUKOKERTO', 1409, 75, 6),
		(17542, 'SUKOWONO', 1409, 75, 6),
		(17543, 'JAMPIT', 1410, 75, 6),
		(17544, 'KALIANYAR', 1410, 75, 6),
		(17545, 'KALIGEDANG', 1410, 75, 6),
		(17546, 'KALISAT', 1410, 75, 6),
		(17547, 'SEMPOL', 1410, 75, 6),
		(17548, 'SUMBER REJO', 1410, 75, 6),
		(17549, 'KERANG', 1411, 75, 6),
		(17550, 'NOGOSARI', 1411, 75, 6),
		(17551, 'PECALONGAN', 1411, 75, 6),
		(17552, 'SUKOSARI LOR', 1411, 75, 6),
		(17553, 'REJOAGUNG', 1412, 75, 6),
		(17554, 'SUKOREJO', 1412, 75, 6),
		(17555, 'SUKOSARI KIDUL', 1412, 75, 6),
		(17556, 'SUMBER GADING', 1412, 75, 6),
		(17557, 'SUMBER WRINGIN', 1412, 75, 6),
		(17558, 'TEGALJATI', 1412, 75, 6),
		(17559, 'GENTONG', 1413, 75, 6),
		(17560, 'KEMUNINGAN', 1413, 75, 6),
		(17561, 'KRETEK', 1413, 75, 6),
		(17562, 'PAGUAN', 1413, 75, 6),
		(17563, 'SUMBERKOKAP', 1413, 75, 6),
		(17564, 'TAMAN', 1413, 75, 6),
		(17565, 'TREBUNGAN', 1413, 75, 6),
		(17566, 'KALIANYAR', 1414, 75, 6),
		(17567, 'KARANGMELOK', 1414, 75, 6),
		(17568, 'KEMIRIAN', 1414, 75, 6),
		(17569, 'MENGEN', 1414, 75, 6),
		(17570, 'SUKOSARI', 1414, 75, 6),
		(17571, 'SUMBER ANOM', 1414, 75, 6),
		(17572, 'SUMBERKEMUNING', 1414, 75, 6),
		(17573, 'TAMANAN', 1414, 75, 6),
		(17574, 'WONOSUKO', 1414, 75, 6),
		(17575, 'CINDOGO', 1415, 75, 6),
		(17576, 'GUNUNGANYAR', 1415, 75, 6),
		(17577, 'JURANGSAPI', 1415, 75, 6),
		(17578, 'KALITAPEN', 1415, 75, 6),
		(17579, 'MANGLI WETAN', 1415, 75, 6),
		(17580, 'MRAWAN', 1415, 75, 6),
		(17581, 'TAAL', 1415, 75, 6),
		(17582, 'TAPEN', 1415, 75, 6),
		(17583, 'WONOKUSUMO', 1415, 75, 6),
		(17584, 'KARANGANYAR', 1416, 75, 6),
		(17585, 'KLABANG', 1416, 75, 6),
		(17586, 'KLABANG AGUNG', 1416, 75, 6),
		(17587, 'MANDIRO', 1416, 75, 6),
		(17588, 'PURNAMA', 1416, 75, 6),
		(17589, 'SEKARPUTIH', 1416, 75, 6),
		(17590, 'TANGGULANGIN', 1416, 75, 6),
		(17591, 'TEGALAMPEL', 1416, 75, 6),
		(17592, 'BATAAN', 1417, 75, 6),
		(17593, 'DAWUHAN', 1417, 75, 6),
		(17594, 'GEBANG', 1417, 75, 6),
		(17595, 'KAJAR', 1417, 75, 6),
		(17596, 'KESEMEK (KASEMEK)', 1417, 75, 6),
		(17597, 'KONCER DARUL ALAM', 1417, 75, 6),
		(17598, 'KONCER KIDUL', 1417, 75, 6),
		(17599, 'LOJAJAR', 1417, 75, 6),
		(17600, 'PEKALANGAN', 1417, 75, 6),
		(17601, 'SUMBERSALAM', 1417, 75, 6),
		(17602, 'TANGSIL KULON', 1417, 75, 6),
		(17603, 'TENGGARANG', 1417, 75, 6),
		(17604, 'BRAMBANG DARUSSOLAH', 1418, 75, 6),
		(17605, 'GUNOSARI', 1418, 75, 6),
		(17606, 'JEBUNG KIDUL', 1418, 75, 6),
		(17607, 'JEBUNG LOR', 1418, 75, 6),
		(17608, 'KEMBANG', 1418, 75, 6),
		(17609, 'PAKISAN', 1418, 75, 6),
		(17610, 'PATEMON', 1418, 75, 6),
		(17611, 'SULEK', 1418, 75, 6),
		(17612, 'TLOGOSARI', 1418, 75, 6),
		(17613, 'TROTOSARI', 1418, 75, 6),
		(17614, 'AMBULU', 1419, 75, 6),
		(17615, 'AMPELAN', 1419, 75, 6),
		(17616, 'BANYUPUTIH', 1419, 75, 6),
		(17617, 'BANYUWULU', 1419, 75, 6),
		(17618, 'BUKOR', 1419, 75, 6),
		(17619, 'GLINGSERAN', 1419, 75, 6),
		(17620, 'GUBRIH', 1419, 75, 6),
		(17621, 'JAMBEWUNGU', 1419, 75, 6),
		(17622, 'JATISARI', 1419, 75, 6),
		(17623, 'JATITAMBAN', 1419, 75, 6),
		(17624, 'SUMBERCANTING', 1419, 75, 6),
		(17625, 'SUMBERMALANG', 1419, 75, 6),
		(17626, 'WRINGIN', 1419, 75, 6),
		(17627, 'BABATAN', 1420, 76, 6),
		(17628, 'BALONGPANGGANG', 1420, 76, 6),
		(17629, 'BANDUNGSEKARAN', 1420, 76, 6),
		(17630, 'BANJARAGUNG', 1420, 76, 6),
		(17631, 'BRANGKAL', 1420, 76, 6),
		(17632, 'DAPET', 1420, 76, 6),
		(17633, 'DOHOAGUNG', 1420, 76, 6),
		(17634, 'GANGGANG', 1420, 76, 6),
		(17635, 'JOMBANGDELIK', 1420, 76, 6),
		(17636, 'KARANGSEMANDING', 1420, 76, 6),
		(17637, 'KEDUNGPRING', 1420, 76, 6),
		(17638, 'KEDUNGSUMBER', 1420, 76, 6),
		(17639, 'KLOTOK', 1420, 76, 6),
		(17640, 'MOJOGEDE', 1420, 76, 6),
		(17641, 'NGAMPEL', 1420, 76, 6),
		(17642, 'NGASIN', 1420, 76, 6),
		(17643, 'PACUH', 1420, 76, 6),
		(17644, 'PINGGIR', 1420, 76, 6),
		(17645, 'PUCUNG', 1420, 76, 6),
		(17646, 'SEKARPUTIH', 1420, 76, 6),
		(17647, 'TANAHLANDEAN', 1420, 76, 6),
		(17648, 'TENGGOR', 1420, 76, 6),
		(17649, 'WAHAS', 1420, 76, 6),
		(17650, 'WONOREJO', 1420, 76, 6),
		(17651, 'WOTANSARI', 1420, 76, 6),
		(17652, 'BALONGMOJO', 1421, 76, 6),
		(17653, 'BALONGTUNJUNG', 1421, 76, 6),
		(17654, 'BANTER', 1421, 76, 6),
		(17655, 'BENGKELOLOR', 1421, 76, 6),
		(17656, 'BULANGKULON', 1421, 76, 6),
		(17657, 'BULUREJO', 1421, 76, 6),
		(17658, 'DELIKSUMBER', 1421, 76, 6),
		(17659, 'DERMO', 1421, 76, 6),
		(17660, 'GLURANPLOSO', 1421, 76, 6),
		(17661, 'JATIREMBE', 1421, 76, 6),
		(17662, 'JOGODALU', 1421, 76, 6),
		(17663, 'KALIPADANG', 1421, 76, 6),
		(17664, 'KARANGANKIDUL', 1421, 76, 6),
		(17665, 'KEDUNGRUKEM', 1421, 76, 6),
		(17666, 'KEDUNGSEKAR', 1421, 76, 6),
		(17667, 'KLAMPOK', 1421, 76, 6),
		(17668, 'LUNDO', 1421, 76, 6),
		(17669, 'METATU', 1421, 76, 6),
		(17670, 'MUNGGUGEBANG', 1421, 76, 6),
		(17671, 'MUNGGUGIANTI', 1421, 76, 6),
		(17672, 'PUNDUTTRATE', 1421, 76, 6),
		(17673, 'SEDAPURKLAGEN', 1421, 76, 6),
		(17674, 'SIRNOBOYO', 1421, 76, 6),
		(17675, 'ABAR-ABIR', 1422, 76, 6),
		(17676, 'BEDANTEN', 1422, 76, 6),
		(17677, 'BUNGAH', 1422, 76, 6),
		(17678, 'GUMENG', 1422, 76, 6),
		(17679, 'INDRODELIK', 1422, 76, 6),
		(17680, 'KEMANGI', 1422, 76, 6),
		(17681, 'KISIK', 1422, 76, 6),
		(17682, 'KRAMAT', 1422, 76, 6),
		(17683, 'MASANGAN', 1422, 76, 6),
		(17684, 'MELIRANG', 1422, 76, 6),
		(17685, 'MOJOPURO GEDE', 1422, 76, 6),
		(17686, 'MOJOPURO WETAN', 1422, 76, 6),
		(17687, 'PEGUNDAN', 1422, 76, 6),
		(17688, 'RACIWETAN', 1422, 76, 6),
		(17689, 'SIDOKUMPUL', 1422, 76, 6),
		(17690, 'SIDOMUKTI', 1422, 76, 6),
		(17691, 'SIDOREJO', 1422, 76, 6),
		(17692, 'SUKOREJO', 1422, 76, 6),
		(17693, 'SUKOWATI', 1422, 76, 6),
		(17694, 'SUNGONLEGOWO', 1422, 76, 6),
		(17695, 'TAJUNG WIDORO', 1422, 76, 6),
		(17696, 'WATUAGUNG', 1422, 76, 6),
		(17697, 'BANJARSARI', 1423, 76, 6),
		(17698, 'BETITING', 1423, 76, 6),
		(17699, 'CAGAKAGUNG', 1423, 76, 6),
		(17700, 'CERME KIDUL', 1423, 76, 6),
		(17701, 'CERME LOR', 1423, 76, 6),
		(17702, 'DADAPKUNING', 1423, 76, 6),
		(17703, 'DAMPAAN', 1423, 76, 6),
		(17704, 'DOORO', 1423, 76, 6),
		(17705, 'DUNGUS', 1423, 76, 6),
		(17706, 'GEDANGKULUT', 1423, 76, 6),
		(17707, 'GURANGANYAR', 1423, 76, 6),
		(17708, 'IKERIKERGEGER', 1423, 76, 6),
		(17709, 'JONO', 1423, 76, 6),
		(17710, 'KAMBINGAN', 1423, 76, 6),
		(17711, 'KANDANGAN', 1423, 76, 6),
		(17712, 'LENGKONG', 1423, 76, 6),
		(17713, 'MOROWUDI', 1423, 76, 6),
		(17714, 'NGABETAN', 1423, 76, 6),
		(17715, 'NGEMBUNG', 1423, 76, 6),
		(17716, 'PADEG', 1423, 76, 6),
		(17717, 'PANDU', 1423, 76, 6),
		(17718, 'SEMAMPIR', 1423, 76, 6),
		(17719, 'SUKOANYAR', 1423, 76, 6),
		(17720, 'TAMBAKBERAS', 1423, 76, 6),
		(17721, 'WEDANI', 1423, 76, 6),
		(17722, 'BAMBE', 1424, 76, 6),
		(17723, 'BANJARAN', 1424, 76, 6),
		(17724, 'CANGKIR', 1424, 76, 6),
		(17725, 'DRIYOREJO', 1424, 76, 6),
		(17726, 'GADUNG', 1424, 76, 6),
		(17727, 'KARANGANDONG', 1424, 76, 6),
		(17728, 'KESAMBENWETAN', 1424, 76, 6),
		(17729, 'KRIKILAN', 1424, 76, 6),
		(17730, 'MOJOSARIREJO', 1424, 76, 6),
		(17731, 'MULUNG', 1424, 76, 6),
		(17732, 'PETIKEN', 1424, 76, 6),
		(17733, 'RANDEGANSARI', 1424, 76, 6),
		(17734, 'SUMPUT', 1424, 76, 6),
		(17735, 'TANJUNGAN', 1424, 76, 6),
		(17736, 'TENARU', 1424, 76, 6),
		(17737, 'WEDOROANOM', 1424, 76, 6),
		(17738, 'AMBENG AMBENG WATANGREJO', 1425, 76, 6),
		(17739, 'BENDUNGAN', 1425, 76, 6),
		(17740, 'DUDUK SAMPEYAN', 1425, 76, 6),
		(17741, 'GLANGGANG', 1425, 76, 6),
		(17742, 'GREDEK', 1425, 76, 6),
		(17743, 'KANDANGAN', 1425, 76, 6),
		(17744, 'KAWISTOWINDU', 1425, 76, 6),
		(17745, 'KEMUDI', 1425, 76, 6),
		(17746, 'KRAMAT KULON', 1425, 76, 6),
		(17747, 'PALEBON', 1425, 76, 6),
		(17748, 'PANDANAN', 1425, 76, 6),
		(17749, 'PANJUNAN', 1425, 76, 6),
		(17750, 'PETISBENEM', 1425, 76, 6),
		(17751, 'SAMIRPLAPAN', 1425, 76, 6),
		(17752, 'SETROHADI', 1425, 76, 6),
		(17753, 'SUMARI', 1425, 76, 6),
		(17754, 'SUMENGKO', 1425, 76, 6),
		(17755, 'TAMBAKREJO', 1425, 76, 6),
		(17756, 'TEBALOAN', 1425, 76, 6),
		(17757, 'TIREM', 1425, 76, 6),
		(17758, 'TUMAPEL', 1425, 76, 6),
		(17759, 'WADAK KIDUL', 1425, 76, 6),
		(17760, 'WADAK LOR', 1425, 76, 6),
		(17761, 'BABAKBAWO', 1426, 76, 6),
		(17762, 'BABAKSARI', 1426, 76, 6),
		(17763, 'BANGERAN', 1426, 76, 6),
		(17764, 'BANYUBIRU', 1426, 76, 6),
		(17765, 'BANYUDONO', 1426, 76, 6),
		(17766, 'BARON', 1426, 76, 6),
		(17767, 'BULANGAN', 1426, 76, 6),
		(17768, 'DUKUH KEMBAR', 1426, 76, 6),
		(17769, 'DUKUN', 1426, 76, 6),
		(17770, 'DUKUNANYAR', 1426, 76, 6),
		(17771, 'GEDONGKEDOAN', 1426, 76, 6),
		(17772, 'IMAAN', 1426, 76, 6),
		(17773, 'JREBENG', 1426, 76, 6),
		(17774, 'KALIBENING', 1426, 76, 6),
		(17775, 'KALIREJO', 1426, 76, 6),
		(17776, 'KARANGCANGKRING', 1426, 76, 6),
		(17777, 'KENINGAR', 1426, 76, 6),
		(17778, 'KETUNGGENG', 1426, 76, 6),
		(17779, 'KRINJING', 1426, 76, 6),
		(17780, 'LOWAYU', 1426, 76, 6),
		(17781, 'MADUMULYOREJO', 1426, 76, 6),
		(17782, 'MANGUNSOKO', 1426, 76, 6),
		(17783, 'MENTARAS', 1426, 76, 6),
		(17784, 'MOJOPETUNG', 1426, 76, 6),
		(17785, 'NGADIPURO', 1426, 76, 6),
		(17786, 'NGARGOMULYO', 1426, 76, 6),
		(17787, 'PADANG BANDUNG', 1426, 76, 6),
		(17788, 'PATEN', 1426, 76, 6),
		(17789, 'PETIYIN TUNGGAL', 1426, 76, 6),
		(17790, 'SAMBOGUNUNG', 1426, 76, 6),
		(17791, 'SAWO', 1426, 76, 6),
		(17792, 'SEKARGADUNG', 1426, 76, 6),
		(17793, 'SEMBUNG ANYAR', 1426, 76, 6),
		(17794, 'SEMBUNGAN KIDUL', 1426, 76, 6),
		(17795, 'SENGI', 1426, 76, 6),
		(17796, 'SEWUKAN', 1426, 76, 6),
		(17797, 'SUMBER', 1426, 76, 6),
		(17798, 'TEBUWUNG', 1426, 76, 6),
		(17799, 'TIREMENGGAL', 1426, 76, 6),
		(17800, 'WATES', 1426, 76, 6),
		(17801, 'WONOKERTO', 1426, 76, 6),
		(17802, 'BEDILAN', 1427, 76, 6),
		(17803, 'GAPUROSUKOLILO', 1427, 76, 6),
		(17804, 'KARANGPOH', 1427, 76, 6),
		(17805, 'KARANGTURI', 1427, 76, 6),
		(17806, 'KEBUNGSON', 1427, 76, 6),
		(17807, 'KEMUTERAN', 1427, 76, 6),
		(17808, 'KRAMATINGGIL', 1427, 76, 6),
		(17809, 'KROMAN', 1427, 76, 6),
		(17810, 'LUMPUR', 1427, 76, 6),
		(17811, 'NGIPIK', 1427, 76, 6),
		(17812, 'PEKAUMAN', 1427, 76, 6),
		(17813, 'PEKELINGAN', 1427, 76, 6),
		(17814, 'PULOPANCIKAN', 1427, 76, 6),
		(17815, 'SIDOKUMPUL', 1427, 76, 6),
		(17816, 'SIDORUKUN', 1427, 76, 6),
		(17817, 'SUKODONO', 1427, 76, 6),
		(17818, 'SUKORAME', 1427, 76, 6),
		(17819, 'TLOGOBENDUNG', 1427, 76, 6),
		(17820, 'TLOGOPATUT', 1427, 76, 6),
		(17821, 'TLOGOPOJOK', 1427, 76, 6),
		(17822, 'TRATE', 1427, 76, 6),
		(17823, 'DAHANREJO', 1428, 76, 6),
		(17824, 'GENDING', 1428, 76, 6),
		(17825, 'GIRI', 1428, 76, 6),
		(17826, 'GULOMANTUNG', 1428, 76, 6),
		(17827, 'INDRO', 1428, 76, 6),
		(17828, 'KARANGKERING', 1428, 76, 6),
		(17829, 'KAWISANYAR', 1428, 76, 6),
		(17830, 'KEBOMAS', 1428, 76, 6),
		(17831, 'KEDANYANG', 1428, 76, 6),
		(17832, 'KEMBANGAN', 1428, 76, 6),
		(17833, 'KLANGONAN', 1428, 76, 6),
		(17834, 'NGARGOSARI', 1428, 76, 6),
		(17835, 'PRAMBANGAN', 1428, 76, 6),
		(17836, 'RANDUAGUNG', 1428, 76, 6),
		(17837, 'SEGOROMADU', 1428, 76, 6),
		(17838, 'SEKARKURUNG', 1428, 76, 6),
		(17839, 'SIDOMORO', 1428, 76, 6),
		(17840, 'SIDOMUKTI', 1428, 76, 6),
		(17841, 'SINGOSARI', 1428, 76, 6),
		(17842, 'SUKOREJO', 1428, 76, 6),
		(17843, 'TENGGULUNAN', 1428, 76, 6),
		(17844, 'BANYUURIP', 1429, 76, 6),
		(17845, 'BELAHANREJO', 1429, 76, 6),
		(17846, 'CERMEN', 1429, 76, 6),
		(17847, 'GLINDAH', 1429, 76, 6),
		(17848, 'KATIMOHO', 1429, 76, 6),
		(17849, 'KEDAMEAN', 1429, 76, 6),
		(17850, 'LAMPAH', 1429, 76, 6),
		(17851, 'MENUNGGAL', 1429, 76, 6),
		(17852, 'MOJOWUKU', 1429, 76, 6),
		(17853, 'NGEPUNG', 1429, 76, 6),
		(17854, 'SIDORAHARJO', 1429, 76, 6),
		(17855, 'SLEMPIT', 1429, 76, 6),
		(17856, 'TANJUNG', 1429, 76, 6),
		(17857, 'TULUNG', 1429, 76, 6),
		(17858, 'TURIREJO', 1429, 76, 6),
		(17859, 'BANJARSARI', 1430, 76, 6),
		(17860, 'BANYUWANGI', 1430, 76, 6),
		(17861, 'BETOYOGUCI', 1430, 76, 6),
		(17862, 'BETOYOKAUMAN', 1430, 76, 6),
		(17863, 'GUMENO', 1430, 76, 6),
		(17864, 'KARANGREJO', 1430, 76, 6),
		(17865, 'LERAN', 1430, 76, 6),
		(17866, 'MANYAR SIDOMUKTI', 1430, 76, 6),
		(17867, 'MANYAR SIDORUKUN', 1430, 76, 6),
		(17868, 'MANYAREJO', 1430, 76, 6),
		(17869, 'MOROBAKUNG', 1430, 76, 6),
		(17870, 'NGAMPEL', 1430, 76, 6),
		(17871, 'PEGANDEN', 1430, 76, 6),
		(17872, 'PEJANGGANAN', 1430, 76, 6),
		(17873, 'PONGANGAN', 1430, 76, 6),
		(17874, 'ROOMO', 1430, 76, 6),
		(17875, 'SEMBAYAT', 1430, 76, 6),
		(17876, 'SUCI', 1430, 76, 6),
		(17877, 'SUKOMULYO', 1430, 76, 6),
		(17878, 'SUMBEREJO', 1430, 76, 6),
		(17879, 'TANGGULREJO', 1430, 76, 6),
		(17880, 'TEBALO', 1430, 76, 6),
		(17881, 'YOSOWILANGUN', 1430, 76, 6),
		(17882, 'BETON', 1431, 76, 6),
		(17883, 'BOBOH', 1431, 76, 6),
		(17884, 'BOTENG', 1431, 76, 6),
		(17885, 'BRINGKANG', 1431, 76, 6),
		(17886, 'DOMAS', 1431, 76, 6),
		(17887, 'DRANCANG', 1431, 76, 6),
		(17888, 'GADINGWATU', 1431, 76, 6),
		(17889, 'GEMPOLKURUNG', 1431, 76, 6),
		(17890, 'HENDROSARI', 1431, 76, 6),
		(17891, 'HULAAN', 1431, 76, 6),
		(17892, 'KEPATIHAN', 1431, 76, 6),
		(17893, 'LABAN', 1431, 76, 6),
		(17894, 'MENGANTI', 1431, 76, 6),
		(17895, 'MOJOTENGAH', 1431, 76, 6),
		(17896, 'PELEMWATU', 1431, 76, 6),
		(17897, 'PENGALANGAN', 1431, 76, 6),
		(17898, 'PRANTI', 1431, 76, 6),
		(17899, 'PUTAT LOR', 1431, 76, 6),
		(17900, 'RANDUPADANGAN', 1431, 76, 6),
		(17901, 'SETRO', 1431, 76, 6),
		(17902, 'SIDOJANGKUNG', 1431, 76, 6),
		(17903, 'SIDOWUNGU', 1431, 76, 6),
		(17904, 'BANYUTENGAH', 1432, 76, 6),
		(17905, 'CAMPUREJO', 1432, 76, 6),
		(17906, 'DALEGAN', 1432, 76, 6),
		(17907, 'DOUDO', 1432, 76, 6),
		(17908, 'KETANEN', 1432, 76, 6),
		(17909, 'PANTENAN', 1432, 76, 6),
		(17910, 'PETUNG', 1432, 76, 6),
		(17911, 'PRUPUH', 1432, 76, 6),
		(17912, 'SERAH', 1432, 76, 6),
		(17913, 'SIWALAN', 1432, 76, 6),
		(17914, 'SUKODONO', 1432, 76, 6),
		(17915, 'SUMURBER', 1432, 76, 6),
		(17916, 'SUROWITI', 1432, 76, 6),
		(17917, 'WOTAN', 1432, 76, 6),
		(17918, 'BALIKTERUS', 1433, 76, 6),
		(17919, 'BULULANJANG', 1433, 76, 6),
		(17920, 'DAUN', 1433, 76, 6),
		(17921, 'DEKATAGUNG', 1433, 76, 6),
		(17922, 'GUNUNGTEGUH', 1433, 76, 6),
		(17923, 'KEBON TELUK DALAM', 1433, 76, 6),
		(17924, 'KOTA KUSUMA', 1433, 76, 6),
		(17925, 'KUMALASA', 1433, 76, 6),
		(17926, 'LEBAK', 1433, 76, 6),
		(17927, 'PATARSELAMAT', 1433, 76, 6),
		(17928, 'PUDAKIT BARAT', 1433, 76, 6),
		(17929, 'PUDAKIT TIMUR', 1433, 76, 6),
		(17930, 'SAWAHMULYA', 1433, 76, 6),
		(17931, 'SIDOGEDUNG BATU', 1433, 76, 6),
		(17932, 'SUNGAI RUJING', 1433, 76, 6),
		(17933, 'SUNGAI TELUK', 1433, 76, 6),
		(17934, 'SUWARI', 1433, 76, 6),
		(17935, 'ASEMPAPAK', 1434, 76, 6),
		(17936, 'BUNDERAN', 1434, 76, 6),
		(17937, 'GEDANGAN', 1434, 76, 6),
		(17938, 'GOLOKAN', 1434, 76, 6),
		(17939, 'KAUMAN', 1434, 76, 6),
		(17940, 'KERTOSONO', 1434, 76, 6),
		(17941, 'LASEM', 1434, 76, 6),
		(17942, 'MOJOASEM', 1434, 76, 6),
		(17943, 'MRIYUNAN', 1434, 76, 6),
		(17944, 'NGAWEN', 1434, 76, 6),
		(17945, 'PENGULU', 1434, 76, 6),
		(17946, 'PURWODADI', 1434, 76, 6),
		(17947, 'RACIKULON', 1434, 76, 6),
		(17948, 'RACITENGAH', 1434, 76, 6),
		(17949, 'RANDUBOTO', 1434, 76, 6),
		(17950, 'SAMBIPONDOK', 1434, 76, 6),
		(17951, 'SEDAGARAN', 1434, 76, 6),
		(17952, 'SIDOMULYO', 1434, 76, 6),
		(17953, 'SROWO', 1434, 76, 6),
		(17954, 'SUKOREJO', 1434, 76, 6),
		(17955, 'WADENG', 1434, 76, 6),
		(17956, 'BUNIAYU', 1435, 76, 6),
		(17957, 'DIPONGGO', 1435, 76, 6),
		(17958, 'GEBANGSARI', 1435, 76, 6),
		(17959, 'GELAM', 1435, 76, 6),
		(17960, 'GREJEG', 1435, 76, 6),
		(17961, 'GUMELAR KIDUL', 1435, 76, 6),
		(17962, 'GUMELAR LOR', 1435, 76, 6),
		(17963, 'KAMULYAN', 1435, 76, 6),
		(17964, 'KARANGPETIR', 1435, 76, 6),
		(17965, 'KARANGPUCUNG', 1435, 76, 6),
		(17966, 'KELOMPANG GUBUG', 1435, 76, 6),
		(17967, 'KEPUH LEGUNDI', 1435, 76, 6),
		(17968, 'KEPUH TELUK', 1435, 76, 6),
		(17969, 'PEKALONGAN', 1435, 76, 6),
		(17970, 'PEROMAAN (PAROMAAN)', 1435, 76, 6),
		(17971, 'PESANTREN', 1435, 76, 6),
		(17972, 'PLANGKAPAN', 1435, 76, 6),
		(17973, 'PREMBUN', 1435, 76, 6),
		(17974, 'PURWODADI', 1435, 76, 6),
		(17975, 'SUKALELA', 1435, 76, 6),
		(17976, 'SUKAONENG', 1435, 76, 6),
		(17977, 'TAMBAK', 1435, 76, 6),
		(17978, 'TANJUNGORI', 1435, 76, 6),
		(17979, 'TELUK JATIDAWANG', 1435, 76, 6),
		(17980, 'WATUAGUNG', 1435, 76, 6),
		(17981, 'BANYUURIP', 1436, 76, 6),
		(17982, 'BOLO', 1436, 76, 6),
		(17983, 'CANGAAN', 1436, 76, 6),
		(17984, 'GLATIK', 1436, 76, 6),
		(17985, 'GOSARI', 1436, 76, 6),
		(17986, 'KARANGREJO', 1436, 76, 6),
		(17987, 'KEBONAGUNG', 1436, 76, 6),
		(17988, 'KETAPANG LOR', 1436, 76, 6),
		(17989, 'NGEMBOH', 1436, 76, 6),
		(17990, 'PANGKAH KULON', 1436, 76, 6),
		(17991, 'PANGKAH WETAN', 1436, 76, 6),
		(17992, 'SEKAPUK', 1436, 76, 6),
		(17993, 'TANJANGAWAN', 1436, 76, 6),
		(17994, 'KEDUNGANYAR', 1437, 76, 6),
		(17995, 'KEPUHKLAGEN', 1437, 76, 6),
		(17996, 'KESAMBEN KULON', 1437, 76, 6),
		(17997, 'LEBANISUKO', 1437, 76, 6),
		(17998, 'LEBANIWARAS', 1437, 76, 6),
		(17999, 'MONDOLUKU', 1437, 76, 6),
		(18000, 'PASINAN LEMAHPUTIH', 1437, 76, 6)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 17")
}
