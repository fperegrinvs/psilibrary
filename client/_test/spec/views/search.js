describe('Testing Search View', function () {

  // setup
  beforeEach(module('client'));

  var TestCtrl, $rootScope, $compile, createController, view, $scope;
  var results = 
        [{"id":1282,"title":"Testing","abstract":"A abstract","content":"Dummy content",
        "publishDate":"2016-04-10T00:00:00Z",
        "author":"Leonardo Stern","journal":"Some journal","entryType":{"id":0,"name":""},
        "EntryTypeId":2,"categories":null,"medlineId":""},
      {"id":1283,"title":"Testing","abstract":"A abstract","content":"Dummy content",
        "publishDate":"2016-04-10T00:00:00Z",
        "author":"Leonardo Stern","journal":"Some journal","entryType":{"id":0,"name":""},
        "EntryTypeId":2,"categories":null,"medlineId":""}];

  beforeEach(inject(function($controller, $templateCache, _$rootScope_, _$compile_, _$httpBackend_) {
    $rootScope = _$rootScope_;
    $scope = $rootScope.$new();
    $compile = _$compile_;
    $httpBackend = _$httpBackend_
    $httpBackend.whenGET(/^.*/).respond('');
    $httpBackend.whenPOST(/^.*/).respond('');

    createController = function(data) {
      var html = '<div>' + $templateCache.get('app/templates/search.html') + '</div>';
      TestCtrl = $controller('searchCtl', { $scope: $scope, $rootScope: $rootScope });

      //join objects
      for (var attrname in data) { $scope[attrname] = data[attrname]; }

      view = $compile(angular.element(html))($scope);
      $scope.$digest();
      return $scope;
    };
  }));

  // tests
  it('Checking the query field, no data', function() {
    createController({});
    expect(view.find("#query").length).toEqual(1)
  });

  it('Checking the query field reflects model data', function() {
    createController({'query': 'abstract'});
    expect(view.find("#query").val()).toEqual('abstract')
  });

  it('Should display facets', function() {
    createController({'data': {'facets': 
      [{'id': 'category', 'name': 'Categorias', 'options': [
        {'id': 2, 'name': 'PK'},
        {'id': 3, 'name': 'Teste'},
        {'id': 4, 'name': 'Outro'},
      ]}]}});
    expect(view.find('.facet-category-item').length).toEqual(3)
  });

  it('Should display options name', function() {
    createController({'data': {'facets': 
      [{'id': 'category', 'name': 'Categorias', 'options': [
        {'id': 2, 'name': 'PK'},
        {'id': 3, 'name': 'Teste'},
        {'id': 4, 'name': 'Outro'},
      ]}]}});

    expect($(view.find('.facet-category-item')[0]).text()).toEqual('PK')
  });

  it('facet option should filter', function() {
    var scope = createController({'data': {'facets': 
      [{'id': 'category', 'name': 'Categorias', 'options': [
        {'id': 2, 'name': 'PK'},
        {'id': 3, 'name': 'Teste'},
        {'id': 4, 'name': 'Outro'},
      ]}]}});

    spyOn(scope, 'search').and.callThrough();;
    $(view.find('.facet-category-item')[0]).click();
    expect(scope.search).toHaveBeenCalled();
  });

  it('should list results', function() {
    var scope = createController({'data': {'results':results}}); 

    expect(view.find('.result-item').length).toEqual(2)
  });

  it('should have 5 pages', function(){
    result = {"pagination":{"pages":[
      {},{},{},{},{}
    ]}};
    var scope = createController(result);
    expect(view.find('.page-item').length - 2).toEqual(5)

  });
});

//{"filters":{'category': ['2']}}