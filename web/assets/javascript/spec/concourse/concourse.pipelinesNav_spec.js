describe("Pipelines Nav", function () {
  var pipelinesNav;

  beforeEach(function () {
    window.pipelineName = "another-pipeline";

    setFixtures(
      '<body><div class="js-pipelinesNav"><div class="js-groups"></div><ul class="js-pipelinesNav-list"></ul><span class="js-pipelinesNav-toggle"></span></div></body>'
    );

    pipelinesNav = new concourse.PipelinesNav($('.js-pipelinesNav'));

    jasmine.Ajax.install();
  });

  afterEach(function() {
    window.pipelineName = undefined;
    jasmine.Ajax.uninstall();
  });

  describe('#bindEvents', function () {
    it('binds on the click of .js-pipelinesNav-toggle', function () {
      pipelinesNav.bindEvents();

      $(".js-pipelinesNav-toggle").trigger('click');
      expect($('body')).toHaveClass('pipelinesNav-visible');

      $(".js-pipelinesNav-toggle").trigger('click');
      expect($('body')).not.toHaveClass('pipelinesNav-visible');
    });

    it('calls to load the pipelines', function() {
      spyOn(pipelinesNav, 'loadPipelines');

      pipelinesNav.bindEvents();

      expect(pipelinesNav.loadPipelines).toHaveBeenCalled();
    });
  });

  describe('#loadPipelines', function() {
    var respondWithSuccess = function(request) {
      var successRequest = request || jasmine.Ajax.requests.mostRecent();
      var successJson = [
      {
        "name": "a-pipeline",
        "url": "/pipelines/a-pipeline",
        "paused": true
      },{
        "name": "another-pipeline",
        "url": "/pipelines/another-pipeline",
        "paused": false
      }];

      successRequest.respondWith({
        "status": 200,
        "contentType": "application/json",
        "responseText":JSON.stringify(successJson)
      });
    };

    var respondWithError = function(request) {
      var errorRequest = request || jasmine.Ajax.requests.mostRecent();
      errorRequest.respondWith({
        "status": 500,
        "contentType": 'application/json'
      });
    };

    it('calls the api endpoint to get the pipelinesNav', function() {
      pipelinesNav.loadPipelines();

      var request = jasmine.Ajax.requests.mostRecent();

      expect(request.url).toBe('/api/v1/pipelines');
      expect(request.method).toBe('GET');

      respondWithSuccess(request);
    });

    describe('when the request is successful', function() {
      it('loads the results into the list', function() {
        expect($('.js-pipelinesNav-list li').length).toEqual(0);

        pipelinesNav.loadPipelines();

        respondWithSuccess();

        expect($('.js-pipelinesNav-list li').length).toEqual(2);

        expect($('.js-pipelinesNav-list li:nth-of-type(1)').data('endpoint')).toEqual('pipelines/a-pipeline');
        expect($('.js-pipelinesNav-list li:nth-of-type(1)').html()).toEqual(
          '<span class="btn-pause fl enabled js-pauseUnpause"><i class="fa fa-fw fa-play"></i></span><a href="/pipelines/a-pipeline">a-pipeline</a>'
        );


        expect($('.js-pipelinesNav-list li:nth-of-type(2)').data('endpoint')).toEqual('pipelines/another-pipeline');
        expect($('.js-pipelinesNav-list li:nth-of-type(2)').html()).toEqual(
          '<span class="btn-pause fl disabled js-pauseUnpause"><i class="fa fa-fw fa-pause"></i></span><a href="/pipelines/another-pipeline">another-pipeline</a>'
        );

        expect($(".js-groups")).not.toHaveClass("paused");

        window.pipelineName = "a-pipeline";
        pipelinesNav.loadPipelines();

        respondWithSuccess();
        expect($(".js-groups")).toHaveClass("paused");
      });

      it('binds events to the .js-pausePipeline buttons via PauseUnpause', function() {
        spyOn(pipelinesNav, 'newPauseUnpause');

        pipelinesNav.loadPipelines();

        respondWithSuccess();

        expect(pipelinesNav.newPauseUnpause).toHaveBeenCalled();
        expect(pipelinesNav.newPauseUnpause.calls.count()).toEqual(2);
      });
    });

    describe("#newPauseUnpause", function() {
      it("creates a new pause unpause from the passed in element and binds the events", function() {
        var pauseUnpauseSpy = jasmine.createSpyObj('pauseUnpause', ['bindEvents']);
        spyOn(concourse, 'PauseUnpause').and.returnValue(pauseUnpauseSpy);

        var myEl = $("<div>");

        pipelinesNav.newPauseUnpause(myEl);

        expect(concourse.PauseUnpause).toHaveBeenCalledWith(myEl, jasmine.any(Function), jasmine.any(Function));
        expect(pauseUnpauseSpy.bindEvents).toHaveBeenCalled();
      });
    });
  });
});