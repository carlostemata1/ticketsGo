CREATE TABLE `tickets` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `usuario` varchar(255),
  `fechaCreacion` date,
  `fechaActualizacion` date,
  `estatus` boolean
);

CREATE TABLE `ticketBorrados` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `usuario` varchar(255),
  `fechaCreacion` date,
  `fechaActualizacion` date,
  `estatus` boolean
);
