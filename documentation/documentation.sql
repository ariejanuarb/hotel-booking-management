CREATE TABLE `booking_hotel`.`booking`
( `id` INT NOT NULL AUTO_INCREMENT ,
  `status` VARCHAR(100) NOT NULL ,
    `pic_name` VARCHAR(100) NOT NULL ,
      `pic_contact` VARCHAR(100) NOT NULL ,
        `invoice` VARCHAR(100) NOT NULL ,
          `event_start` DATETIME NOT NULL ,
            `event_end` DATETIME NOT NULL ,
              `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
                `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
                    PRIMARY KEY  (`id`)) ENGINE = InnoDB;