describe('Category Service tests', function() {
    test_init()

    // instantiate service
    var categoryService, configSettings, baseUrl;

    beforeEach(inject(function (_$httpBackend_, _categoryService_, config) {
        categoryService = _categoryService_;
        $httpBackend = _$httpBackend_;
        configSettings = config;
        baseUrl = config.serverUrl;
        dummy = null;
        error = null;
    }));


    it('test list, ok', function () {
        serviceOk(categoryService.List, null, baseUrl + '/category', {foo: "bar"});
    });

    it('test list error', function(){
        serviceFail(categoryService.List, null, baseUrl + '/category', "error");
    })

    it('test get, ok', function () {
        serviceOk(categoryService.Get, 20, baseUrl + '/category/20', {foo: "bar"});
    });

    it('test get error', function(){
        serviceFail(categoryService.Get, 20, baseUrl + '/category/20', "error");
    })

    it('test create, ok', function () {
       var body = {foo: "bar"};
        serviceOk(categoryService.Create, body, baseUrl + '/category/create', body, body);
    });

    it('test update, fail', function () {
       var body = {foo: "bar"};
        serviceFail(categoryService.Create, body, baseUrl + '/category/create', body, body);
    });

    it('test update, ok', function () {
       var body = {foo: "bar"};
        serviceOk(categoryService.Update, body, baseUrl + '/category/update', body, body);
    });

    it('test update, fail', function () {
       var body = {foo: "bar"};
        serviceFail(categoryService.Update, body, baseUrl + '/category/update', body, body);
    });
});