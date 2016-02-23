describe('Category Service tests', function() {
  beforeEach(module('client'))
 	//var service, dummy, http;

    // instantiate service
    var categoryService, $httpBackend, configSettings, baseUrl, dummy, error;

    beforeEach(inject(function (_$httpBackend_, _categoryService_, config) {
        categoryService = _categoryService_;
        $httpBackend = _$httpBackend_;
        configSettings = config;
        baseUrl = config.serverUrl;
        dummy = null;
        error = null;
    }));

    getResult = function(data){
      dummy = data;
    }

    getError = function(e){
        error = e;
    }

    serviceOk = function(method, param, url, response, body){
        var call = !body ? $httpBackend.expectGET(url) : $httpBackend.expectPOST(url, body);
        call.respond(200, response);

        method(param).then(getResult, getError);

        $httpBackend.flush();
        expect(dummy).not.toBe(null);
        expect(error).toBe(null)
      }

    serviceFail = function(method, param, url, response, body){
        var call = !body ? $httpBackend.expectGET(url) : $httpBackend.expectPOST(url, body);
        call.respond(500, response);

        method(param).then(getResult, getError);

        $httpBackend.flush();
        expect(dummy).toBe(null);
        expect(error).not.toBe(undefined)
      }


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