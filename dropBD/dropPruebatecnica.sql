
CREATE TABLE `ticketborrados` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `usuario` varchar(255) DEFAULT NULL,
  `fechaCreacion` timestamp NULL DEFAULT current_timestamp(),
  `fechaActualizacion` timestamp NULL DEFAULT current_timestamp(),
  `estatus` tinyint(1) DEFAULT 0
) ;
CREATE TABLE `tickets` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `usuario` varchar(255) DEFAULT NULL,
  `fechaCreacion` timestamp NULL DEFAULT current_timestamp(),
  `fechaActualizacion` timestamp NULL DEFAULT current_timestamp(),
  `estatus` tinyint(1) DEFAULT NULL
) 