-- phpMyAdmin SQL Dump
-- version 4.1.12
-- http://www.phpmyadmin.net
--
-- Client :  localhost:8889
-- Généré le :  Dim 15 Juin 2014 à 07:21
-- Version du serveur :  5.5.34
-- Version de PHP :  5.5.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Base de données :  `blackloop`
--

-- --------------------------------------------------------

--
-- Structure de la table `produit`
--

CREATE TABLE `produit` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `titre` varchar(100) NOT NULL,
  `description` text NOT NULL,
  `url_image` varchar(200) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=7 ;

--
-- Contenu de la table `produit`
--

INSERT INTO `produit` (`id`, `titre`, `description`, `url_image`, `created_at`, `updated_at`) VALUES
(1, 'Mojito', '• Réalisez la recette "Mojito" directement dans le verre. \r\n• Placer les feuilles de menthe dans le verre, ajoutez le sucre et le jus de citrons. Piler consciencieusement afin d''exprimer l''essence de la menthe mais sans la broyer. Ajouter le rhum, remplir le verre à moitié de glaçons et compléter avec de l''eau gazeuse. Mélanger doucement et servir avec une paille.\r\n• Servir dans un verre de type "tumbler". \r\n• Décor: Décorer de feuilles de menthe fraîches et d''une tranche de citron.\r\n• Bien que la recette originale ne contienne pas d''angostura, vous pouvez y ajouter quelques gouttes afin de le rendre un peu plus sec.', '/img/1.jpg', '0000-00-00 00:00:00', '2014-06-15 05:16:15'),
(2, 'Piña Colada', '• Réalisez la recette "Piña Colada" au mixer. \r\n• Dans un blender (mixer), versez tous les ingrédients, y compris les glaçons et mixez le tout. C''est prêt ! Versez dans le verre et dégustez. Peut aussi se réaliser au shaker si c''est juste pour une personne.\r\n• Servir dans un verre de type "verre à vin". \r\n• Décor: Décorer avec un morceau d''ananas et une cerise confite.\r\n• Vous pouvez ajouter une touche d''onctuosité en ajoutant une cuillère à soupe de crème fraîche dans le mixer.', '/img/2.jpg', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 'Margarita', '• Réalisez la recette "Margarita" au shaker . \r\n• Frapper les ingrédients au shaker avec des glaçons puis verser dans le verre givré au citron et au sel fin...\r\n\r\nPour givrer facilement le verre, passer le citron sur le bord du verre et tremper les bords dans le sel.\r\n• Servir dans un verre de type "verre à margarita". \r\n• Décor: Décorer d''une tranche de citron vert...', '/img/3.jpg', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(4, 'Dirty Martini', '• Réalisez la recette "Dirty Martini" au shaker . \r\n• Frappez les ingrédients avec des glaçons au shaker et servir sans glace. Utilisez la saumure des olives dénoyautées (jus de conservation) de celles-ci.\r\n• Servir dans un verre de type "verre à martini". \r\n• Décor: Deux olives dénoyautées au fond du verre\r\n• Un Martini Dry avec de la saumure qui lui donne une couleur trouble d''où le nom "dirty" (sale). Très en vogue aux USA.', '/img/4.jpg', '0000-00-00 00:00:00', '2014-06-15 05:20:11');
