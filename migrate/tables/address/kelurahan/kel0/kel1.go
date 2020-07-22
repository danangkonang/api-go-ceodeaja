package kel0

import "github.com/danangkonang/rest-api/config"

func Kel1() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
		(1001, 'TIPAR JAYA (TIPARRAYA)', 94, 5, 1),
		(1002, 'ALAM JAYA', 95, 5, 1),
		(1003, 'GANDASARI', 95, 5, 1),
		(1004, 'JATAKE', 95, 5, 1),
		(1005, 'KERONCONG', 95, 5, 1),
		(1006, 'MANIS JAYA', 95, 5, 1),
		(1007, 'PASIR JAYA', 95, 5, 1),
		(1008, 'CIKANDE', 96, 5, 1),
		(1009, 'DANG DEUR', 96, 5, 1),
		(1010, 'JAYANTI', 96, 5, 1),
		(1011, 'PABUARAN', 96, 5, 1),
		(1012, 'PANGKAT', 96, 5, 1),
		(1013, 'PASIR GINTUNG', 96, 5, 1),
		(1014, 'PASIR MUNCANG', 96, 5, 1),
		(1015, 'SUMUR BANDUNG', 96, 5, 1),
		(1016, 'BOJONG JAYA', 97, 5, 1),
		(1017, 'BUGEL', 97, 5, 1),
		(1018, 'CIMONE', 97, 5, 1),
		(1019, 'CIMONE JAYA', 97, 5, 1),
		(1020, 'GERENDENG', 97, 5, 1),
		(1021, 'KARAWACI', 97, 5, 1),
		(1022, 'KARAWACI BARU', 97, 5, 1),
		(1023, 'KOANG JAYA', 97, 5, 1),
		(1024, 'MARGA SARI', 97, 5, 1),
		(1025, 'NAMBO JAYA', 97, 5, 1),
		(1026, 'NUSA JAYA', 97, 5, 1),
		(1027, 'PABUARAN', 97, 5, 1),
		(1028, 'PABUARAN TUMPENG', 97, 5, 1),
		(1029, 'PASAR BARU', 97, 5, 1),
		(1030, 'SUKA JADI', 97, 5, 1),
		(1031, 'SUMUR PANCING (PACING)', 97, 5, 1),
		(1032, 'BENCONGAN', 98, 5, 1),
		(1033, 'BENCONGAN INDAH', 98, 5, 1),
		(1034, 'BOJONG NANGKA', 98, 5, 1),
		(1035, 'CURUG SANGERANG (CURUG SANGERENG)', 98, 5, 1),
		(1036, 'KELAPA DUA', 98, 5, 1),
		(1037, 'PAKULONAN BARAT', 98, 5, 1),
		(1038, 'BELIMBING', 99, 5, 1),
		(1039, 'CENGKLONG', 99, 5, 1),
		(1040, 'DADAP', 99, 5, 1),
		(1041, 'JATIMULYA', 99, 5, 1),
		(1042, 'KOSAMBI BARAT', 99, 5, 1),
		(1043, 'KOSAMBI TIMUR', 99, 5, 1),
		(1044, 'RAWA BURUNG', 99, 5, 1),
		(1045, 'RAWA RENGAS', 99, 5, 1),
		(1046, 'SALEMBARAN JATI', 99, 5, 1),
		(1047, 'SALEMBARAN RAYA', 99, 5, 1),
		(1048, 'JENGKOL', 100, 5, 1),
		(1049, 'KEMUNING', 100, 5, 1),
		(1050, 'KOPER', 100, 5, 1),
		(1051, 'KRESEK', 100, 5, 1),
		(1052, 'PASIR AMPO', 100, 5, 1),
		(1053, 'PATRA SANA', 100, 5, 1),
		(1054, 'RANCA ILAT', 100, 5, 1),
		(1055, 'RENGED', 100, 5, 1),
		(1056, 'TALOK', 100, 5, 1),
		(1057, 'BAKUNG', 101, 5, 1),
		(1058, 'BLUKBUK', 101, 5, 1),
		(1059, 'CIRUMPAK', 101, 5, 1),
		(1060, 'KRONJO', 101, 5, 1),
		(1061, 'MUNCUNG', 101, 5, 1),
		(1062, 'PAGEDANGAN ILIR', 101, 5, 1),
		(1063, 'PAGEDANGAN UDIK', 101, 5, 1),
		(1064, 'PAGENJAHAN', 101, 5, 1),
		(1065, 'PASILIAN', 101, 5, 1),
		(1066, 'PASIR', 101, 5, 1),
		(1067, 'BABAKAN', 102, 5, 1),
		(1068, 'BABAT', 102, 5, 1),
		(1069, 'BOJONG KAMAL', 102, 5, 1),
		(1070, 'CARINGIN', 102, 5, 1),
		(1071, 'CIANGIR', 102, 5, 1),
		(1072, 'CIRARAB', 102, 5, 1),
		(1073, 'KAMUNING (KEMUNING)', 102, 5, 1),
		(1074, 'LEGOK', 102, 5, 1),
		(1075, 'PALA SARI', 102, 5, 1),
		(1076, 'RANCAGONG', 102, 5, 1),
		(1077, 'SERDANG WETAN', 102, 5, 1),
		(1078, 'BANYU ASIH', 103, 5, 1),
		(1079, 'GUNUNG SARI', 103, 5, 1),
		(1080, 'JATI WARINGIN', 103, 5, 1),
		(1081, 'KEDUNG DALEM', 103, 5, 1),
		(1082, 'KETAPANG', 103, 5, 1),
		(1083, 'MARGA MULYA', 103, 5, 1),
		(1084, 'MAUK BARAT', 103, 5, 1),
		(1085, 'MAUK TIMUR', 103, 5, 1),
		(1086, 'SASAK', 103, 5, 1),
		(1087, 'TANJUNG ANOM', 103, 5, 1),
		(1088, 'TEGAL KUNIR KIDUL', 103, 5, 1),
		(1089, 'TEGAL KUNIR LOR', 103, 5, 1),
		(1090, 'CIJERUK', 104, 5, 1),
		(1091, 'GANDA RIA', 104, 5, 1),
		(1092, 'JENGGOT', 104, 5, 1),
		(1093, 'KEDAUNG', 104, 5, 1),
		(1094, 'KLUTUK', 104, 5, 1),
		(1095, 'KOSAMBI DALAM', 104, 5, 1),
		(1096, 'MEKAR BARU', 104, 5, 1),
		(1097, 'WALIWIS', 104, 5, 1),
		(1098, 'KARANG ANYAR', 105, 5, 1),
		(1099, 'KARANG SARI', 105, 5, 1),
		(1100, 'KEDAUNG BARU', 105, 5, 1),
		(1101, 'KEDAUNG WETAN', 105, 5, 1),
		(1102, 'MEKAR SARI', 105, 5, 1),
		(1103, 'NEGLASARI', 105, 5, 1),
		(1104, 'SELAPAJANG JAYA', 105, 5, 1),
		(1105, 'CICALENGKA', 106, 5, 1),
		(1106, 'CIHUNI', 106, 5, 1),
		(1107, 'CIJANTRA', 106, 5, 1),
		(1108, 'JATAKE', 106, 5, 1),
		(1109, 'KADU SIRUNG', 106, 5, 1),
		(1110, 'KARANG TENGAH', 106, 5, 1),
		(1111, 'LENGKONG KULON', 106, 5, 1),
		(1112, 'MALANG NENGAH', 106, 5, 1),
		(1113, 'MEDANG', 106, 5, 1),
		(1114, 'PAGEDANGAN', 106, 5, 1),
		(1115, 'SITU GADUNG', 106, 5, 1),
		(1116, 'BUARAN BAMBU', 107, 5, 1),
		(1117, 'BUARAN MANGGA', 107, 5, 1),
		(1118, 'BUNISARI (BONASARI)', 107, 5, 1),
		(1119, 'GAGA', 107, 5, 1),
		(1120, 'KALIBARU', 107, 5, 1),
		(1121, 'KIARA PAYUNG', 107, 5, 1),
		(1122, 'KOHOD', 107, 5, 1),
		(1123, 'KRAMAT', 107, 5, 1),
		(1124, 'LAKSANA', 107, 5, 1),
		(1125, 'PAKU ALAM', 107, 5, 1),
		(1126, 'PAKUHAJI', 107, 5, 1),
		(1127, 'RAWA BONI', 107, 5, 1),
		(1128, 'SUKAWALI', 107, 5, 1),
		(1129, 'SURYA BAHARI', 107, 5, 1),
		(1130, 'CIAKAR', 108, 5, 1),
		(1131, 'MEKAR BAKTI', 108, 5, 1),
		(1132, 'MEKAR JAYA', 108, 5, 1),
		(1133, 'PANONGAN', 108, 5, 1),
		(1134, 'PEUSAR', 108, 5, 1),
		(1135, 'RANCA IYUH', 108, 5, 1),
		(1136, 'RANCA KALAPA', 108, 5, 1),
		(1137, 'SERDANG KULON', 108, 5, 1),
		(1138, 'GELAM JAYA', 109, 5, 1),
		(1139, 'KUTA BARU', 109, 5, 1),
		(1140, 'KUTA BUMI', 109, 5, 1),
		(1141, 'KUTA JAYA', 109, 5, 1),
		(1142, 'PANGADEGAN', 109, 5, 1),
		(1143, 'PASAR KEMIS', 109, 5, 1),
		(1144, 'SINDANG SARI', 109, 5, 1),
		(1145, 'SUKAASIH', 109, 5, 1),
		(1146, 'SUKAMANTRI', 109, 5, 1),
		(1147, 'GEBANG RAYA', 110, 5, 1),
		(1148, 'GEMBOR', 110, 5, 1),
		(1149, 'PERIUK', 110, 5, 1),
		(1150, 'PERIUK JAYA', 110, 5, 1),
		(1151, 'SANGIANG JAYA', 110, 5, 1),
		(1152, 'CIPETE', 111, 5, 1),
		(1153, 'KUNCIRAN', 111, 5, 1),
		(1154, 'KUNCIRAN INDAH', 111, 5, 1),
		(1155, 'KUNCIRAN JAYA', 111, 5, 1),
		(1156, 'NEROGTOG', 111, 5, 1),
		(1157, 'PAKOJAN', 111, 5, 1),
		(1158, 'PANUNGGANGAN', 111, 5, 1),
		(1159, 'PANUNGGANGAN TIMUR', 111, 5, 1),
		(1160, 'PANUNGGANGAN UTARA', 111, 5, 1),
		(1161, 'PINANG', 111, 5, 1),
		(1162, 'SUDIMARA PINANG', 111, 5, 1),
		(1163, 'DAON', 112, 5, 1),
		(1164, 'JAMBU KARYA', 112, 5, 1),
		(1165, 'LEMBANG SARI', 112, 5, 1),
		(1166, 'MEKARSARI', 112, 5, 1),
		(1167, 'PANGARENGAN', 112, 5, 1),
		(1168, 'RAJEG', 112, 5, 1),
		(1169, 'RAJEGMULYA', 112, 5, 1),
		(1170, 'RANCA BANGO', 112, 5, 1),
		(1171, 'SUKA MANAH', 112, 5, 1),
		(1172, 'SUKA SARI', 112, 5, 1),
		(1173, 'SUKA TANI', 112, 5, 1),
		(1174, 'TANJAKAN', 112, 5, 1),
		(1175, 'TANJAKAN MEKAR', 112, 5, 1),
		(1176, 'KARET', 113, 5, 1),
		(1177, 'KAYU AGUNG', 113, 5, 1),
		(1178, 'KAYU BONGKOK', 113, 5, 1),
		(1179, 'MEKAR JAYA', 113, 5, 1),
		(1180, 'PISANGAN JAYA', 113, 5, 1),
		(1181, 'PONDOK JAYA', 113, 5, 1),
		(1182, 'SARAKAN', 113, 5, 1),
		(1183, 'SEPATAN', 113, 5, 1),
		(1184, 'GEMPOL SARI', 114, 5, 1),
		(1185, 'JATI MULYA', 114, 5, 1),
		(1186, 'KAMPUNG KELOR', 114, 5, 1),
		(1187, 'KEDAUNG BARAT', 114, 5, 1),
		(1188, 'LEBAK WANGI', 114, 5, 1),
		(1189, 'PONDOK KELOR', 114, 5, 1),
		(1190, 'SANGIANG', 114, 5, 1),
		(1191, 'TANAH MERAH', 114, 5, 1),
		(1192, 'BADAK ANOM', 115, 5, 1),
		(1193, 'SINDANG ASIH', 115, 5, 1),
		(1194, 'SINDANG JAYA', 115, 5, 1),
		(1195, 'SINDANG PANON', 115, 5, 1),
		(1196, 'SINDANG SONO', 115, 5, 1),
		(1197, 'SUKA HARJA', 115, 5, 1),
		(1198, 'WANA KERTA', 115, 5, 1),
		(1199, 'CIKAREO', 116, 5, 1),
		(1200, 'CIKASUNGKA', 116, 5, 1),
		(1201, 'CIKUYA', 116, 5, 1),
		(1202, 'CIREUNDEU', 116, 5, 1),
		(1203, 'MUNJUL', 116, 5, 1),
		(1204, 'PASANGGRAHAN', 116, 5, 1),
		(1205, 'SOLEAR', 116, 5, 1),
		(1206, 'BUARAN JATI', 117, 5, 1),
		(1207, 'GINTUNG', 117, 5, 1),
		(1208, 'KARANG SERANG', 117, 5, 1),
		(1209, 'KOSAMBI', 117, 5, 1),
		(1210, 'MEKAR KONDANG', 117, 5, 1),
		(1211, 'PEKAYON', 117, 5, 1),
		(1212, 'RAWA KIDANG', 117, 5, 1),
		(1213, 'SUKADIRI', 117, 5, 1),
		(1214, 'BENDA', 118, 5, 1),
		(1215, 'BUNAR', 118, 5, 1),
		(1216, 'BUNI AYU', 118, 5, 1),
		(1217, 'KALI ASIN', 118, 5, 1),
		(1218, 'KUBANG', 118, 5, 1),
		(1219, 'MERAK', 118, 5, 1),
		(1220, 'PARAHU', 118, 5, 1),
		(1221, 'SUKA MULYA', 118, 5, 1),
		(1222, 'BABAKAN', 119, 5, 1),
		(1223, 'BUARAN INDAH', 119, 5, 1),
		(1224, 'CIKOKOL', 119, 5, 1),
		(1225, 'KELAPA INDAH', 119, 5, 1),
		(1226, 'SUKA ASIH', 119, 5, 1),
		(1227, 'SUKARASA', 119, 5, 1),
		(1228, 'SUKASARI', 119, 5, 1),
		(1229, 'TANAH TINGGI', 119, 5, 1),
		(1230, 'BABAKAN ASEM', 120, 5, 1),
		(1231, 'BOJONG RENGED', 120, 5, 1),
		(1232, 'KAMPUNG BESAR', 120, 5, 1),
		(1233, 'KAMPUNG MELAYU BARAT', 120, 5, 1),
		(1234, 'KAMPUNG MELAYU TIMUR', 120, 5, 1),
		(1235, 'KEBON CAU', 120, 5, 1),
		(1236, 'LEMO', 120, 5, 1),
		(1237, 'MUARA', 120, 5, 1),
		(1238, 'PANGKALAN', 120, 5, 1),
		(1239, 'TANJUNG BURUNG', 120, 5, 1),
		(1240, 'TANJUNG PASIR', 120, 5, 1),
		(1241, 'TEGAL ANGUS', 120, 5, 1),
		(1242, 'TELUK NAGA', 120, 5, 1),
		(1243, 'BANTAR PANJANG', 121, 5, 1),
		(1244, 'CILELES', 121, 5, 1),
		(1245, 'CISEREH', 121, 5, 1),
		(1246, 'KADU AGUNG', 121, 5, 1),
		(1247, 'MARGA SARI', 121, 5, 1),
		(1248, 'MATA GARA', 121, 5, 1),
		(1249, 'PASIR BOLANG', 121, 5, 1),
		(1250, 'PASIR NANGKA', 121, 5, 1),
		(1251, 'PEMATANG', 121, 5, 1),
		(1252, 'PETE', 121, 5, 1),
		(1253, 'SODONG', 121, 5, 1),
		(1254, 'TAPOS', 121, 5, 1),
		(1255, 'TEGALSARI', 121, 5, 1),
		(1256, 'TIGARAKSA', 121, 5, 1),
		(1257, 'CIPAYUNG', 122, 6, 1),
		(1258, 'CIPUTAT', 122, 6, 1),
		(1259, 'JOMBANG', 122, 6, 1),
		(1260, 'SARUA (SERUA)', 122, 6, 1),
		(1261, 'SARUA INDAH', 122, 6, 1),
		(1262, 'SAWAH BARU', 122, 6, 1),
		(1263, 'SAWAH LAMA', 122, 6, 1),
		(1264, 'CEMPAKA PUTIH', 123, 6, 1),
		(1265, 'CIREUNDEU', 123, 6, 1),
		(1266, 'PISANGAN', 123, 6, 1),
		(1267, 'PONDOK RANJI', 123, 6, 1),
		(1268, 'REMPOA', 123, 6, 1),
		(1269, 'RENGAS', 123, 6, 1),
		(1270, 'BAMBU APUS', 124, 6, 1),
		(1271, 'BENDA BARU', 124, 6, 1),
		(1272, 'KEDAUNG', 124, 6, 1),
		(1273, 'PAMULANG BARAT', 124, 6, 1),
		(1274, 'PAMULANG TIMUR', 124, 6, 1),
		(1275, 'PONDOK BENDA', 124, 6, 1),
		(1276, 'PONDOK CABE ILIR', 124, 6, 1),
		(1277, 'PONDOK CABE UDIK', 124, 6, 1),
		(1278, 'JURANG MANGU BARAT', 125, 6, 1),
		(1279, 'JURANG MANGU TIMUR', 125, 6, 1),
		(1280, 'PERIGI BARU', 125, 6, 1),
		(1281, 'PERIGI LAMA', 125, 6, 1),
		(1282, 'PONDOK AREN', 125, 6, 1),
		(1283, 'PONDOK BETUNG', 125, 6, 1),
		(1284, 'PONDOK JAYA', 125, 6, 1),
		(1285, 'PONDOK KACANG BARAT', 125, 6, 1),
		(1286, 'PONDOK KACANG TIMUR', 125, 6, 1),
		(1287, 'PONDOK KARYA', 125, 6, 1),
		(1288, 'PONDOK PUCUNG', 125, 6, 1),
		(1289, 'BUARAN', 126, 6, 1),
		(1290, 'CIATER', 126, 6, 1),
		(1291, 'CILENGGANG', 126, 6, 1),
		(1292, 'LENGKONG GUDANG', 126, 6, 1),
		(1293, 'LENGKONG GUDANG TIMUR', 126, 6, 1),
		(1294, 'LENGKONG WETAN', 126, 6, 1),
		(1295, 'RAWA BUNTU', 126, 6, 1),
		(1296, 'RAWA MEKAR JAYA', 126, 6, 1),
		(1297, 'SERPONG', 126, 6, 1),
		(1298, 'JELUPANG', 127, 6, 1),
		(1299, 'LENGKONG KARYA', 127, 6, 1),
		(1300, 'PAKU JAYA', 127, 6, 1),
		(1301, 'PAKUALAM', 127, 6, 1),
		(1302, 'PAKULONAN', 127, 6, 1),
		(1303, 'PONDOK JAGUNG', 127, 6, 1),
		(1304, 'PONDOK JAGUNG TIMUR', 127, 6, 1),
		(1305, 'CENGKARENG BARAT', 128, 7, 2),
		(1306, 'CENGKARENG TIMUR', 128, 7, 2),
		(1307, 'DURI KOSAMBI', 128, 7, 2),
		(1308, 'KAPUK', 128, 7, 2),
		(1309, 'KEDAUNG KALI ANGKE', 128, 7, 2),
		(1310, 'RAWA BUAYA', 128, 7, 2),
		(1311, 'GROGOL', 129, 7, 2),
		(1312, 'JELAMBAR', 129, 7, 2),
		(1313, 'JELAMBAR BARU', 129, 7, 2),
		(1314, 'TANJUNG DUREN SELATAN', 129, 7, 2),
		(1315, 'TANJUNG DUREN UTARA', 129, 7, 2),
		(1316, 'TOMANG', 129, 7, 2),
		(1317, 'WIJAYA KUSUMA', 129, 7, 2),
		(1318, 'KALIDERES', 130, 7, 2),
		(1319, 'KAMAL', 130, 7, 2),
		(1320, 'PEGADUNGAN', 130, 7, 2),
		(1321, 'SEMANAN', 130, 7, 2),
		(1322, 'TEGAL ALUR', 130, 7, 2),
		(1323, 'DURI KEPA', 131, 7, 2),
		(1324, 'KEBON JERUK', 131, 7, 2),
		(1325, 'KEDOYA SELATAN', 131, 7, 2),
		(1326, 'KEDOYA UTARA', 131, 7, 2),
		(1327, 'KELAPA DUA', 131, 7, 2),
		(1328, 'SUKABUMI SELATAN (UDIK)', 131, 7, 2),
		(1329, 'SUKABUMI UTARA (ILIR)', 131, 7, 2),
		(1330, 'JOGLO', 132, 7, 2),
		(1331, 'KEMBANGAN SELATAN', 132, 7, 2),
		(1332, 'KEMBANGAN UTARA', 132, 7, 2),
		(1333, 'MERUYA SELATAN (UDIK)', 132, 7, 2),
		(1334, 'MERUYA UTARA (ILIR)', 132, 7, 2),
		(1335, 'SRENGSENG', 132, 7, 2),
		(1336, 'JATIPULO', 133, 7, 2),
		(1337, 'KEMANGGISAN', 133, 7, 2),
		(1338, 'KOTA BAMBU SELATAN', 133, 7, 2),
		(1339, 'KOTA BAMBU UTARA', 133, 7, 2),
		(1340, 'PALMERAH', 133, 7, 2),
		(1341, 'SLIPI', 133, 7, 2),
		(1342, 'CEMPAKA PUTIH BARAT', 134, 8, 2),
		(1343, 'CEMPAKA PUTIH TIMUR', 134, 8, 2),
		(1344, 'RAWASARI', 134, 8, 2),
		(1345, 'CIDENG', 135, 8, 2),
		(1346, 'DURI PULO', 135, 8, 2),
		(1347, 'GAMBIR', 135, 8, 2),
		(1348, 'KEBON KELAPA', 135, 8, 2),
		(1349, 'PETOJO SELATAN', 135, 8, 2),
		(1350, 'PETOJO UTARA', 135, 8, 2),
		(1351, 'GALUR', 136, 8, 2),
		(1352, 'JOHAR BARU', 136, 8, 2),
		(1353, 'KAMPUNG RAWA', 136, 8, 2),
		(1354, 'TANAH TINGGI', 136, 8, 2),
		(1355, 'CEMPAKA BARU', 137, 8, 2),
		(1356, 'GUNUNG SAHARI SELATAN', 137, 8, 2),
		(1357, 'HARAPAN MULYA', 137, 8, 2),
		(1358, 'KEBON KOSONG', 137, 8, 2),
		(1359, 'KEMAYORAN', 137, 8, 2),
		(1360, 'SERDANG', 137, 8, 2),
		(1361, 'SUMUR BATU', 137, 8, 2),
		(1362, 'UTAN PANJANG', 137, 8, 2),
		(1363, 'CIKINI', 138, 8, 2),
		(1364, 'GONDANGDIA', 138, 8, 2),
		(1365, 'KEBON SIRIH', 138, 8, 2),
		(1366, 'MENTENG', 138, 8, 2),
		(1367, 'PEGANGSAAN', 138, 8, 2),
		(1368, 'GUNUNG SAHARI UTARA', 139, 8, 2),
		(1369, 'KARANG ANYAR', 139, 8, 2),
		(1370, 'KARTINI', 139, 8, 2),
		(1371, 'MANGGA DUA SELATAN', 139, 8, 2),
		(1372, 'PASAR BARU', 139, 8, 2),
		(1373, 'BUNGUR', 140, 8, 2),
		(1374, 'KENARI', 140, 8, 2),
		(1375, 'KRAMAT', 140, 8, 2),
		(1376, 'KWITANG', 140, 8, 2),
		(1377, 'PASEBAN', 140, 8, 2),
		(1378, 'SENEN', 140, 8, 2),
		(1379, 'CILANDAK BARAT', 141, 9, 2),
		(1380, 'CIPETE SELATAN', 141, 9, 2),
		(1381, 'GANDARIA SELATAN', 141, 9, 2),
		(1382, 'LEBAK BULUS', 141, 9, 2),
		(1383, 'PONDOK LABU', 141, 9, 2),
		(1384, 'CIGANJUR', 142, 9, 2),
		(1385, 'CIPEDAK', 142, 9, 2),
		(1386, 'JAGAKARSA', 142, 9, 2),
		(1387, 'LENTENG AGUNG', 142, 9, 2),
		(1388, 'SRENGSENG SAWAH', 142, 9, 2),
		(1389, 'TANJUNG BARAT', 142, 9, 2),
		(1390, 'CIPETE UTARA', 143, 9, 2),
		(1391, 'GANDARIA UTARA', 143, 9, 2),
		(1392, 'GUNUNG', 143, 9, 2),
		(1393, 'KRAMAT PELA', 143, 9, 2),
		(1394, 'MELAWAI', 143, 9, 2),
		(1395, 'PETOGOGAN', 143, 9, 2),
		(1396, 'PULO', 143, 9, 2),
		(1397, 'RAWA BARAT', 143, 9, 2),
		(1398, 'SELONG', 143, 9, 2),
		(1399, 'SENAYAN', 143, 9, 2),
		(1400, 'CIPULIR', 144, 9, 2),
		(1401, 'GROGOL SELATAN', 144, 9, 2),
		(1402, 'GROGOL UTARA', 144, 9, 2),
		(1403, 'KEBAYORAN LAMA SELATAN', 144, 9, 2),
		(1404, 'KEBAYORAN LAMA UTARA', 144, 9, 2),
		(1405, 'PONDOK PINANG', 144, 9, 2),
		(1406, 'BANGKA', 145, 9, 2),
		(1407, 'KUNINGAN BARAT', 145, 9, 2),
		(1408, 'MAMPANG PRAPATAN', 145, 9, 2),
		(1409, 'PELA MAMPANG', 145, 9, 2),
		(1410, 'TEGAL PARANG', 145, 9, 2),
		(1411, 'CIKOKO', 146, 9, 2),
		(1412, 'DUREN TIGA', 146, 9, 2),
		(1413, 'KALIBATA', 146, 9, 2),
		(1414, 'PANCORAN', 146, 9, 2),
		(1415, 'PENGADEGAN', 146, 9, 2),
		(1416, 'RAWA JATI', 146, 9, 2),
		(1417, 'CILANDAK TIMUR', 147, 9, 2),
		(1418, 'JATI PADANG', 147, 9, 2),
		(1419, 'KEBAGUSAN', 147, 9, 2),
		(1420, 'PASAR MINGGU', 147, 9, 2),
		(1421, 'PEJATEN BARAT', 147, 9, 2),
		(1422, 'PEJATEN TIMUR', 147, 9, 2),
		(1423, 'RAGUNAN', 147, 9, 2),
		(1424, 'BINTARO', 148, 9, 2),
		(1425, 'PESANGGRAHAN', 148, 9, 2),
		(1426, 'PETUKANGAN SELATAN', 148, 9, 2),
		(1427, 'PETUKANGAN UTARA', 148, 9, 2),
		(1428, 'ULUJAMI', 148, 9, 2),
		(1429, 'GUNTUR', 149, 9, 2),
		(1430, 'KARET', 149, 9, 2),
		(1431, 'KARET KUNINGAN', 149, 9, 2),
		(1432, 'KARET SEMANGGI', 149, 9, 2),
		(1433, 'KUNINGAN TIMUR', 149, 9, 2),
		(1434, 'MENTENG ATAS', 149, 9, 2),
		(1435, 'PASAR MANGGIS', 149, 9, 2),
		(1436, 'SETIABUDI', 149, 9, 2),
		(1437, 'BUKIT DURI', 150, 9, 2),
		(1438, 'KEBON BARU', 150, 9, 2),
		(1439, 'MANGGARAI', 150, 9, 2),
		(1440, 'MANGGARAI SELATAN', 150, 9, 2),
		(1441, 'MENTENG DALAM', 150, 9, 2),
		(1442, 'TEBET BARAT', 150, 9, 2),
		(1443, 'TEBET TIMUR', 150, 9, 2),
		(1444, 'CAKUNG BARAT', 151, 10, 2),
		(1445, 'CAKUNG TIMUR', 151, 10, 2),
		(1446, 'JATINEGARA', 151, 10, 2),
		(1447, 'PENGGILINGAN', 151, 10, 2),
		(1448, 'PULO GEBANG', 151, 10, 2),
		(1449, 'RAWA TERATE', 151, 10, 2),
		(1450, 'UJUNG MENTENG', 151, 10, 2),
		(1451, 'CIBUBUR', 152, 10, 2),
		(1452, 'CIRACAS', 152, 10, 2),
		(1453, 'KELAPA DUA WETAN', 152, 10, 2),
		(1454, 'RAMBUTAN', 152, 10, 2),
		(1455, 'SUSUKAN', 152, 10, 2),
		(1456, 'DUREN SAWIT', 153, 10, 2),
		(1457, 'KLENDER', 153, 10, 2),
		(1458, 'MALAKA JAYA', 153, 10, 2),
		(1459, 'MALAKA SARI', 153, 10, 2),
		(1460, 'PONDOK BAMBU', 153, 10, 2),
		(1461, 'PONDOK KELAPA', 153, 10, 2),
		(1462, 'PONDOK KOPI', 153, 10, 2),
		(1463, 'BALEKAMBANG', 154, 10, 2),
		(1464, 'BATUAMPAR', 154, 10, 2),
		(1465, 'CAWANG', 154, 10, 2),
		(1466, 'CILILITAN', 154, 10, 2),
		(1467, 'DUKUH', 154, 10, 2),
		(1468, 'KAMPUNG TENGAH', 154, 10, 2),
		(1469, 'KRAMAT JATI', 154, 10, 2),
		(1470, 'CIPINANG MELAYU', 155, 10, 2),
		(1471, 'HALIM PERDANA KUSUMAH', 155, 10, 2),
		(1472, 'KEBON PALA', 155, 10, 2),
		(1473, 'MAKASAR', 155, 10, 2),
		(1474, 'PINANG RANTI', 155, 10, 2),
		(1475, 'KAYU MANIS', 156, 10, 2),
		(1476, 'KEBON MANGGIS', 156, 10, 2),
		(1477, 'PAL MERIAM', 156, 10, 2),
		(1478, 'PISANGAN BARU', 156, 10, 2),
		(1479, 'UTAN KAYU SELATAN', 156, 10, 2),
		(1480, 'UTAN KAYU UTARA', 156, 10, 2),
		(1481, 'BARU', 157, 10, 2),
		(1482, 'CIJANTUNG', 157, 10, 2),
		(1483, 'GEDONG', 157, 10, 2),
		(1484, 'KALISARI', 157, 10, 2),
		(1485, 'PEKAYON', 157, 10, 2),
		(1486, 'CIPINANG', 158, 10, 2),
		(1487, 'JATI', 158, 10, 2),
		(1488, 'JATINEGARA KAUM', 158, 10, 2),
		(1489, 'KAYU PUTIH', 158, 10, 2),
		(1490, 'PISANGAN TIMUR', 158, 10, 2),
		(1491, 'PULO GADUNG', 158, 10, 2),
		(1492, 'RAWAMANGUN', 158, 10, 2),
		(1493, 'CILINCING', 159, 11, 2),
		(1494, 'KALI BARU', 159, 11, 2),
		(1495, 'MARUNDA', 159, 11, 2),
		(1496, 'ROROTAN', 159, 11, 2),
		(1497, 'SEMPER BARAT', 159, 11, 2),
		(1498, 'SEMPER TIMUR', 159, 11, 2),
		(1499, 'SUKAPURA', 159, 11, 2),
		(1500, 'KELAPA GADING BARAT', 160, 11, 2),
		(1501, 'KELAPA GADING TIMUR', 160, 11, 2),
		(1502, 'PEGANGSAAN DUA', 160, 11, 2),
		(1503, 'KOJA (UTARA  SELATAN)', 161, 11, 2),
		(1504, 'LAGOA', 161, 11, 2),
		(1505, 'RAWA BADAK SELATAN', 161, 11, 2),
		(1506, 'RAWA BADAK UTARA', 161, 11, 2),
		(1507, 'TUGU SELATAN', 161, 11, 2),
		(1508, 'TUGU UTARA', 161, 11, 2),
		(1509, 'ANCOL', 162, 11, 2),
		(1510, 'PADEMANGAN BARAT', 162, 11, 2),
		(1511, 'PADEMANGAN TIMUR', 162, 11, 2),
		(1512, 'KAMAL MUARA', 163, 11, 2),
		(1513, 'KAPUK MUARA', 163, 11, 2),
		(1514, 'PEJAGALAN', 163, 11, 2),
		(1515, 'PENJARINGAN', 163, 11, 2),
		(1516, 'PLUIT', 163, 11, 2),
		(1517, 'KEBON BAWANG', 164, 11, 2),
		(1518, 'PAPANGGO', 164, 11, 2),
		(1519, 'SUNGAI BAMBU', 164, 11, 2),
		(1520, 'SUNTER AGUNG', 164, 11, 2),
		(1521, 'SUNTER JAYA', 164, 11, 2),
		(1522, 'TANJUNG PRIOK', 164, 11, 2),
		(1523, 'WARAKAS', 164, 11, 2),
		(1524, 'PULAU PARI', 165, 12, 2),
		(1525, 'PULAU TIDUNG', 165, 12, 2),
		(1526, 'PULAU UNTUNG JAWA', 165, 12, 2),
		(1527, 'PULAU HARAPAN', 166, 12, 2),
		(1528, 'PULAU KELAPA', 166, 12, 2),
		(1529, 'PULAU PANGGANG', 166, 12, 2),
		(1530, 'CAMPAKA', 167, 13, 3),
		(1531, 'CIROYOM', 167, 13, 3),
		(1532, 'DUNGUS CARIANG', 167, 13, 3),
		(1533, 'GARUDA', 167, 13, 3),
		(1534, 'KEBON JERUK', 167, 13, 3),
		(1535, 'MALEBER (MALEER)', 167, 13, 3),
		(1536, 'ANTAPANI KIDUL', 168, 13, 3),
		(1537, 'ANTAPANI KULON', 168, 13, 3),
		(1538, 'ANTAPANI TENGAH', 168, 13, 3),
		(1539, 'ANTAPANI WETAN', 168, 13, 3),
		(1540, 'CISARANTEN KULON', 169, 13, 3),
		(1541, 'CISARENTEN BINA HARAPAN', 169, 13, 3),
		(1542, 'SINDANG JAYA', 169, 13, 3),
		(1543, 'SUKAMISKIN', 169, 13, 3),
		(1544, 'ANCOLMEKAR', 170, 13, 3),
		(1545, 'ARJASARI', 170, 13, 3),
		(1546, 'BAROS', 170, 13, 3),
		(1547, 'BATUKARUT', 170, 13, 3),
		(1548, 'LEBAKWANGI', 170, 13, 3),
		(1549, 'MANGUNJAYA', 170, 13, 3),
		(1550, 'MEKARJAYA', 170, 13, 3),
		(1551, 'PATROLSARI', 170, 13, 3),
		(1552, 'PINGGIRSARI', 170, 13, 3),
		(1553, 'RANCAKOLE', 170, 13, 3),
		(1554, 'WARGALUYU', 170, 13, 3),
		(1555, 'CIBADAK', 171, 13, 3),
		(1556, 'KARANGANYAR', 171, 13, 3),
		(1557, 'KARASAK', 171, 13, 3),
		(1558, 'NYENGSERET', 171, 13, 3),
		(1559, 'PANJUNAN', 171, 13, 3),
		(1560, 'PELINDUNG HEWAN', 171, 13, 3),
		(1561, 'BABAKAN', 172, 13, 3),
		(1562, 'BABAKAN CIPARAY', 172, 13, 3),
		(1563, 'CIRANGRANG', 172, 13, 3),
		(1564, 'MARGAHAYU UTARA', 172, 13, 3),
		(1565, 'MARGASUKA', 172, 13, 3),
		(1566, 'SUKAHAJI', 172, 13, 3),
		(1567, 'ANDIR', 173, 13, 3),
		(1568, 'BALEENDAH', 173, 13, 3),
		(1569, 'BOJONGMALAKA', 173, 13, 3),
		(1570, 'JELEKONG', 173, 13, 3),
		(1571, 'MALAKASARI', 173, 13, 3),
		(1572, 'MANGGAHANG', 173, 13, 3),
		(1573, 'RANCAMANYAR', 173, 13, 3),
		(1574, 'WARGAMEKAR', 173, 13, 3),
		(1575, 'BATUNUNGGAL', 174, 13, 3),
		(1576, 'KUJANGSARI', 174, 13, 3),
		(1577, 'MENGGER', 174, 13, 3),
		(1578, 'WATES', 174, 13, 3),
		(1579, 'CARINGIN', 175, 13, 3),
		(1580, 'CIBUNTU', 175, 13, 3),
		(1581, 'CIGONDEWAH KALER', 175, 13, 3),
		(1582, 'CIGONDEWAH KIDUL', 175, 13, 3),
		(1583, 'CIGONDEWAH RAHAYU', 175, 13, 3),
		(1584, 'CIJERAH', 175, 13, 3),
		(1585, 'GEMPOLSARI', 175, 13, 3),
		(1586, 'WARUNG MUNCANG', 175, 13, 3),
		(1587, 'CIHAPIT', 176, 13, 3),
		(1588, 'CITARUM', 176, 13, 3),
		(1589, 'TAMANSARI', 176, 13, 3),
		(1590, 'BINONG', 177, 13, 3),
		(1591, 'CIBANGKONG', 177, 13, 3),
		(1592, 'GUMURUH', 177, 13, 3),
		(1593, 'KACAPIRING', 177, 13, 3),
		(1594, 'KEBON GEDANG', 177, 13, 3),
		(1595, 'KEBONWARU', 177, 13, 3),
		(1596, 'MALEER', 177, 13, 3),
		(1597, 'SAMOJA', 177, 13, 3),
		(1598, 'BABAKAN ASIH', 178, 13, 3),
		(1599, 'BABAKAN TAROGONG', 178, 13, 3),
		(1600, 'JAMIKA', 178, 13, 3),
		(1601, 'KOPO', 178, 13, 3),
		(1602, 'SUKA ASIH', 178, 13, 3),
		(1603, 'CIBADUYUT', 179, 13, 3),
		(1604, 'CIBADUYUT KIDUL', 179, 13, 3),
		(1605, 'CIBADUYUT WETAN', 179, 13, 3),
		(1606, 'KEBON LEGA', 179, 13, 3),
		(1607, 'MEKARWANGI', 179, 13, 3),
		(1608, 'SITUSAEUR', 179, 13, 3),
		(1609, 'BOJONGSARI', 180, 13, 3),
		(1610, 'BOJONGSOANG', 180, 13, 3),
		(1611, 'BUAHBATU', 180, 13, 3),
		(1612, 'CIPAGALO', 180, 13, 3),
		(1613, 'LENGKONG', 180, 13, 3),
		(1614, 'TEGALLUAR', 180, 13, 3),
		(1615, 'CIJAURA (MARGASENANG)', 181, 13, 3),
		(1616, 'JATISARI', 181, 13, 3),
		(1617, 'MARGASARI', 181, 13, 3),
		(1618, 'SEKEJATI', 181, 13, 3),
		(1619, 'BANDASARI', 182, 13, 3),
		(1620, 'CANGKUANG', 182, 13, 3),
		(1621, 'CILUNCAT', 182, 13, 3),
		(1622, 'JATISARI', 182, 13, 3),
		(1623, 'NAGRAK', 182, 13, 3),
		(1624, 'PANANJUNG', 182, 13, 3),
		(1625, 'TANJUNGSARI', 182, 13, 3),
		(1626, 'CIGADUNG', 183, 13, 3),
		(1627, 'CIHAUR GEULIS', 183, 13, 3),
		(1628, 'NEGLASARI', 183, 13, 3),
		(1629, 'SUKALUYU', 183, 13, 3),
		(1630, 'CICADAS', 184, 13, 3),
		(1631, 'CIKUTRA', 184, 13, 3),
		(1632, 'PADASUKA', 184, 13, 3),
		(1633, 'PASIRLAYUNG', 184, 13, 3),
		(1634, 'SUKAMAJU', 184, 13, 3),
		(1635, 'SUKAPADA', 184, 13, 3),
		(1636, 'CIPADUNG', 185, 13, 3),
		(1637, 'CISURUPAN', 185, 13, 3),
		(1638, 'PALASARI', 185, 13, 3),
		(1639, 'PASIRBIRU', 185, 13, 3),
		(1640, 'BABAKAN PEUTEUY', 186, 13, 3),
		(1641, 'CICALENGKA KULON', 186, 13, 3),
		(1642, 'CICALENGKA WETAN', 186, 13, 3),
		(1643, 'CIKUYA', 186, 13, 3),
		(1644, 'DAMPIT', 186, 13, 3),
		(1645, 'MARGAASIH', 186, 13, 3),
		(1646, 'NAGROG', 186, 13, 3),
		(1647, 'NARAWITA', 186, 13, 3),
		(1648, 'PANENJOAN', 186, 13, 3),
		(1649, 'TANJUNGWANGI', 186, 13, 3),
		(1650, 'TENJOLAYA', 186, 13, 3),
		(1651, 'WALUYA', 186, 13, 3),
		(1652, 'ARJUNA', 187, 13, 3),
		(1653, 'HUSEN SASTRANEGARA', 187, 13, 3),
		(1654, 'PAJAJARAN', 187, 13, 3),
		(1655, 'PAMOYANAN', 187, 13, 3),
		(1656, 'PASIR KALIKI', 187, 13, 3),
		(1657, 'SUKARAJA', 187, 13, 3),
		(1658, 'CIHANYIR', 188, 13, 3),
		(1659, 'CIKANCUNG', 188, 13, 3),
		(1660, 'CIKASUNGKA', 188, 13, 3),
		(1661, 'CILULUK', 188, 13, 3),
		(1662, 'HEGARMANAH', 188, 13, 3),
		(1663, 'MANDALASARI', 188, 13, 3),
		(1664, 'MEKARLAKSANA', 188, 13, 3),
		(1665, 'SRIRAHAYU', 188, 13, 3),
		(1666, 'TANJUNGLAYA', 188, 13, 3),
		(1667, 'CILENGKRANG', 189, 13, 3),
		(1668, 'CIPANJALU', 189, 13, 3),
		(1669, 'CIPOREAT', 189, 13, 3),
		(1670, 'GIRIMEKAR', 189, 13, 3),
		(1671, 'JATIENDAH', 189, 13, 3),
		(1672, 'MELATIWANGI', 189, 13, 3),
		(1673, 'CIBIRU HILIR', 190, 13, 3),
		(1674, 'CIBIRU WETAN', 190, 13, 3),
		(1675, 'CILEUNYI KULON', 190, 13, 3),
		(1676, 'CILEUNYI WETAN', 190, 13, 3),
		(1677, 'CIMEKAR', 190, 13, 3),
		(1678, 'CINUNUK', 190, 13, 3),
		(1679, 'CAMPAKAMULYA', 191, 13, 3),
		(1680, 'CIKALONG', 191, 13, 3),
		(1681, 'CIMAUNG', 191, 13, 3),
		(1682, 'CIPINANG', 191, 13, 3),
		(1683, 'JAGABAYA', 191, 13, 3),
		(1684, 'MALASARI', 191, 13, 3),
		(1685, 'MEKARSARI', 191, 13, 3),
		(1686, 'PASIRHUNI', 191, 13, 3),
		(1687, 'SUKAMAJU', 191, 13, 3),
		(1688, 'WARJABAKTI', 191, 13, 3),
		(1689, 'CIBEUNYING', 192, 13, 3),
		(1690, 'CIBURIAL', 192, 13, 3),
		(1691, 'CIKADUT', 192, 13, 3),
		(1692, 'CIMENYAN', 192, 13, 3),
		(1693, 'MANDALAMEKAR', 192, 13, 3),
		(1694, 'MEKARMANIK', 192, 13, 3),
		(1695, 'MEKARSALUYU', 192, 13, 3),
		(1696, 'PADASUKA', 192, 13, 3),
		(1697, 'SINDANGLAYA', 192, 13, 3),
		(1698, 'BABAKAN PENGHULU', 193, 13, 3),
		(1699, 'CISARANTEN WETAN', 193, 13, 3),
		(1700, 'PAKEMITAN', 193, 13, 3),
		(1701, 'SUKAMULYA', 193, 13, 3),
		(1702, 'BABAKAN', 194, 13, 3),
		(1703, 'BUMIWANGI', 194, 13, 3),
		(1704, 'CIHEULANG', 194, 13, 3),
		(1705, 'CIKONENG', 194, 13, 3),
		(1706, 'CIPARAY', 194, 13, 3),
		(1707, 'GUNUNGLEUTIK', 194, 13, 3),
		(1708, 'MANGGUNGHARJA', 194, 13, 3),
		(1709, 'MEKAR LAKSANA', 194, 13, 3),
		(1710, 'MEKARSARI', 194, 13, 3),
		(1711, 'PAKUTANDANG', 194, 13, 3),
		(1712, 'SARIMAHI', 194, 13, 3),
		(1713, 'SERANGMEKAR', 194, 13, 3),
		(1714, 'SIGARACIPTA', 194, 13, 3),
		(1715, 'SUMBERSARI', 194, 13, 3),
		(1716, 'CIWIDEY', 195, 13, 3),
		(1717, 'LEBAKMUNCANG', 195, 13, 3),
		(1718, 'NENGKELAN', 195, 13, 3),
		(1719, 'PANUNDAAN', 195, 13, 3),
		(1720, 'PANYOCOKAN', 195, 13, 3),
		(1721, 'RAWABOGO', 195, 13, 3),
		(1722, 'SUKAWENING', 195, 13, 3),
		(1723, 'CIPAGANTI', 196, 13, 3),
		(1724, 'DAGO', 196, 13, 3),
		(1725, 'LEBAK GEDE', 196, 13, 3),
		(1726, 'LEBAK SILIWANGI', 196, 13, 3),
		(1727, 'SADANG SERANG', 196, 13, 3),
		(1728, 'SEKELOA', 196, 13, 3),
		(1729, 'CANGKUANG KULON', 197, 13, 3),
		(1730, 'CANGKUANG WETAN', 197, 13, 3),
		(1731, 'CITEUREUP', 197, 13, 3),
		(1732, 'DAYEUHKOLOT', 197, 13, 3),
		(1733, 'PASAWAHAN', 197, 13, 3),
		(1734, 'SUKAPURA', 197, 13, 3),
		(1735, 'CIMENERANG (CIMINCRANG)', 198, 13, 3),
		(1736, 'CISARANTEN KIDUL', 198, 13, 3),
		(1737, 'RANCABALONG', 198, 13, 3),
		(1738, 'RANCANUMPANG', 198, 13, 3),
		(1739, 'CIBEET', 199, 13, 3),
		(1740, 'DUKUH', 199, 13, 3),
		(1741, 'IBUN', 199, 13, 3),
		(1742, 'KARYALAKSANA', 199, 13, 3),
		(1743, 'LAKSANA', 199, 13, 3),
		(1744, 'LAMPEGAN', 199, 13, 3),
		(1745, 'MEKARWANGI', 199, 13, 3),
		(1746, 'NEGLASARI', 199, 13, 3),
		(1747, 'PANGGUH', 199, 13, 3),
		(1748, 'SUDI', 199, 13, 3),
		(1749, 'TALUN', 199, 13, 3),
		(1750, 'TANGGULUN', 199, 13, 3),
		(1751, 'BANYUSARI', 200, 13, 3),
		(1752, 'CILAMPENI', 200, 13, 3),
		(1753, 'GANDASARI', 200, 13, 3),
		(1754, 'KATAPANG', 200, 13, 3),
		(1755, 'PANGAUBAN', 200, 13, 3),
		(1756, 'SANGKANHURIP', 200, 13, 3),
		(1757, 'SUKAMUKTI', 200, 13, 3),
		(1758, 'CIBEUREUM', 201, 13, 3),
		(1759, 'CIHAWUK', 201, 13, 3),
		(1760, 'CIKEMBANG', 201, 13, 3),
		(1761, 'NEGLAWANGI', 201, 13, 3),
		(1762, 'RESMI TINGAL', 201, 13, 3),
		(1763, 'SANTOSA', 201, 13, 3),
		(1764, 'SUKAPURA', 201, 13, 3),
		(1765, 'TARUMAJAYA', 201, 13, 3),
		(1766, 'BABAKAN SARI', 202, 13, 3),
		(1767, 'BABAKAN SURABAYA', 202, 13, 3),
		(1768, 'CICAHEUM', 202, 13, 3),
		(1769, 'KEBON KANGKUNG', 202, 13, 3),
		(1770, 'KEBUN JAYANTI', 202, 13, 3),
		(1771, 'SUKAPURA', 202, 13, 3),
		(1772, 'BUNINAGARA', 203, 13, 3),
		(1773, 'CIBODAS', 203, 13, 3),
		(1774, 'CILAME', 203, 13, 3),
		(1775, 'GAJAHMEKAR', 203, 13, 3),
		(1776, 'JATISARI', 203, 13, 3),
		(1777, 'JELEGONG', 203, 13, 3),
		(1778, 'KOPO', 203, 13, 3),
		(1779, 'KUTAWARINGIN', 203, 13, 3),
		(1780, 'PADASUKA', 203, 13, 3),
		(1781, 'PAMEUNTASAN', 203, 13, 3),
		(1782, 'SUKAMULYA', 203, 13, 3),
		(1783, 'JATIHANDAP', 204, 13, 3),
		(1784, 'KARANG PAMULANG', 204, 13, 3),
		(1785, 'PASIR IMPUN', 204, 13, 3),
		(1786, 'SINDANG JAYA', 204, 13, 3),
		(1787, 'CIGONDEWAH HILIR', 205, 13, 3),
		(1788, 'LAGADAR', 205, 13, 3),
		(1789, 'MARGAASIH', 205, 13, 3),
		(1790, 'MEKAR RAHAYU', 205, 13, 3),
		(1791, 'NANJUNG', 205, 13, 3),
		(1792, 'RAHAYU', 205, 13, 3),
		(1793, 'MARGAHAYU SELATAN', 206, 13, 3),
		(1794, 'MARGAHAYU TENGAH', 206, 13, 3),
		(1795, 'SAYATI', 206, 13, 3),
		(1796, 'SUKAMENAK', 206, 13, 3),
		(1797, 'SULAEMAN', 206, 13, 3),
		(1798, 'BOJONG', 207, 13, 3),
		(1799, 'CIARO', 207, 13, 3),
		(1800, 'CIHERANG', 207, 13, 3),
		(1801, 'CITAMAN', 207, 13, 3),
		(1802, 'GANJAR SABAR', 207, 13, 3),
		(1803, 'MANDALAWANGI', 207, 13, 3),
		(1804, 'NAGREG', 207, 13, 3),
		(1805, 'NAGREG KENDAN', 207, 13, 3),
		(1806, 'BANJARSARI', 208, 13, 3),
		(1807, 'LAMAJANG', 208, 13, 3),
		(1808, 'MARGALUYU', 208, 13, 3),
		(1809, 'MARGAMEKAR', 208, 13, 3),
		(1810, 'MARGAMUKTI', 208, 13, 3),
		(1811, 'MARGAMULYA', 208, 13, 3),
		(1812, 'PANGALENGAN', 208, 13, 3),
		(1813, 'PULOSARI', 208, 13, 3),
		(1814, 'SUKALUYU', 208, 13, 3),
		(1815, 'SUKAMANAH', 208, 13, 3),
		(1816, 'TRIBAKTIMULYA', 208, 13, 3),
		(1817, 'WANASUKA', 208, 13, 3),
		(1818, 'WARNASARI', 208, 13, 3),
		(1819, 'CIPADUNG KIDUL', 209, 13, 3),
		(1820, 'CIPADUNG KULON', 209, 13, 3),
		(1821, 'CIPADUNG WETAN', 209, 13, 3),
		(1822, 'MEKARMULYA', 209, 13, 3),
		(1823, 'CIBODAS', 210, 13, 3),
		(1824, 'CIKONENG', 210, 13, 3),
		(1825, 'CISONDARI', 210, 13, 3),
		(1826, 'CUKANGGENTENG', 210, 13, 3),
		(1827, 'MARGAMULYA', 210, 13, 3),
		(1828, 'MEKARMAJU', 210, 13, 3),
		(1829, 'MEKARSARI', 210, 13, 3),
		(1830, 'PASIRJAMBU', 210, 13, 3),
		(1831, 'SUGIHMUKTI', 210, 13, 3),
		(1832, 'TENJOLAYA', 210, 13, 3),
		(1833, 'ALAMENDAH', 211, 13, 3),
		(1834, 'CIPELAH', 211, 13, 3),
		(1835, 'INDRAGIRI', 211, 13, 3),
		(1836, 'PATENGAN', 211, 13, 3),
		(1837, 'SUKARESMI', 211, 13, 3),
		(1838, 'BOJONGLOA', 212, 13, 3),
		(1839, 'BOJONGSALAM', 212, 13, 3),
		(1840, 'CANGKUANG', 212, 13, 3),
		(1841, 'HAURPUGUR', 212, 13, 3),
		(1842, 'JELEGONG', 212, 13, 3),
		(1843, 'LINGGAR', 212, 13, 3),
		(1844, 'NANJUNGMEKAR', 212, 13, 3),
		(1845, 'RANCAEKEK KENCANA', 212, 13, 3),
		(1846, 'RANCAEKEK KULON', 212, 13, 3),
		(1847, 'RANCAEKEK WETAN', 212, 13, 3),
		(1848, 'SANGIANG', 212, 13, 3),
		(1849, 'SUKAMANAH', 212, 13, 3),
		(1850, 'SUKAMULYA', 212, 13, 3),
		(1851, 'TEGAL SUMEDANG', 212, 13, 3),
		(1852, 'CIPAMOKOLAN', 213, 13, 3),
		(1853, 'DARWATI', 213, 13, 3),
		(1854, 'MANJAHLEGA (CISARANTENKIDUL)', 213, 13, 3),
		(1855, 'MEKAR MULYA (MEKARJAYA)', 213, 13, 3),
		(1856, 'ANCOL', 214, 13, 3),
		(1857, 'BALONG GEDE', 214, 13, 3),
		(1858, 'CIATEUL', 214, 13, 3),
		(1859, 'CIGERELENG', 214, 13, 3),
		(1860, 'CISEUREUH', 214, 13, 3),
		(1861, 'PASIRLUYU', 214, 13, 3),
		(1862, 'PUNGKUR', 214, 13, 3),
		(1863, 'BOJONGEMAS', 215, 13, 3),
		(1864, 'CIBODAS', 215, 13, 3),
		(1865, 'LANGENSARI', 215, 13, 3),
		(1866, 'PADAMUKTI', 215, 13, 3),
		(1867, 'PANYADAP', 215, 13, 3),
		(1868, 'RANCAKASUMBA', 215, 13, 3),
		(1869, 'SOLOKAN JERUK', 215, 13, 3),
		(1870, 'CIPEDES', 216, 13, 3),
		(1871, 'HARJOSARI', 216, 13, 3),
		(1872, 'JADIREJO', 216, 13, 3),
		(1873, 'KAMPUNG MELAYU', 216, 13, 3),
		(1874, 'KAMPUNG TENGAH', 216, 13, 3),
		(1875, 'KEDUNGSARI', 216, 13, 3),
		(1876, 'PASTEUR', 216, 13, 3),
		(1877, 'PULAU KARAM', 216, 13, 3),
		(1878, 'SUKABUNGAH', 216, 13, 3),
		(1879, 'SUKAGALIH', 216, 13, 3),
		(1880, 'SUKAJADI', 216, 13, 3),
		(1881, 'SUKAWARNA', 216, 13, 3),
		(1882, 'BABAKAN CIAMIS', 217, 13, 3),
		(1883, 'BRAGA', 217, 13, 3),
		(1884, 'KEBON PISANG', 217, 13, 3),
		(1885, 'MERDEKA', 217, 13, 3),
		(1886, 'CIGENDING', 218, 13, 3),
		(1887, 'PASANGGRAHAN', 218, 13, 3),
		(1888, 'PASIR ENDAH', 218, 13, 3),
		(1889, 'PASIRJATI', 218, 13, 3),
		(1890, 'PASIRWANGI', 218, 13, 3),
		(1891, 'BATUJAJAR BARAT', 219, 14, 3),
		(1892, 'BATUJAJAR TIMUR', 219, 14, 3),
		(1893, 'CANGKORAH', 219, 14, 3),
		(1894, 'GALANGGANG', 219, 14, 3),
		(1895, 'GIRIASIH', 219, 14, 3),
		(1896, 'PANGAUBAN', 219, 14, 3),
		(1897, 'SELACAU', 219, 14, 3),
		(1898, 'CIHAMPELAS', 220, 14, 3),
		(1899, 'CIPATIK', 220, 14, 3),
		(1900, 'CITAPEN', 220, 14, 3),
		(1901, 'MEKARJAYA', 220, 14, 3),
		(1902, 'MEKARMUKTI', 220, 14, 3),
		(1903, 'PATARUMAN', 220, 14, 3),
		(1904, 'SINGAJAYA', 220, 14, 3),
		(1905, 'SITUWANGI', 220, 14, 3),
		(1906, 'TANJUNGJAYA', 220, 14, 3),
		(1907, 'TANJUNGWANGI', 220, 14, 3),
		(1908, 'CIKALONG', 221, 14, 3),
		(1909, 'CIPADA', 221, 14, 3),
		(1910, 'CIPTAGUMATI', 221, 14, 3),
		(1911, 'CISOMANG BARAT', 221, 14, 3),
		(1912, 'GANJARSARI', 221, 14, 3),
		(1913, 'KANANGASARI', 221, 14, 3),
		(1914, 'MANDALAMUKTI', 221, 14, 3),
		(1915, 'MANDALASARI', 221, 14, 3),
		(1916, 'MEKARJAYA', 221, 14, 3),
		(1917, 'PUTERAN', 221, 14, 3),
		(1918, 'RENDE', 221, 14, 3),
		(1919, 'TENJOLAUT', 221, 14, 3),
		(1920, 'WANGUNJAYA', 221, 14, 3),
		(1921, 'BATULAYANG', 222, 14, 3),
		(1922, 'BONGAS', 222, 14, 3),
		(1923, 'BUDIHARJA', 222, 14, 3),
		(1924, 'CILILIN', 222, 14, 3),
		(1925, 'KARANGANYAR', 222, 14, 3),
		(1926, 'KARANGTANJUNG', 222, 14, 3),
		(1927, 'KARYAMUKTI', 222, 14, 3),
		(1928, 'KIDANGPANANJUNG', 222, 14, 3),
		(1929, 'MUKAPAYUNG', 222, 14, 3),
		(1930, 'NANGGERANG', 222, 14, 3),
		(1931, 'RANCAPANGGUNG', 222, 14, 3),
		(1932, 'CIPATAT', 223, 14, 3),
		(1933, 'CIPTAHARJA', 223, 14, 3),
		(1934, 'CIRAWAMEKAR', 223, 14, 3),
		(1935, 'CITATAH', 223, 14, 3),
		(1936, 'GUNUNGMASIGIT', 223, 14, 3),
		(1937, 'KERTAMUKTI', 223, 14, 3),
		(1938, 'MANDALASARI', 223, 14, 3),
		(1939, 'MANDALAWANGI', 223, 14, 3),
		(1940, 'NYALINDUNG', 223, 14, 3),
		(1941, 'RAJAMANDALA KULON', 223, 14, 3),
		(1942, 'SARIMUKTI', 223, 14, 3),
		(1943, 'SUMURBANDUNG', 223, 14, 3),
		(1944, 'BARANANGSIANG', 224, 14, 3),
		(1945, 'CIBENDA', 224, 14, 3),
		(1946, 'CICANGKANG HILIR', 224, 14, 3),
		(1947, 'CIJAMBU', 224, 14, 3),
		(1948, 'CIJENUK', 224, 14, 3),
		(1949, 'CINTAASIH', 224, 14, 3),
		(1950, 'CITALEM', 224, 14, 3),
		(1951, 'GIRIMUKTI', 224, 14, 3),
		(1952, 'KARANGSARI', 224, 14, 3),
		(1953, 'MEKARSARI', 224, 14, 3),
		(1954, 'NEGLASARI', 224, 14, 3),
		(1955, 'SARINAGEN', 224, 14, 3),
		(1956, 'SIRNAGALIH', 224, 14, 3),
		(1957, 'SUKAMULYA', 224, 14, 3),
		(1958, 'BUNIJAYA', 225, 14, 3),
		(1959, 'CELAK', 225, 14, 3),
		(1960, 'CILANGARI', 225, 14, 3),
		(1961, 'GUNUNGHALU', 225, 14, 3),
		(1962, 'SINDANGJAYA', 225, 14, 3),
		(1963, 'SIRNAJAYA', 225, 14, 3),
		(1964, 'SUKASARI', 225, 14, 3),
		(1965, 'TAMANJAYA', 225, 14, 3),
		(1966, 'WARGASALUYU', 225, 14, 3),
		(1967, 'BOJONGKONENG', 226, 14, 3),
		(1968, 'CILAME', 226, 14, 3),
		(1969, 'CIMANGGU', 226, 14, 3),
		(1970, 'CIMAREME', 226, 14, 3),
		(1971, 'GADOBANGKONG', 226, 14, 3),
		(1972, 'MARGAJAYA', 226, 14, 3),
		(1973, 'MEKARSARI', 226, 14, 3),
		(1974, 'NGAMPRAH', 226, 14, 3),
		(1975, 'PAKUHAJI', 226, 14, 3),
		(1976, 'SUKATANI', 226, 14, 3),
		(1977, 'TANIMULYA', 226, 14, 3),
		(1978, 'CEMPAKAMEKAR', 227, 14, 3),
		(1979, 'CIBURUY', 227, 14, 3),
		(1980, 'CIMERANG', 227, 14, 3),
		(1981, 'CIPEUNDEUY', 227, 14, 3),
		(1982, 'JAYAMEKAR', 227, 14, 3),
		(1983, 'KERTAJAYA', 227, 14, 3),
		(1984, 'KERTAMULYA', 227, 14, 3),
		(1985, 'LAKSANAMEKAR', 227, 14, 3),
		(1986, 'PADALARANG', 227, 14, 3),
		(1987, 'TAGOGAPU', 227, 14, 3),
		(1988, 'CIGUGUR GIRANG', 228, 14, 3),
		(1989, 'CIHANJUANG', 228, 14, 3),
		(1990, 'CIHANJUANG RAHAYU', 228, 14, 3),
		(1991, 'CIHIDEUNG', 228, 14, 3),
		(1992, 'CIWARUGA', 228, 14, 3),
		(1993, 'KARYAWANGI', 228, 14, 3),
		(1994, 'SARIWANGI', 228, 14, 3),
		(1995, 'BOJONG', 229, 14, 3),
		(1996, 'BOJONGSALAM', 229, 14, 3),
		(1997, 'CIBEDUG', 229, 14, 3),
		(1998, 'CIBITUNG', 229, 14, 3),
		(1999, 'CICADAS', 229, 14, 3),
		(2000, 'CINENGAH', 229, 14, 3)
	`)
}