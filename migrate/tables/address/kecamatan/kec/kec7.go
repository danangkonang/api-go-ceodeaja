package kec

import "github.com/danangkonang/ceodeaja-go/config"

func Kec7() {
	db := config.Connect()
	db.Exec(`INSERT INTO kecamatan (id,kecamatan,kota_id,provinsi_id) VALUES
		(6001, 'MIOMAFFO TIMUR', 435, 32),
		(6002, 'MUSI', 435, 32),
		(6003, 'MUTIS', 435, 32),
		(6004, 'NAIBENU', 435, 32),
		(6005, 'NOEMUTI', 435, 32),
		(6006, 'NOEMUTI TIMUR', 435, 32),
		(6007, 'BOMBARAI (BOMBERAY)', 436, 33),
		(6008, 'FAKFAK', 436, 33),
		(6009, 'FAKFAK BARAT', 436, 33),
		(6010, 'FAKFAK TENGAH', 436, 33),
		(6011, 'FAKFAK TIMUR', 436, 33),
		(6012, 'KARAS', 436, 33),
		(6013, 'KOKAS', 436, 33),
		(6014, 'KRAMONGMONGGA (KRAMAMONGGA)', 436, 33),
		(6015, 'TELUK PATIPI', 436, 33),
		(6016, 'BURUWAY', 437, 33),
		(6017, 'KAIMANA', 437, 33),
		(6018, 'KAMBRAW (KAMBERAU)', 437, 33),
		(6019, 'TELUK ARGUNI ATAS', 437, 33),
		(6020, 'TELUK ARGUNI BAWAH (YERUSI)', 437, 33),
		(6021, 'TELUK ETNA', 437, 33),
		(6022, 'YAMOR', 437, 33),
		(6023, 'MANOKWARI BARAT', 438, 33),
		(6024, 'MANOKWARI SELATAN', 438, 33),
		(6025, 'MANOKWARI TIMUR', 438, 33),
		(6026, 'MANOKWARI UTARA', 438, 33),
		(6027, 'MASNI', 438, 33),
		(6028, 'PRAFI', 438, 33),
		(6029, 'SIDEY', 438, 33),
		(6030, 'TANAH RUBUH', 438, 33),
		(6031, 'WARMARE', 438, 33),
		(6032, 'DATARAN ISIM', 439, 33),
		(6033, 'MOMI WAREN', 439, 33),
		(6034, 'NENEY (NENEI)', 439, 33),
		(6035, 'ORANSBARI', 439, 33),
		(6036, 'RANSIKI', 439, 33),
		(6037, 'TAHOTA (TOHOTA)', 439, 33),
		(6038, 'AIFAT', 440, 33),
		(6039, 'AIFAT SELATAN', 440, 33),
		(6040, 'AIFAT TIMUR', 440, 33),
		(6041, 'AIFAT TIMUR JAUH', 440, 33),
		(6042, 'AIFAT TIMUR SELATAN', 440, 33),
		(6043, 'AIFAT TIMUR TENGAH', 440, 33),
		(6044, 'AIFAT UTARA', 440, 33),
		(6045, 'AITINYO', 440, 33),
		(6046, 'AITINYO BARAT', 440, 33),
		(6047, 'AITINYO RAYA', 440, 33),
		(6048, 'AITINYO TENGAH', 440, 33),
		(6049, 'AITINYO UTARA', 440, 33),
		(6050, 'AYAMARU', 440, 33),
		(6051, 'AYAMARU BARAT', 440, 33),
		(6052, 'AYAMARU JAYA', 440, 33),
		(6053, 'AYAMARU SELATAN', 440, 33),
		(6054, 'AYAMARU SELATAN JAYA', 440, 33),
		(6055, 'AYAMARU TENGAH', 440, 33),
		(6056, 'AYAMARU TIMUR', 440, 33),
		(6057, 'AYAMARU TIMUR SELATAN', 440, 33),
		(6058, 'AYAMARU UTARA', 440, 33),
		(6059, 'AYAMARU UTARA TIMUR', 440, 33),
		(6060, 'MARE', 440, 33),
		(6061, 'MARE SELATAN', 440, 33),
		(6062, 'ANGGI', 441, 33),
		(6063, 'ANGGI GIDA', 441, 33),
		(6064, 'CATUBOUW', 441, 33),
		(6065, 'DIDOHU', 441, 33),
		(6066, 'HINGK', 441, 33),
		(6067, 'MEMBEY', 441, 33),
		(6068, 'MENYAMBOUW (MINYAMBOUW)', 441, 33),
		(6069, 'SURUREY', 441, 33),
		(6070, 'TAIGE', 441, 33),
		(6071, 'TESTEGA', 441, 33),
		(6072, 'AYAU', 442, 33),
		(6073, 'BATANTA SELATAN', 442, 33),
		(6074, 'BATANTA UTARA', 442, 33),
		(6075, 'KEPULAUAN AYAU', 442, 33),
		(6076, 'KEPULAUAN SEMBILAN', 442, 33),
		(6077, 'KOFIAU', 442, 33),
		(6078, 'KOTA WAISAI', 442, 33),
		(6079, 'MEOS MANSAR', 442, 33),
		(6080, 'MISOOL (MISOOL UTARA)', 442, 33),
		(6081, 'MISOOL BARAT', 442, 33),
		(6082, 'MISOOL SELATAN', 442, 33),
		(6083, 'MISOOL TIMUR', 442, 33),
		(6084, 'SALAWATI BARAT', 442, 33),
		(6085, 'SALAWATI TENGAH', 442, 33),
		(6086, 'SALAWATI UTARA (SAMATE)', 442, 33),
		(6087, 'SUPNIN', 442, 33),
		(6088, 'TELUK MAYALIBIT', 442, 33),
		(6089, 'TIPLOL MAYALIBIT', 442, 33),
		(6090, 'WAIGEO BARAT', 442, 33),
		(6091, 'WAIGEO BARAT KEPULAUAN', 442, 33),
		(6092, 'WAIGEO SELATAN', 442, 33),
		(6093, 'WAIGEO TIMUR', 442, 33),
		(6094, 'WAIGEO UTARA', 442, 33),
		(6095, 'WARWABOMI', 442, 33),
		(6096, 'AIMAS', 443, 33),
		(6097, 'BERAUR', 443, 33),
		(6098, 'KLABOT', 443, 33),
		(6099, 'KLAMONO', 443, 33),
		(6100, 'KLASO', 443, 33),
		(6101, 'KLAWAK', 443, 33),
		(6102, 'KLAYILI', 443, 33),
		(6103, 'MAKBON', 443, 33),
		(6104, 'MARIAT', 443, 33),
		(6105, 'MAUDUS', 443, 33),
		(6106, 'MAYAMUK', 443, 33),
		(6107, 'MOISEGEN', 443, 33),
		(6108, 'SALAWATI', 443, 33),
		(6109, 'SALAWATI SELATAN', 443, 33),
		(6110, 'SAYOSA', 443, 33),
		(6111, 'SEGET', 443, 33),
		(6112, 'SEGUN', 443, 33),
		(6113, 'SORONG', 443, 33),
		(6114, 'SORONG BARAT', 443, 33),
		(6115, 'SORONG KEPULAUAN', 443, 33),
		(6116, 'SORONG MANOI', 443, 33),
		(6117, 'SORONG TIMUR', 443, 33),
		(6118, 'SORONG UTARA', 443, 33),
		(6119, 'FOKOUR', 444, 33),
		(6120, 'INANWATAN (INAWATAN)', 444, 33),
		(6121, 'KAIS (MATEMANI KAIS)', 444, 33),
		(6122, 'KOKODA', 444, 33),
		(6123, 'KOKODA UTARA', 444, 33),
		(6124, 'KONDA', 444, 33),
		(6125, 'MATEMANI', 444, 33),
		(6126, 'MOSWAREN', 444, 33),
		(6127, 'SAIFI', 444, 33),
		(6128, 'SAWIAT', 444, 33),
		(6129, 'SEREMUK', 444, 33),
		(6130, 'TEMINABUAN', 444, 33),
		(6131, 'WAYER', 444, 33),
		(6132, 'ABUN', 445, 33),
		(6133, 'AMBERBAKEN', 445, 33),
		(6134, 'FEF (PEEF)', 445, 33),
		(6135, 'KEBAR', 445, 33),
		(6136, 'KWOOR', 445, 33),
		(6137, 'MIYAH (MEYAH)', 445, 33),
		(6138, 'MORAID', 445, 33),
		(6139, 'MUBRANI', 445, 33),
		(6140, 'SAUSAPOR', 445, 33),
		(6141, 'SENOPI', 445, 33),
		(6142, 'SYUJAK', 445, 33),
		(6143, 'YEMBUN', 445, 33),
		(6144, 'ARANDAY', 446, 33),
		(6145, 'AROBA', 446, 33),
		(6146, 'BABO', 446, 33),
		(6147, 'BINTUNI', 446, 33),
		(6148, 'BISCOOP', 446, 33),
		(6149, 'DATARAN BEIMES', 446, 33),
		(6150, 'FAFURWAR (IRORUTU)', 446, 33),
		(6151, 'KAITARO', 446, 33),
		(6152, 'KAMUNDAN', 446, 33),
		(6153, 'KURI', 446, 33),
		(6154, 'MANIMERI', 446, 33),
		(6155, 'MASYETA', 446, 33),
		(6156, 'MAYADO', 446, 33),
		(6157, 'MERDEY', 446, 33),
		(6158, 'MOSKONA BARAT', 446, 33),
		(6159, 'MOSKONA SELATAN', 446, 33),
		(6160, 'MOSKONA TIMUR', 446, 33),
		(6161, 'MOSKONA UTARA', 446, 33),
		(6162, 'SUMURI (SIMURI)', 446, 33),
		(6163, 'TEMBUNI', 446, 33),
		(6164, 'TOMU', 446, 33),
		(6165, 'TUHIBA', 446, 33),
		(6166, 'WAMESA (IDOOR)', 446, 33),
		(6167, 'WERIAGAR', 446, 33),
		(6168, 'KURI WAMESA', 447, 33),
		(6169, 'NAIKERE (WASIOR BARAT)', 447, 33),
		(6170, 'NIKIWAR', 447, 33),
		(6171, 'RASIEI', 447, 33),
		(6172, 'ROON', 447, 33),
		(6173, 'ROSWAR', 447, 33),
		(6174, 'RUMBERPON', 447, 33),
		(6175, 'SOUG JAYA', 447, 33),
		(6176, 'TELUK DUAIRI (WASIOR UTARA)', 447, 33),
		(6177, 'WAMESA', 447, 33),
		(6178, 'WASIOR', 447, 33),
		(6179, 'WINDESI', 447, 33),
		(6180, 'WONDIBOY (WASIOR SELATAN)', 447, 33),
		(6181, 'AGATS', 448, 34),
		(6182, 'AKAT', 448, 34),
		(6183, 'ATSY / ATSJ', 448, 34),
		(6184, 'AYIP', 448, 34),
		(6185, 'BETCBAMU', 448, 34),
		(6186, 'DER KOUMUR', 448, 34),
		(6187, 'FAYIT', 448, 34),
		(6188, 'JETSY', 448, 34),
		(6189, 'JOERAT', 448, 34),
		(6190, 'KOLF BRAZA', 448, 34),
		(6191, 'KOPAY', 448, 34),
		(6192, 'PANTAI KASUARI', 448, 34),
		(6193, 'PULAU TIGA', 448, 34),
		(6194, 'SAFAN', 448, 34),
		(6195, 'SAWA ERMA', 448, 34),
		(6196, 'SIRETS', 448, 34),
		(6197, 'SUATOR', 448, 34),
		(6198, 'SURU-SURU', 448, 34),
		(6199, 'UNIR SIRAU', 448, 34),
		(6200, 'AIMANDO PADAIDO', 449, 34),
		(6201, 'ANDEY (ANDEI)', 449, 34),
		(6202, 'BIAK BARAT', 449, 34),
		(6203, 'BIAK KOTA', 449, 34),
		(6204, 'BIAK TIMUR', 449, 34),
		(6205, 'BIAK UTARA', 449, 34),
		(6206, 'BONDIFUAR', 449, 34),
		(6207, 'BRUYADORI', 449, 34),
		(6208, 'NUMFOR BARAT', 449, 34),
		(6209, 'NUMFOR TIMUR', 449, 34),
		(6210, 'ORIDEK', 449, 34),
		(6211, 'ORKERI', 449, 34),
		(6212, 'PADAIDO', 449, 34),
		(6213, 'POIRU', 449, 34),
		(6214, 'SAMOFA', 449, 34),
		(6215, 'SWANDIWE', 449, 34),
		(6216, 'WARSA', 449, 34),
		(6217, 'YAWOSI', 449, 34),
		(6218, 'YENDIDORI', 449, 34),
		(6219, 'AMBATKWI (AMBATKUI)', 450, 34),
		(6220, 'ARIMOP', 450, 34),
		(6221, 'BOMAKIA', 450, 34),
		(6222, 'FIRIWAGE', 450, 34),
		(6223, 'FOFI', 450, 34),
		(6224, 'INIYANDIT', 450, 34),
		(6225, 'JAIR', 450, 34),
		(6226, 'KAWAGIT', 450, 34),
		(6227, 'KI', 450, 34),
		(6228, 'KOMBAY', 450, 34),
		(6229, 'KOMBUT', 450, 34),
		(6230, 'KOUH', 450, 34),
		(6231, 'MANDOBO', 450, 34),
		(6232, 'MANGGELUM', 450, 34),
		(6233, 'MINDIPTANA', 450, 34),
		(6234, 'NINATI', 450, 34),
		(6235, 'SESNUK', 450, 34),
		(6236, 'SUBUR', 450, 34),
		(6237, 'WAROPKO', 450, 34),
		(6238, 'YANIRUMA', 450, 34),
		(6239, 'BOWOBADO', 451, 34),
		(6240, 'KAPIRAYA', 451, 34),
		(6241, 'TIGI', 451, 34),
		(6242, 'TIGI BARAT', 451, 34),
		(6243, 'TIGI TIMUR', 451, 34),
		(6244, 'DOGIYAI', 452, 34),
		(6245, 'KAMU', 452, 34),
		(6246, 'KAMU SELATAN', 452, 34),
		(6247, 'KAMU TIMUR', 452, 34),
		(6248, 'KAMU UTARA (IKRAR/IKRAT)', 452, 34),
		(6249, 'MAPIA', 452, 34),
		(6250, 'MAPIA BARAT', 452, 34),
		(6251, 'MAPIA TENGAH', 452, 34),
		(6252, 'PIYAIYE (SUKIKAI)', 452, 34),
		(6253, 'SUKIKAI SELATAN', 452, 34),
		(6254, 'AGISIGA', 453, 34),
		(6255, 'BIANDOGA', 453, 34),
		(6256, 'HITADIPA', 453, 34),
		(6257, 'HOMEO (HOMEYO)', 453, 34),
		(6258, 'SUGAPA', 453, 34),
		(6259, 'WANDAI', 453, 34),
		(6260, 'ABEPURA', 454, 34),
		(6261, 'AIRU', 454, 34),
		(6262, 'DEMTA', 454, 34),
		(6263, 'DEPAPRE', 454, 34),
		(6264, 'EBUNGFAU (EBUNGFA)', 454, 34),
		(6265, 'GRESI SELATAN', 454, 34),
		(6266, 'HERAM', 454, 34),
		(6267, 'JAYAPURA SELATAN', 454, 34),
		(6268, 'JAYAPURA UTARA', 454, 34),
		(6269, 'KAUREH', 454, 34),
		(6270, 'KEMTUK', 454, 34),
		(6271, 'KEMTUK GRESI', 454, 34),
		(6272, 'MUARA TAMI', 454, 34),
		(6273, 'NAMBLUONG', 454, 34),
		(6274, 'NIMBOKRANG', 454, 34),
		(6275, 'NIMBORAN', 454, 34),
		(6276, 'RAVENIRARA', 454, 34),
		(6277, 'SENTANI', 454, 34),
		(6278, 'SENTANI BARAT', 454, 34),
		(6279, 'SENTANI TIMUR', 454, 34),
		(6280, 'UNURUM GUAY', 454, 34),
		(6281, 'WAIBU', 454, 34),
		(6282, 'YAPSI', 454, 34),
		(6283, 'YOKARI', 454, 34),
		(6284, 'ASOLOGAIMA (ASALOGAIMA)', 455, 34),
		(6285, 'ASOLOKOBAL', 455, 34),
		(6286, 'ASOTIPO', 455, 34),
		(6287, 'BOLAKME', 455, 34),
		(6288, 'BPIRI', 455, 34),
		(6289, 'BUGI', 455, 34),
		(6290, 'HUBIKIAK', 455, 34),
		(6291, 'HUBIKOSI (HOBIKOSI)', 455, 34),
		(6292, 'IBELE', 455, 34),
		(6293, 'ITLAY HISAGE', 455, 34),
		(6294, 'KORAGI', 455, 34),
		(6295, 'KURULU', 455, 34),
		(6296, 'LIBAREK', 455, 34),
		(6297, 'MAIMA', 455, 34),
		(6298, 'MOLAGALOME', 455, 34),
		(6299, 'MULIAMA', 455, 34),
		(6300, 'MUSATFAK', 455, 34),
		(6301, 'NAPUA', 455, 34),
		(6302, 'PELEBAGA', 455, 34),
		(6303, 'PIRAMID', 455, 34),
		(6304, 'PISUGI', 455, 34),
		(6305, 'POPUGOBA', 455, 34),
		(6306, 'SIEPKOSI', 455, 34),
		(6307, 'SILO KARNO DOGA', 455, 34),
		(6308, 'TAELAREK', 455, 34),
		(6309, 'TRIKORA', 455, 34),
		(6310, 'USILIMO', 455, 34),
		(6311, 'WADANGKU', 455, 34),
		(6312, 'WALAIK', 455, 34),
		(6313, 'WALELAGAMA', 455, 34),
		(6314, 'WAME', 455, 34),
		(6315, 'WAMENA', 455, 34),
		(6316, 'WELESI', 455, 34),
		(6317, 'WESAPUT', 455, 34),
		(6318, 'WITA WAYA', 455, 34),
		(6319, 'WOLLO (WOLO)', 455, 34),
		(6320, 'WOUMA', 455, 34),
		(6321, 'YALENGGA', 455, 34),
		(6322, 'ARSO', 456, 34),
		(6323, 'ARSO TIMUR', 456, 34),
		(6324, 'SENGGI', 456, 34),
		(6325, 'SKAMTO (SKANTO)', 456, 34),
		(6326, 'TOWE', 456, 34),
		(6327, 'WARIS', 456, 34),
		(6328, 'WEB', 456, 34),
		(6329, 'ANGKAISERA', 457, 34),
		(6330, 'KEPULAUAN AMBAI', 457, 34),
		(6331, 'KOSIWO', 457, 34),
		(6332, 'POOM', 457, 34),
		(6333, 'PULAU KURUDU', 457, 34),
		(6334, 'PULAU YERUI', 457, 34),
		(6335, 'RAIMBAWI', 457, 34),
		(6336, 'TELUK AMPIMOI', 457, 34),
		(6337, 'WONAWA', 457, 34),
		(6338, 'YAPEN BARAT', 457, 34),
		(6339, 'YAPEN SELATAN', 457, 34),
		(6340, 'YAPEN TIMUR', 457, 34),
		(6341, 'YAPEN UTARA', 457, 34),
		(6342, 'BALINGGA', 458, 34),
		(6343, 'DIMBA', 458, 34),
		(6344, 'GAMELIA', 458, 34),
		(6345, 'KUYAWAGE', 458, 34),
		(6346, 'MAKKI (MAKI)', 458, 34),
		(6347, 'MALAGAINERI (MALAGINERI)', 458, 34),
		(6348, 'PIRIME', 458, 34),
		(6349, 'POGA', 458, 34),
		(6350, 'TIOM', 458, 34),
		(6351, 'TIOMNERI', 458, 34),
		(6352, 'BENUKI', 459, 34),
		(6353, 'MAMBERAMO HILIR/ILIR', 459, 34),
		(6354, 'MAMBERAMO HULU/ULU', 459, 34),
		(6355, 'MAMBERAMO TENGAH', 459, 34),
		(6356, 'MAMBERAMO TENGAH TIMUR', 459, 34),
		(6357, 'ROFAER (RUFAER)', 459, 34),
		(6358, 'SAWAI', 459, 34),
		(6359, 'WAROPEN ATAS', 459, 34),
		(6360, 'ERAGAYAM', 460, 34),
		(6361, 'ILUGWA', 460, 34),
		(6362, 'KELILA', 460, 34),
		(6363, 'KOBAKMA', 460, 34),
		(6364, 'MEGABILIS (MEGAMBILIS)', 460, 34),
		(6365, 'ASSUE', 461, 34),
		(6366, 'BAMGI', 461, 34),
		(6367, 'CITAKMITAK', 461, 34),
		(6368, 'EDERA', 461, 34),
		(6369, 'HAJU', 461, 34),
		(6370, 'KAIBAR', 461, 34),
		(6371, 'MINYAMUR', 461, 34),
		(6372, 'NAMBIOMAN BAPAI (MAMBIOMAN)', 461, 34),
		(6373, 'OBAA', 461, 34),
		(6374, 'PASSUE', 461, 34),
		(6375, 'PASSUE BAWAH', 461, 34),
		(6376, 'SYAHCAME', 461, 34),
		(6377, 'TI ZAIN', 461, 34),
		(6378, 'VENAHA', 461, 34),
		(6379, 'YAKOMI', 461, 34),
		(6380, 'ANIMHA', 462, 34),
		(6381, 'ELIGOBEL', 462, 34),
		(6382, 'ILYAWAB', 462, 34),
		(6383, 'JAGEBOB', 462, 34),
		(6384, 'KAPTEL', 462, 34),
		(6385, 'KIMAAM', 462, 34),
		(6386, 'KURIK', 462, 34),
		(6387, 'MALIND', 462, 34),
		(6388, 'MERAUKE', 462, 34),
		(6389, 'MUTING', 462, 34),
		(6390, 'NAUKENJERAI', 462, 34),
		(6391, 'NGGUTI (NGGUNTI)', 462, 34),
		(6392, 'OKABA', 462, 34),
		(6393, 'SEMANGGA', 462, 34),
		(6394, 'SOTA', 462, 34),
		(6395, 'TABONJI', 462, 34),
		(6396, 'TANAH MIRING', 462, 34),
		(6397, 'TUBANG', 462, 34),
		(6398, 'ULILIN', 462, 34),
		(6399, 'WAAN', 462, 34),
		(6400, 'AGIMUGA', 463, 34),
		(6401, 'JILA', 463, 34),
		(6402, 'JITA', 463, 34),
		(6403, 'KUALA KENCANA', 463, 34),
		(6404, 'MIMIKA BARAT (MIBAR)', 463, 34),
		(6405, 'MIMIKA BARAT JAUH', 463, 34),
		(6406, 'MIMIKA BARAT TENGAH', 463, 34),
		(6407, 'MIMIKA BARU', 463, 34),
		(6408, 'MIMIKA TIMUR', 463, 34),
		(6409, 'MIMIKA TIMUR JAUH', 463, 34),
		(6410, 'MIMIKA TIMUR TENGAH', 463, 34),
		(6411, 'TEMBAGAPURA', 463, 34),
		(6412, 'DIPA', 464, 34),
		(6413, 'MAKIMI', 464, 34),
		(6414, 'MENOU', 464, 34),
		(6415, 'MOORA', 464, 34),
		(6416, 'NABIRE', 464, 34),
		(6417, 'NABIRE BARAT', 464, 34),
		(6418, 'NAPAN', 464, 34),
		(6419, 'SIRIWO', 464, 34),
		(6420, 'TELUK KIMI', 464, 34),
		(6421, 'TELUK UMAR', 464, 34),
		(6422, 'UWAPA', 464, 34),
		(6423, 'WANGGAR', 464, 34),
		(6424, 'WAPOGA', 464, 34),
		(6425, 'YARO (YARO KABISAY)', 464, 34),
		(6426, 'YAUR', 464, 34),
		(6427, 'ALAMA', 465, 34),
		(6428, 'DAL', 465, 34),
		(6429, 'EMBETPEN', 465, 34),
		(6430, 'GEAREK', 465, 34),
		(6431, 'GESELMA (GESELEMA)', 465, 34),
		(6432, 'INIKGAL', 465, 34),
		(6433, 'INIYE', 465, 34),
		(6434, 'KEGAYEM', 465, 34),
		(6435, 'KENYAM', 465, 34),
		(6436, 'KILMID', 465, 34),
		(6437, 'KORA', 465, 34),
		(6438, 'KOROPTAK', 465, 34),
		(6439, 'KREPKURI', 465, 34),
		(6440, 'MAM', 465, 34),
		(6441, 'MAPENDUMA', 465, 34),
		(6442, 'MBUA (MBUGA)', 465, 34),
		(6443, 'MBUA TENGAH', 465, 34),
		(6444, 'MBULMU YALMA', 465, 34),
		(6445, 'MEBAROK', 465, 34),
		(6446, 'MOBA', 465, 34),
		(6447, 'NENGGEAGIN', 465, 34),
		(6448, 'NIRKURI', 465, 34),
		(6449, 'PARO', 465, 34),
		(6450, 'PASIR PUTIH', 465, 34),
		(6451, 'PIJA', 465, 34),
		(6452, 'WOSAK', 465, 34),
		(6453, 'WUSI', 465, 34),
		(6454, 'WUTPAGA', 465, 34),
		(6455, 'YAL', 465, 34),
		(6456, 'YENGGELO', 465, 34),
		(6457, 'YIGI', 465, 34),
		(6458, 'ARADIDE', 466, 34),
		(6459, 'BIBIDA', 466, 34),
		(6460, 'BOGOBAIDA', 466, 34),
		(6461, 'DUMADAMA', 466, 34),
		(6462, 'EKADIDE', 466, 34),
		(6463, 'KEBO', 466, 34),
		(6464, 'PANIAI BARAT', 466, 34),
		(6465, 'PANIAI TIMUR', 466, 34),
		(6466, 'YATAMO', 466, 34),
		(6467, 'ABOY', 467, 34),
		(6468, 'ALEMSOM', 467, 34),
		(6469, 'AWINBON', 467, 34),
		(6470, 'BATANI', 467, 34),
		(6471, 'BATOM', 467, 34),
		(6472, 'BIME', 467, 34),
		(6473, 'BORME', 467, 34),
		(6474, 'EIPUMEK', 467, 34),
		(6475, 'IWUR (OKIWUR)', 467, 34),
		(6476, 'JETFA', 467, 34),
		(6477, 'KALOMDOL', 467, 34),
		(6478, 'KAWOR', 467, 34),
		(6479, 'KIWIROK', 467, 34),
		(6480, 'KIWIROK TIMUR', 467, 34),
		(6481, 'MOFINOP', 467, 34),
		(6482, 'MURKIM', 467, 34),
		(6483, 'NONGME', 467, 34),
		(6484, 'OK AOM', 467, 34),
		(6485, 'OKBAB', 467, 34),
		(6486, 'OKBAPE', 467, 34),
		(6487, 'OKBEMTAU', 467, 34),
		(6488, 'OKBIBAB', 467, 34),
		(6489, 'OKHIKA', 467, 34),
		(6490, 'OKLIP', 467, 34),
		(6491, 'OKSAMOL', 467, 34),
		(6492, 'OKSEBANG', 467, 34),
		(6493, 'OKSIBIL', 467, 34),
		(6494, 'OKSOP', 467, 34),
		(6495, 'PAMEK', 467, 34),
		(6496, 'PEPERA', 467, 34),
		(6497, 'SERAMBAKON', 467, 34),
		(6498, 'TARUP', 467, 34),
		(6499, 'TEIRAPLU', 467, 34),
		(6500, 'WEIME', 467, 34),
		(6501, 'AGADUGUME', 468, 34),
		(6502, 'BEOGA', 468, 34),
		(6503, 'DOUFU', 468, 34),
		(6504, 'GOME', 468, 34),
		(6505, 'ILAGA', 468, 34),
		(6506, 'POGOMA', 468, 34),
		(6507, 'SINAK', 468, 34),
		(6508, 'WANGBE', 468, 34),
		(6509, 'FAWI', 469, 34),
		(6510, 'ILU', 469, 34),
		(6511, 'JIGONIKME', 469, 34),
		(6512, 'MEWOLUK (MEWULOK)', 469, 34),
		(6513, 'MULIA', 469, 34),
		(6514, 'TINGGINAMBUT', 469, 34),
		(6515, 'TORERE', 469, 34),
		(6516, 'YAMO', 469, 34),
		(6517, 'APAWER HULU', 470, 34),
		(6518, 'BONGGO', 470, 34),
		(6519, 'BONGGO TIMUR', 470, 34),
		(6520, 'PANTAI BARAT', 470, 34),
		(6521, 'PANTAI TIMUR', 470, 34),
		(6522, 'PANTAI TIMUR BARAT', 470, 34),
		(6523, 'SARMI', 470, 34),
		(6524, 'SARMI SELATAN', 470, 34),
		(6525, 'SARMI TIMUR', 470, 34),
		(6526, 'TOR ATAS', 470, 34),
		(6527, 'KEPULAUAN ARURI', 471, 34),
		(6528, 'SUPIORI BARAT', 471, 34),
		(6529, 'SUPIORI SELATAN', 471, 34),
		(6530, 'SUPIORI TIMUR', 471, 34),
		(6531, 'SUPIORI UTARA', 471, 34),
		(6532, 'AIRGARAM', 472, 34),
		(6533, 'ANAWI', 472, 34),
		(6534, 'AWEKU', 472, 34),
		(6535, 'BEWANI', 472, 34),
		(6536, 'BIUK', 472, 34),
		(6537, 'BOGONUK', 472, 34),
		(6538, 'BOKONDINI', 472, 34),
		(6539, 'BOKONERI', 472, 34),
		(6540, 'DANIME', 472, 34),
		(6541, 'DOW', 472, 34),
		(6542, 'DUNDU (NDUNDU)', 472, 34),
		(6543, 'EGIAM', 472, 34),
		(6544, 'GEYA', 472, 34),
		(6545, 'GIKA', 472, 34),
		(6546, 'GILUBANDU (GILUMBANDU/GILIMBANDU)', 472, 34),
		(6547, 'GOYAGE', 472, 34),
		(6548, 'GUNDAGI (GUDAGE)', 472, 34),
		(6549, 'KAI', 472, 34),
		(6550, 'KAMBONERI', 472, 34),
		(6551, 'KANGGIME (KANGGIMA )', 472, 34),
		(6552, 'KARUBAGA', 472, 34),
		(6553, 'KEMBU', 472, 34),
		(6554, 'KONDAGA (KONDA)', 472, 34),
		(6555, 'KUARI', 472, 34),
		(6556, 'KUBU', 472, 34),
		(6557, 'LI ANOGOMMA', 472, 34),
		(6558, 'NABUNAGE', 472, 34),
		(6559, 'NELAWI', 472, 34),
		(6560, 'NUMBA', 472, 34),
		(6561, 'NUNGGAWI (MUNGGAWI)', 472, 34),
		(6562, 'PANAGA', 472, 34),
		(6563, 'POGANERI', 472, 34),
		(6564, 'TAGIME', 472, 34),
		(6565, 'TAGINERI', 472, 34),
		(6566, 'TELENGGEME', 472, 34),
		(6567, 'TIMORI', 472, 34),
		(6568, 'UMAGI', 472, 34),
		(6569, 'WAKUWO', 472, 34),
		(6570, 'WARI (TAIYEVE)', 472, 34),
		(6571, 'WENAM', 472, 34),
		(6572, 'WINA', 472, 34),
		(6573, 'WONOKI (WONIKI)', 472, 34),
		(6574, 'WUGI', 472, 34),
		(6575, 'WUNIN (WUMIN)', 472, 34),
		(6576, 'YUKO', 472, 34),
		(6577, 'YUNERI', 472, 34),
		(6578, 'DEMBA MASIREI', 473, 34),
		(6579, 'INGGERUS', 473, 34),
		(6580, 'KIRIHI', 473, 34),
		(6581, 'MASIREI', 473, 34),
		(6582, 'OUDATE WAROPEN', 473, 34),
		(6583, 'RISEI SAYATI', 473, 34),
		(6584, 'UREIFASEI', 473, 34),
		(6585, 'WAPOGA INGGERUS', 473, 34),
		(6586, 'WAROPEN BAWAH', 473, 34),
		(6587, 'AMUMA', 474, 34),
		(6588, 'ANGGRUK', 474, 34),
		(6589, 'BOMELA', 474, 34),
		(6590, 'DEKAI', 474, 34),
		(6591, 'DIRWEMNA (DIRUWENA)', 474, 34),
		(6592, 'DURAM', 474, 34),
		(6593, 'ENDOMEN', 474, 34),
		(6594, 'HEREAPINI (HEREANINI)', 474, 34),
		(6595, 'HILIPUK', 474, 34),
		(6596, 'HOGIO (HUGIO)', 474, 34),
		(6597, 'HOLUON', 474, 34),
		(6598, 'KABIANGGAMA (KABIANGGEMA)', 474, 34),
		(6599, 'KAYO', 474, 34),
		(6600, 'KONA', 474, 34),
		(6601, 'KOROPUN (KORUPUN)', 474, 34),
		(6602, 'KOSAREK', 474, 34),
		(6603, 'KURIMA', 474, 34),
		(6604, 'KWELEMDUA (KWELAMDUA)', 474, 34),
		(6605, 'KWIKMA', 474, 34),
		(6606, 'LANGDA', 474, 34),
		(6607, 'LOLAT', 474, 34),
		(6608, 'MUGI', 474, 34),
		(6609, 'MUSAIK', 474, 34),
		(6610, 'NALCA', 474, 34),
		(6611, 'NINIA', 474, 34),
		(6612, 'NIPSAN', 474, 34),
		(6613, 'OBIO', 474, 34),
		(6614, 'PANGGEMA', 474, 34),
		(6615, 'PASEMA', 474, 34),
		(6616, 'PRONGGOLI (PROGGOLI)', 474, 34),
		(6617, 'PULDAMA', 474, 34),
		(6618, 'SAMENAGE', 474, 34),
		(6619, 'SELA', 474, 34),
		(6620, 'SEREDELA (SEREDALA)', 474, 34),
		(6621, 'SILIMO', 474, 34),
		(6622, 'SOBA', 474, 34),
		(6623, 'SOBAHAM', 474, 34),
		(6624, 'SOLOIKMA', 474, 34),
		(6625, 'SUMO', 474, 34),
		(6626, 'SUNTAMON', 474, 34),
		(6627, 'SURU SURU', 474, 34),
		(6628, 'TALAMBO', 474, 34),
		(6629, 'TANGMA', 474, 34),
		(6630, 'UBAHAK', 474, 34),
		(6631, 'UBALIHI', 474, 34),
		(6632, 'UKHA', 474, 34),
		(6633, 'WALMA', 474, 34),
		(6634, 'WERIMA', 474, 34),
		(6635, 'WUSUMA', 474, 34),
		(6636, 'YAHULIAMBUT', 474, 34),
		(6637, 'YOGOSEM', 474, 34),
		(6638, 'ABENAHO', 475, 34),
		(6639, 'APALAPSILI', 475, 34),
		(6640, 'BENAWA', 475, 34),
		(6641, 'ELELIM', 475, 34),
		(6642, 'WELAREK', 475, 34)
	`)
}
