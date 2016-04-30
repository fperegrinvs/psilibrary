describe('Search Service tests', function() {
    test_init()

    // instantiate service
    var searchService, configSettings, baseUrl;

    beforeEach(inject(function (_$httpBackend_, _searchService_, config) {
        searchService = _searchService_;
        $httpBackend = _$httpBackend_;
        configSettings = config;
        baseUrl = config.serverUrl;
        dummy = null;
        error = null;
    }));


    it('test search, ok', function () {
       var body = {foo: "bar"};
        serviceOk(searchService.Search, body, baseUrl + '/search', body, body);
    });

    it('test search, fail', function () {
       var body = {foo: "bar"};
        serviceFail(searchService.Search, body, baseUrl + '/search', body, body);
    });
});