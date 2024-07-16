package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upRemoveAnnotationId, downRemoveAnnotationId)
}

func upRemoveAnnotationId(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
create table annotation_dg_tmp
(
    user_id    varchar(255) default ''    not null,
    item_id    varchar(255) default ''    not null,
    item_type  varchar(255) default ''    not null,
    play_count integer      default 0,
    play_date  datetime,
    rating     integer      default 0,
    starred    bool         default FALSE not null,
    starred_at datetime,
    unique (user_id, item_id, item_type)
);

insert into annotation_dg_tmp(user_id, item_id, item_type, play_count, play_date, rating, starred, starred_at)
select user_id,
       item_id,
       item_type,
       play_count,
       play_date,
       rating,
       starred,
       starred_at
from annotation;

drop table annotation;

alter table annotation_dg_tmp
    rename to annotation;

create index annotation_play_count
    on annotation (play_count);

create index annotation_play_date
    on annotation (play_date);

create index annotation_rating
    on annotation (rating);

create index annotation_starred
    on annotation (starred);

create index annotation_starred_at
    on annotation (starred_at);

`)
	return err
}

func downRemoveAnnotationId(ctx context.Context, tx *sql.Tx) error {
	return nil
}
