CREATE TABLE "public"."t_deposit" (
                                      "id" int4 NOT NULL,
                                      "user_id" int4 NOT NULL,
                                      "amount" numeric(10,2) NOT NULL,
                                      "create_time" timestamp(6) NOT NULL DEFAULT now(),
                                      CONSTRAINT "t_deposit_pkey" PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."t_deposit"
    OWNER TO "postgres";

COMMENT ON COLUMN "public"."t_deposit"."id" IS '主键ID';

COMMENT ON COLUMN "public"."t_deposit"."user_id" IS '用户ID';

COMMENT ON COLUMN "public"."t_deposit"."amount" IS '金额';

COMMENT ON COLUMN "public"."t_deposit"."create_time" IS '创建时间';

COMMENT ON TABLE "public"."t_deposit" IS '充值表';

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

CREATE TABLE "public"."t_withdraw" (
                                       "id" int4 NOT NULL,
                                       "user_id" int4 NOT NULL,
                                       "amount" numeric(10,2) NOT NULL,
                                       "create_time" timestamp(6) NOT NULL DEFAULT now(),
                                       CONSTRAINT "t_withdraw_pkey" PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."t_withdraw"
    OWNER TO "postgres";

COMMENT ON COLUMN "public"."t_withdraw"."id" IS '主键ID';

COMMENT ON COLUMN "public"."t_withdraw"."user_id" IS '用户ID';

COMMENT ON COLUMN "public"."t_withdraw"."amount" IS '金额';

COMMENT ON COLUMN "public"."t_withdraw"."create_time" IS '创建时间';

COMMENT ON TABLE "public"."t_withdraw" IS '提现表';