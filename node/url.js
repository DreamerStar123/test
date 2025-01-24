var request = require('request');
var options = {
  'method': 'GET',
  'url': 'https://www.upwork.com/jobs/~021881745060515567376',
  'headers': {
    'Cookie': '__cf_bm=5poa2G1oaQjV_tPuwWZsza97N6xRhrwYf6WvOrmvMxA-1737669517-1.0.1.1-vAqYIxlBet7jNrJrLb.Pc70YLj9dSmktyxNjoRxSDgZ17BUZpl21REMmwniTBuXj0tX8G4.C7TDdnFTHHaxp8Q; _cfuvid=0BzwmmfNBnAzjs70PNK_KrhkfaA8d3xfQy3ddcMumsM-1737668166692-0.0.1.1-604800000; country_code=US; visitor_gql_token=oauth2v2_fabd79806e271098d794111b51a8019c; visitor_id=206.217.134.34.1737668166632000; vjd_gql_token=oauth2v2_884b9b763924352edaacdef1ff4baf7a; AWSALBTG=Fivevc6n1xstCfuMg4Q/rb0tKV5+5B2kToUM9N8TVvpKZ/yIYI3K+KCqDfAWDTo8WcflueN0S8AlVUnECBC2dm9p5RLgDG7j1ZhunrevEy3gwAjiMS5FNW6JQTsrzYmLmCkpzy3Q9VnS8icKmOZ3xovzNHhPwrWKoA8oNNXivfaZ; AWSALBTGCORS=Fivevc6n1xstCfuMg4Q/rb0tKV5+5B2kToUM9N8TVvpKZ/yIYI3K+KCqDfAWDTo8WcflueN0S8AlVUnECBC2dm9p5RLgDG7j1ZhunrevEy3gwAjiMS5FNW6JQTsrzYmLmCkpzy3Q9VnS8icKmOZ3xovzNHhPwrWKoA8oNNXivfaZ; __cflb=02DiuEXPXZVk436fJfSVuuwDqLqkhavJc4NnzgNMA3cqy; cookie_domain=.upwork.com; cookie_prefix=; enabled_ff=TONB2256Air3Migration,!RMTAir3Talent,!MP16400Air3Migration,!RMTAir3Hired,!CI10857Air3Dot0,CI9570Air2Dot5,air2Dot76,!RMTAir3Offer,OTBnrOn,!RMTAir3Offers,!CI10270Air2Dot5QTAllocations,CI11132Air2Dot75,!air2Dot76Qt,!SSINavUser,i18nOn,!CI12577UniversalSearch,CI17409DarkModeUI,!i18nGA,SSINavUserBpa,JPAir3,!RMTAir3Home'
  }
};
request(options, function (error, response) {
  if (error) throw new Error(error);
  console.log(response.statusCode);
});
