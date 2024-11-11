CREATE TABLE "public"."t_transaction" (
                                          "id" int4 NOT NULL,
                                          "from_id" int4 NOT NULL,
                                          "to_id" int4 NOT NULL,
                                          "amount" numeric(10,2) NOT NULL,
                                          "create_time" timestamp(6) NOT NULL DEFAULT now(),
                                          CONSTRAINT "t_transaction_pkey" PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."t_transaction"
    OWNER TO "postgres";

COMMENT ON COLUMN "public"."t_transaction"."id" IS '主键ID';

COMMENT ON COLUMN "public"."t_transaction"."from_id" IS '发起人ID';

COMMENT ON COLUMN "public"."t_transaction"."to_id" IS '接收人ID';

COMMENT ON COLUMN "public"."t_transaction"."amount" IS '金额';

COMMENT ON COLUMN "public"."t_transaction"."create_time" IS '创建时间';

COMMENT ON TABLE "public"."t_transaction" IS '交易表';