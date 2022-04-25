create table if not exists url_redirect
(
    incomming_url varchar,
    out_url       varchar,
    code         integer
);

insert into url_redirect(incomming_url, out_url, code) VALUES ('/hello', '/good/byu', 302);
insert into url_redirect(incomming_url, out_url, code) VALUES ('/hello/again', '/hello', 301);
insert into url_redirect(incomming_url, out_url, code) VALUES ('/good/', '/good/byu', 301);
insert into url_redirect(incomming_url, out_url, code) VALUES ('{buy/\.*}', '/byu', 302);




