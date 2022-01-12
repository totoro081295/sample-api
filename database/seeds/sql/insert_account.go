package sql

var InsertAccount = ` insert into accounts
(id, email, password, created_by, updated_by, created_at, updated_at, deleted_at)
values
('708ee960-4cbb-4770-9a2f-4a632c409428', 'test@example.com', '$2a$10$mECDvehgAc6ZcTSS4y4eBeP8h5vEGm0tU0M3zHPF3BnG11GF1gsTK', 'nzJCeqDCTcRtA89UomfE2UT16kv1', 'nzJCeqDCTcRtA89UomfE2UT16kv1', now(), now(), null);`
