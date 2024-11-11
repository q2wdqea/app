CREATE TABLE "public"."t_wallet" (
                                     "id" int4 NOT NULL,
                                     "user_id" int4 NOT NULL,
                                     "balance" numeric(10,2) NOT NULL DEFAULT 0,
                                     "create_time" timestamp(6) NOT NULL DEFAULT now(),
                                     CONSTRAINT "wallet_pkey" PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."t_wallet"
    OWNER TO "postgres";

COMMENT ON COLUMN "public"."t_wallet"."id" IS '主键ID';

COMMENT ON COLUMN "public"."t_wallet"."user_id" IS '用户ID';

COMMENT ON COLUMN "public"."t_wallet"."balance" IS '余额';

COMMENT ON COLUMN "public"."t_wallet"."create_time" IS '创建时间';

COMMENT ON TABLE "public"."t_wallet" IS '钱包表';