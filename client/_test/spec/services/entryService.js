describe('Entry Service tests', function() {
    test_init()

    // instantiate service
    var entryService, configSettings, baseUrl;

    beforeEach(inject(function (_$httpBackend_, _entryService_, config) {
        entryService = _entryService_;
        $httpBackend = _$httpBackend_;
        configSettings = config;
        baseUrl = config.serverUrl;
        dummy = null;
        error = null;
    }));

    it('test list, ok', function () {
        serviceOk(entryService.List, null, baseUrl + '/entry', {foo: "bar"});
    });


    it('test get, ok', function () {
        serviceOk(entryService.Get, 20, baseUrl + '/entry/20', {foo: "bar"});
    });

    it('test get error', function(){
        serviceFail(entryService.Get, 20, baseUrl + '/entry/20', "error");
    })

    it('test create, ok', function () {
       var body = {foo: "bar"};
        serviceOk(entryService.Create, body, baseUrl + '/entry/create', body, body);
    });

    it('test update, fail', function () {
       var body = {foo: "bar"};
        serviceFail(entryService.Create, body, baseUrl + '/entry/create', body, body);
    });

    it('test update, ok', function () {
       var body = {foo: "bar"};
        serviceOk(entryService.Update, body, baseUrl + '/entry/update', body, body);
    });

    it('test update, fail', function () {
       var body = {foo: "bar"};
        serviceFail(entryService.Update, body, baseUrl + '/entry/update', body, body);
    });
});