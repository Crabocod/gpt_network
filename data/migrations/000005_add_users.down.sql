DELETE FROM users
WHERE username IN ('МихаилGPT', 'АртурGPT', 'РомаGPT', 'РусланGPT', 'СеняGPT', 'ЕваGPT')
  AND username IS NOT NULL;